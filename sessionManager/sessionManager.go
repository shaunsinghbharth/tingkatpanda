package sessionManager

import (
	"encoding/json"
	"fmt"
	. "github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"os"
)

type SessionManager struct{
	Users []Users `json: "Users"`
}

type Users struct{
	UUID 		string `json: "UUID"`
	Type		string `json: "Type"`
	Username  	string `json: "Username"`
	Password  	string `json: "Password"`
	Firstname 	string `json: "Firstname"`
	Lastname  	string `json: "Lastname"`
}


var ManagerSingletonInstance *SessionManager

func (session *SessionManager)Init() *SessionManager{
	session.Users = []Users{}

	userFile, _ := os.Open("JSON/user.json")
	bytes, err := ioutil.ReadAll(userFile)

	if err == nil{
		//fmt.Println(string(bytes))
		json.Unmarshal(bytes, &session.Users)
		fmt.Println(session)
	}
	defer userFile.Close()

	ManagerSingletonInstance = session
	return ManagerSingletonInstance
}

func GetSessionManager() *SessionManager{
	return ManagerSingletonInstance
}

func (session *SessionManager)GetUUID(index int) string{
	return session.Users[index].UUID
}

func (session *SessionManager)GetUserName(index int) string{
	return session.Users[index].Username
}

func (session *SessionManager)GetFirstName(index int) string{
	return session.Users[index].Firstname
}

func (session *SessionManager)GetLastName(index int) string{
	return session.Users[index].Lastname
}

func (session *SessionManager)GetPassword(index int) string{
	return session.Users[index].Password
}

func (session *SessionManager)RegisterUser(username string, password string, firstname string, lastname string) *http.Cookie{
	tempSession, _ := NewV4()
	//if err != nil{
		var tempUser = Users{UUID: tempSession.String(), Username: username, Password: password, Firstname: firstname, Lastname: lastname}
		session.Users = append(session.Users, tempUser)
	//}


	setcookie := &http.Cookie{
		Name:       "BharthPHD",
		Value:      session.Users[len(session.Users) - 1].UUID,
		Path:       "/",
	}

	session.Dump()
	return setcookie
}

func (session *SessionManager)CreateSession(username string, password string) *http.Cookie{
	fmt.Println("LOGIN")
	for index := 0;index < len(session.Users); index = index + 1{
		fmt.Println(session.Users[index].Username, username)
		if session.Users[index].Username == username && session.Users[index].Password == password{
			token, _ := NewV4()
			setcookie := &http.Cookie{
				Name:       "BharthPHD",
				Value:      token.String(),
				Path: "/",
			}
			session.Users[index].UUID = token.String()

			return setcookie
		}
	}
	fmt.Println("USERNAME or PASSWORD INVALID")
	return nil
}

func (session *SessionManager)ValidSession(req *http.Request) bool{
	cookie , err := req.Cookie("BharthPHD")

	if cookie != nil || err != nil{
		return true
	}
	return false
}

func (session *SessionManager)GetCurrentUser(req *http.Request) string{
	cookie , err := req.Cookie("BharthPHD")

	if err == nil{
		for _,v := range session.Users{
			if v.UUID == cookie.Value{
				return v.Username
			}
		}
	}
	return "-1"
}

func (session *SessionManager)GetCurrentUserObject(req *http.Request) Users {
	cookie , err := req.Cookie("BharthPHD")

	if err == nil{
		for _,v := range session.Users{
			if v.UUID == cookie.Value{
				return v
			}
		}
	}
	return Users{}
}

func (session *SessionManager)Dump() error{
	var jsonData []byte
	fmt.Println(session)
	jsonData, err := json.MarshalIndent(session.Users, "", "    ")
	if err != nil{
		fmt.Println("DUMP ERROR")
	}
	fmt.Println("JSON DUMP")
	fmt.Println(string(jsonData))

	ioutil.WriteFile("JSON/user.json", jsonData, 0777)
	return err
}

func (session *SessionManager)DestroySession(req *http.Request) *http.Cookie{
	setcookie := &http.Cookie{
		Name:       "BharthPHD",
		Value:      "-1",
		Path: "/",
		MaxAge: -1,
	}
	return setcookie
}

func (session *SessionManager)EditUser(uuid string, password string, username string, firstname string, lastname string){
	for i,v := range session.Users{
		if v.UUID == uuid{
			session.Users[i].Username = username
			session.Users[i].Firstname = firstname
			session.Users[i].Lastname = lastname
			session.Users[i].Password = password
		}
	}
}

func (session *SessionManager)IsAdmin(req *http.Request) bool{
	tempUser := session.GetCurrentUserObject(req)

	if tempUser.Type == "admin"{
		return true
	}
	return false
}

func (session *SessionManager)DeleteUser(username string) bool{
	for i,v := range session.Users{
		if v.Username == username{
			session.Users[i].Username = ""
			session.Users[i].Type = ""
			session.Users[i].UUID = ""
			session.Users[i].Password = ""
			session.Users[i].Firstname = ""
			session.Users[i].Lastname = ""

			return true
		}
	}
	return false
}

func (session *SessionManager)DeleteSession(username string) bool{
	for i,v := range session.Users{
		if v.Username == username{
			session.Users[i].UUID = ""

			return true
		}
	}
	return false
}