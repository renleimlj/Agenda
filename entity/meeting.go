package object

// Meeting :

import (
  "strings"
  "fmt"
  "time"

)

type Meeting struct {
  Title string
  Spon string
  Pr []string  
  Start string
  End string
}

func (meeting Meeting) Init(Pr []string, St string, Et string, Title string) {
  meeting.Pr = Pr
  meeting.Start = St
  meeting.End = Et
  meeting.Title = Title
}

func (meeting Meeting) CopyMeeting (ano_meeting Meeting) {
  meeting.Pr = ano_meeting.Pr
  meeting.Start = ano_meeting.Start
  meeting.End = ano_meeting.End
  meeting.Title = ano_meeting.Title
}

func (meeting Meeting) GetParticipator() []string {
    return meeting.Pr
}

func (meeting Meeting) SetParticipator(prs []string) {
  var length = len(prs)
  for i := 0; i < length; i++ {
    meeting.Pr[i]= prs[i]
  }
}

func (meeting Meeting) GetStartTime() string {
  return meeting.Start
}

func (meeting Meeting) SetStartTime() (St string) {
  meeting.Start = St
}

func (meeting Meeting) GetSponsor() string{
  return meeting.Spon
} 

func (meeting Meeting) SetSponsor(Sp string) {
  meeting.Spon = Sp
} 

func (meeting Meeting) GetEndTime() string {
  return meeting.End
}

func (meeting Meeting) SetEndTime() (Et string) {
  meeting.End = Et
}

func (meeting Meeting) GetTitle() string {
  return meeting.Title
}

func (meeting Meeting) SetTitle(Title string) {
  meeting.Title = Title
}

func (meeting Meeting) IsParticipator(Un string) bool {
  if Un == "" {
    fmt.Println("Null Participator Name!")
    return false
  }
  var i int
  for i= 0; i< len(meeting.Pr); i++ {
    if strings.EqualFold(meeting.Pr[i], Un)== true {
      return true
    }
  }
  return false
}

func CheckIfMeetingTimeOverlap(startTime string, endTime string) bool {
  meetingsList := getAllMeetings(); 
  for tit, meeting := range meetingsList; {
    if meeting.Start < endTime && meeting.Start > startTime {
      return true
    } else if meeting.End < endTime && meeting.End > startTime{
      return true
    }
  }
  return false
}

func CheckIfPrTimeOverlap(Un string, startTime string, endTime string) bool {
  flag := CheckIfMeetingTimeOverlap(startTime, endTime)
  if flag == true {
    return true
  } else {
    meetingsList := getAllMeetings(); 
    for tit, meeting := range meetingsList; {
      for pr := meeting.Pr {
        if pr == Un || meeting.spon == Un {
          if meeting.Start < endTime && meeting.Start > startTime {
            return true
          } else if meeting.End < endTime && meeting.End > startTime{
            return true
          }
        }
      }
    }
  }
  return false
}

func (meeting Meeting) DeleteParticipator(Un string) {
  if Un == "" {
    fmt.Println("Null Participator Name!")
    return
  }
  var i int
  num := len(meeting.Pr)
  for i= 0; i< num; i++ {
    if strings.EqualFold(meeting.Pr[i], un) == true {
      meeting.Pr = append(meeting.Pr[:i], meeting.Pr[i+1:]...)
      break
    }
  }
  num = len(meeting.Pr)
  if num == 0 {
    DeleteMeeting(meeting.Title)
  }
}

func (meeting Meeting) AddParticipator(Un string) bool {
  if Un == "" {
    fmt.Println("Null Participator Name!")
    return false
  }
  var i int
  var flag bool
  flag = true
  num := len(meeting.Pr)
  for i= 0; i< num; i++ {
    if strings.EqualFold(meeting.Pr[i], un) == true {
      flag = false
      fmt.Println("AddParticipator failed! The user has already in!")  
      return false
    }
  }
  if flag == true {
      // if time not overlap
    isOverlap := checkIfPrTimeOverlap(Un, meetings.Start, meetings.End)
    if isOverlap == true {
      fmt.Println("Overlap\n")
      return false
    }
    meeting.Pr = append(meeting.Pr, un)
    return true
  }
  return false
}

func QueryMeetingByTitle(Title string) Meeting {
  meetingsList := getAllMeetings();
  var count int
  for title, meeting := range meetingsList; i++ {
    if title == Title {
      fmt.Println("The meeting you search exits!\n")
      return meeting
    }
  }

  fmt.Println("No such meeting!\n") 
  return nil
}

func DeleteMeeting(Title string) {
    
}

func ClearMeeting() {
    
}

    
