package object

import (
	"os"
	"io"
	"log"
	"encoding/json"
	"path/filepath"
	"Agenda/logger"
	"strings"
)

type User struct {
	Name, Password, Email, Phone string
}


//fmt.Println(strings.FieldsFunc("widuunhellonword", split)) // [widuu hello word]根据n字符分割

func split(s rune) bool {
	if s == '/' {
 		return true
	}
	return false
}

func getAllUsers() map[string]User {

}

func storeAllUsers(map[string]User) {

}

func getAllMeetings() map[string]Meeting {
	
}

func storeAllMeetings(map[string]Meeting) {
	
}
