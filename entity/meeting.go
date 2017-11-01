package entity

// Meeting :

import (
  "fmt"
  "os"
  "encoding/json"
  "strings"
  "io/ioutil"
)

type Meeting struct {
  Title string
  Spon string
  Pr []string  
  Start string
  End string
}

type OriginMeeting struct {
  Title ,Spon, Pr, Start, End string
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

func (meeting Meeting) SetStartTime(St string) {
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

func (meeting Meeting) SetEndTime(Et string) {
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
  for _, meeting := range meetingsList {
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
    if strings.EqualFold(meeting.Pr[i], Un) == true {
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
    if strings.EqualFold(meeting.Pr[i], Un) == true {
      flag = false
      fmt.Println("AddParticipator failed! The user has already in!")  
      return false
    }
  }
  if flag == true {
      // if time not overlap
    isOverlap := CheckIfPrTimeOverlap(Un, meeting.Start, meeting.End)
    if isOverlap == true {
      fmt.Println("Overlap")
      return false
    }
    meeting.Pr = append(meeting.Pr, Un)
    return true
  }
  return false
}

func QueryMeetingByTitle(Title string) Meeting {
  meetingsList := getAllMeetings();
  for title, meeting := range meetingsList {
    if title == Title {
      fmt.Println("The meeting you search exits!")
      return meeting
    }
  }

  fmt.Println("No such meeting!") 
  var m Meeting
  return m
}


func ClearMeeting() {
    
}






//ricky part
func split(s rune) bool {
  if s == '/' {
    return true
  }
  return false
}

func getAllMeetings() map[string]Meeting {
  originMeetings := map[string]OriginMeeting{}
  allMeetings := map[string]Meeting{}
  bytes1,_ := ioutil.ReadFile("meetings.json")
  json.Unmarshal(bytes1, &originMeetings)
  for key, value := range originMeetings {
    var temp Meeting
    temp.Spon = value.Spon
    temp.Title = value.Title
    temp.Start = value.Start
    temp.End = value.End
    temp.Pr = strings.FieldsFunc(value.Pr, split)
    allMeetings[key] = temp
  }
  return allMeetings
}

func saveAllMeetings(in map[string]Meeting) {
  allMeetings := map[string]OriginMeeting{}
  for key, value := range in {
    var temp OriginMeeting
    temp.Title = value.Title
    temp.Start = value.Start
    temp.End = value.End
    temp.Spon = value.Spon
    temp.Pr = strings.Join(value.Pr, "/")
    allMeetings[key] = temp
  }
  bytes, _ := json.Marshal(allMeetings)
  fout, _ := os.Create("meetings.json")
  defer fout.Close()
  fout.Write(bytes)
}

func TimeConflig2Meeting(meeting1 Meeting, meeting2 Meeting) bool {
  if meeting1.End < meeting2.Start {
    return false
  }
  if meeting2.End < meeting1.Start {
    return false
  }
  return true
}

func isParOrSpon(meeting Meeting, name string) bool {
  if name == meeting.Spon {
    return true
  }
  for value := range meeting.Pr {
    if meeting.Pr[value] == name {
        return true
    }
  }
  return false
}

func TimeConfligWithExisting(meeting Meeting) bool {
  allMeetings := map[string]Meeting{}
  allMeetings = getAllMeetings()
  sponser:=meeting.Spon
  for _, value := range allMeetings {
    if isParOrSpon(value, sponser) {
      if TimeConflig2Meeting(meeting, value) {
        return true
      }
    }
  }
  for i := range meeting.Pr {
    for _, value := range allMeetings {
      if isParOrSpon(value, meeting.Pr[i]) {
        if TimeConflig2Meeting(meeting, value) {
          return true
        }
      }
    }
  }
  return false
}

func CreateMeeting(title string, pr string, st string, et string) {
  allMeetings := map[string]Meeting{}
  allMeetings = getAllMeetings()
  bytes,_ := ioutil.ReadFile("CurUser")
  curuser := string(bytes)

  temp := strings.FieldsFunc(pr, split)
  for value := range temp {
    if (!UserExists(temp[value])) {
      fmt.Println("Participators not exists")
        return
    }
  }
  
  
  _, exists := allMeetings[title]
  if exists {
    fmt.Println("Meeting already exists")
    return
  }

  if (st > et) {
    fmt.Println("Start time should be earlier than end time")
    return
  }

  var meeting Meeting
  meeting.Title = title
  meeting.Spon = curuser
  meeting.Pr = temp
  meeting.Start = st
  meeting.End = et
  allMeetings[title] = meeting

  if TimeConfligWithExisting(meeting) {
    fmt.Println("Sponser or Participators time-conflict with existing meetings")
    return
  }

  fmt.Println("Create meeting successfully")
  saveAllMeetings(allMeetings)
}

func DeleteMeeting(title string) {
  bytes,_ := ioutil.ReadFile("CurUser")
  curuser := string(bytes)
  allMeetings := map[string]Meeting{}
  allMeetings = getAllMeetings()
  value, exists := allMeetings[title]
  if !exists {
    fmt.Println("Meeting does not exist")
    return
  }
  if (value.Spon != curuser) {
    fmt.Println("Sorry, you are not sponser of this meeting")
      return
  }
  delete(allMeetings, title)
  fmt.Println("Delete meeting successfully")
  saveAllMeetings(allMeetings)
}

    
