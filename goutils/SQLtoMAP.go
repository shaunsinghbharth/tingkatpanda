package goutils

import (
	"database/sql"
)

func SQLtoMap(results *sql.Rows) []map[string]interface{} {
	columns, _ := results.Columns()
	var resultMap map[string]interface{}

	var resultArray = make([]map[string]interface{}, 0)

	for results.Next() {
		resultMap = make(map[string]interface{})
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i,_ := range values {
			pointers[i] = &values[i]
		}

		results.Scan(pointers...)

		for i, val := range values {
			//fmt.Printf("Adding key=%s val=%v\n", columns[i], val)
			switch value := val.(type) {
			case string:
				resultMap[columns[i]] = value
			case int64:
				resultMap[columns[i]] = value
			case []uint8:
				resultMap[columns[i]] = string(value)
			case float64, float32:
				resultMap[columns[i]] = value
			}
		}
		resultArray = append(resultArray, resultMap)
	}

	//fmt.Println("SQLtoMAP ", resultArray)
	return resultArray
}
