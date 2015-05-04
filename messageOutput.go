/*
*/

package message 

import (
  "fmt"
)


var msg *message = nil

var INFO int = 0
var WARN int = 1
var ERROR int = 2
var DEBUG int = 3

// Define the message object
/*
  mode
    0: info
    1: warn
    2: error
    3: debug
*/
type message struct {
  mChan chan []interface{}
  quitChan chan bool
  logMode int
}



func GetInstance(l int) *message {
  if msg == nil {
    if l < 0 || l > 3 { l = 0 }
    msg = new(message)
    c := make(chan []interface{})
    q := make(chan bool)

    msg = &message{
      mChan: c,
      quitChan: q,
      logMode: l,
    }
    go printMachine()

  }else{
    msg.SetLogMode(l)
  }

  return msg
}

func (m *message) DestroyInstance(){
  msg.quitChan <- true
}

func (m *message) SetLogMode(l int){
  if l < 0 || l > 3 { l = 0 }
  m.logMode = l
}

func printMachine(){
  fi := false
  for ;!fi; {
  select{
    case m := <-msg.mChan:
      fmt.Println(m)
    case <-msg.quitChan:
      fi = true
    }
  }
  
  defer close(msg.mChan)
  defer close(msg.quitChan)
}


func (m *message) TestWrite(msg... interface{}){
  m.mChan <- msg
}
func (m *message) TestWriteInfo(msg... interface{}){
  if m.logMode >= INFO {
      m.mChan <- append([]interface{}{"INFO:"},msg...)
  }
}
func (m *message) TestWriteWarn(msg... interface{}){
  if m.logMode >= WARN {
      m.mChan <- append([]interface{}{"WARN:"},msg...)
  }
}
func (m *message) TestWriteError(msg... interface{}){
  if m.logMode >= ERROR {
      m.mChan <- append([]interface{}{"ERROR:"},msg...)
  }
}
func (m *message) TestWriteDebug(msg... interface{}){
  if m.logMode >= DEBUG {
      m.mChan <- append([]interface{}{"DEBUG:"},msg...)
  }
}

func Write(message string) {
	fmt.Println(message)
}

func WriteInfo(message interface{}) {
  fmt.Println("INFO", message)
}
func WriteWarn(message interface{}) {
  fmt.Println("WARN", message)
}
func WriteError(message interface{}) {
  fmt.Println("ERROR", message)
}
func WriteDebug(message interface{}) {
  fmt.Println("DEBUG", message)
}
