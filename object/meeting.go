package object

// Meeting :

import (
 	"strings"
  "fmt"

)

type Meeting struct {
	pr []string
	st string
  et string
	title string
}
func (meeting Meeting) init(pr []string, st string, et string, title string) {
	meeting.pr = pr
	meeting.st = st
	meeting.et = et
	meeting.title = title
}

func (meeting Meeting) CopyMeeting (ano_meeting Meeting) {
	meeting.pr = ano_meeting.pr
	meeting.st = ano_meeting.st
	meeting.et = ano_meeting.et
	meeting.title = ano_meeting.title
}

func (meeting Meeting) GetParticipator() []string {
    return pr
}

func (meeting Meeting) SetParticipator(prs []string) {
	var length= len(prs)
	for i := 0; i < length; i++ {
		meeting.pr[i]= prs[i]
	}
}

func (meeting Meeting) GetStartTime() string {
	return meeting.st
}

func (meeting Meeting) SetStartTime() (st string) {
  meeting.st = st
}

func (meeting Meeting) GetEndTime() string {
	return meeting.et
}

func (meeting Meeting) SetEndTime() (et string) {
	meeting.et = t_endTime
}

func (meeting Meeting) GetTitle() string {
	return meeting.title
}

func _meeting Meeting) SetTitle(title string) {
	meeting.title = title
}

func (meeting Meeting) IsParticipator(un string) bool {
  var i int
	for i= 0; i< len(meeting.pr); i++ {
		if strings.EqualFold(meeting.pr[i], un)== true {
	    	return true
		}
	}
	return false
}

func (meeting Meeting) DeleteParticipator(un string) {
	var i int
	num := len(meeting.pr)
	for i= 0; i< num; i++ {
		if strings.EqualFold(meeting.pr[i], un) == true {
	    meeting.pr = append(meeting.pr[:i], meeting.pr[i+1:]...)
			break
		}
	}
  num = len(meeting.pr)
  if num == 0 {
    //delete meeting
  }
}

func (meeting Meeting)  AddParticipator(un string) bool {
  var i int
  var flag bool
  flag= true
  num := len(meeting.pr)
	for i= 0; i< num; i++ {
		if strings.EqualFold(meeting.pr[i], un) == true {
	    flag = false
      fmt.Println("AddParticipator failed! The user has already in!")  
	    return false
		}
	}
	if flag == true {
		meeting.pr = append(meeting.pr, un)
		return true
	}
	return false
}
