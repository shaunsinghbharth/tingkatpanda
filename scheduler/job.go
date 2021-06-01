package scheduler

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

var (
	ErrTimeFormat           = errors.New("time format error")
	ErrParamsNotAdapted     = errors.New("the number of params is not adapted")
	ErrNotAFunction         = errors.New("only functions can be schedule into the job queue")
	ErrPeriodNotSpecified   = errors.New("unspecified job period")
	ErrParameterCannotBeNil = errors.New("nil parameters cannot be used with reflection")
)

// Job struct keeping information about job
type Job struct {
	interval uint64                   // pause interval * unit between runs
	jobFunc  string                   // the job jobFunc to run, func[jobFunc]
	unit     timeUnit                 // time units, ,e.g. 'minutes', 'hours'...
	atTime   time.Duration            // optional time at which this job runs
	err      error                    // error related to job
	loc      *time.Location           // optional timezone that the atTime is in
	lastRun  time.Time                // datetime of last run
	nextRun  time.Time                // datetime of next run
	startDay time.Weekday             // Specific day of the week to start on
	funcs    map[string]interface{}   // Map for the function task store
	fparams  map[string][]interface{} // Map for function and  params of function
	lock     bool                     // lock the job from running at same time form multiple instances
	tags     []string                 // allow the user to tag jobs with certain labels
}

// NewJob creates a new job with the time interval.
func NewJob(interval uint64) *Job {
	return &Job{
		interval: interval,
		loc:      loc,
		lastRun:  time.Unix(0, 0),
		nextRun:  time.Unix(0, 0),
		startDay: time.Sunday,
		funcs:    make(map[string]interface{}),
		fparams:  make(map[string][]interface{}),
		tags:     []string{},
	}
}

// True if the job should be run now
func (j *Job) shouldRun() bool {
	return time.Now().Unix() >= j.nextRun.Unix()
}

//Run the job and immediately reschedule it
func (j *Job) run() ([]reflect.Value, error) {
	if j.lock {
		if locker == nil {
			return nil, fmt.Errorf("trying to lock %s with nil locker", j.jobFunc)
		}
		key := getFunctionKey(j.jobFunc)

		locker.Lock(key)
		defer locker.Unlock(key)
	}
	result, err := callJobFuncWithParams(j.funcs[j.jobFunc], j.fparams[j.jobFunc])
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Err should be checked to ensure an error didn't occur creating the job
func (j *Job) Err() error {
	return j.err
}

// Do specifies the jobFunc that should be called every time the job runs
func (j *Job) Do(jobFun interface{}, params ...interface{}) error {
	if j.err != nil {
		return j.err
	}

	typ := reflect.TypeOf(jobFun)
	if typ.Kind() != reflect.Func {
		return ErrNotAFunction
	}
	fname := getFunctionName(jobFun)
	j.funcs[fname] = jobFun
	j.fparams[fname] = params
	j.jobFunc = fname

	now := time.Now().In(j.loc)
	if !j.nextRun.After(now) {
		j.scheduleNextRun()
	}

	return nil
}

// At schedules job at specific time of day
//	s.Every(1).Day().At("10:30:01").Do(task)
//	s.Every(1).Monday().At("10:30:01").Do(task)
func (j *Job) At(t string) *Job {
	hour, min, sec, err := formatTime(t)
	if err != nil {
		j.err = ErrTimeFormat
		return j
	}
	// save atTime start as duration from midnight
	j.atTime = time.Duration(hour)*time.Hour + time.Duration(min)*time.Minute + time.Duration(sec)*time.Second
	return j
}

// GetAt returns the specific time of day the job will run at
//	s.Every(1).Day().At("10:30").GetAt() == "10:30"
func (j *Job) GetAt() string {
	return fmt.Sprintf("%d:%d", j.atTime/time.Hour, (j.atTime%time.Hour)/time.Minute)
}

func (j *Job) periodDuration() (time.Duration, error) {
	interval := time.Duration(j.interval)
	var periodDuration time.Duration

	switch j.unit {
	case seconds:
		periodDuration = interval * time.Second
	case minutes:
		periodDuration = interval * time.Minute
	case hours:
		periodDuration = interval * time.Hour
	case days:
		periodDuration = interval * time.Hour * 24
	case weeks:
		periodDuration = interval * time.Hour * 24 * 7
	default:
		return 0, ErrPeriodNotSpecified
	}
	return periodDuration, nil
}

// roundToMidnight truncate time to midnight
func (j *Job) roundToMidnight(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, j.loc)
}

// scheduleNextRun Compute the instant when this job should run next
func (j *Job) scheduleNextRun() error {
	now := time.Now()
	if j.lastRun == time.Unix(0, 0) {
		j.lastRun = now
	}

	periodDuration, err := j.periodDuration()
	if err != nil {
		return err
	}

	switch j.unit {
	case seconds, minutes, hours:
		j.nextRun = j.lastRun.Add(periodDuration)
	case days:
		j.nextRun = j.roundToMidnight(j.lastRun)
		j.nextRun = j.nextRun.Add(j.atTime)
	case weeks:
		j.nextRun = j.roundToMidnight(j.lastRun)
		dayDiff := int(j.startDay)
		dayDiff -= int(j.nextRun.Weekday())
		if dayDiff != 0 {
			j.nextRun = j.nextRun.Add(time.Duration(dayDiff) * 24 * time.Hour)
		}
		j.nextRun = j.nextRun.Add(j.atTime)
	}

	// advance to next possible schedule
	for j.nextRun.Before(now) || j.nextRun.Before(j.lastRun) {
		j.nextRun = j.nextRun.Add(periodDuration)
	}

	return nil
}


// Loc sets the location for which to interpret "At"
//	s.Every(1).Day().At("10:30").Loc(time.UTC).Do(task)
func (j *Job) Loc(loc *time.Location) *Job {
	j.loc = loc
	return j
}

// NextScheduledTime returns the time of when this job is to run next
func (j *Job) NextScheduledTime() time.Time {
	return j.nextRun
}

// set the job's unit with seconds,minutes,hours...
func (j *Job) mustInterval(i uint64) error {
	if j.interval != i {
		return fmt.Errorf("interval must be %d", i)
	}
	return nil
}

// From schedules the next run of the job
func (j *Job) From(t *time.Time) *Job {
	j.nextRun = *t
	return j
}

// setUnit sets unit type
func (j *Job) setUnit(unit timeUnit) *Job {
	j.unit = unit
	return j
}

// Seconds set the unit with seconds
func (j *Job) Seconds() *Job {
	return j.setUnit(seconds)
}

// Minutes set the unit with minute
func (j *Job) Minutes() *Job {
	return j.setUnit(minutes)
}

// Hours set the unit with hours
func (j *Job) Hours() *Job {
	return j.setUnit(hours)
}

// Second sets the unit with second
func (j *Job) Second() *Job {
	j.mustInterval(1)
	return j.Seconds()
}

// Minute sets the unit  with minute, which interval is 1
func (j *Job) Minute() *Job {
	j.mustInterval(1)
	return j.Minutes()
}

// Hour sets the unit with hour, which interval is 1
func (j *Job) Hour() *Job {
	j.mustInterval(1)
	return j.Hours()
}

// GetWeekday returns which day of the week the job will run on
// This should only be used when .Weekday(...) was called on the job.
func (j *Job) GetWeekday() time.Weekday {
	return j.startDay
}

// Lock prevents job to run from multiple instances of gocron
func (j *Job) Lock() *Job {
	j.lock = true
	return j
}
