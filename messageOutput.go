/*
  Message: Message is a way to manage al message from your system controling its output depending a loglevel configuration
  The log level system follow next values:
    0: info
    1: warn
    2: error
    3: debug

  files:
  -message
*/

package message

import (
  "fmt"
  "io"
  "log"
  "os"
  "time"
)

// The Message object used by sigleton pattern
var msg *Message = nil

// Constants definitions
const INFO int = 0
const WARN int = 1
const ERROR int = 2
const DEBUG int = 3

const layout = "2006-01-02 15:04:00"


// Message struct
// Message its an object which contains all the atributes to manage the message writes into channel
type Message struct {
  // where to write messages
  Writer io.Writer
  // logger for info messages
  logInfo *log.Logger
  // logger for warn messages
  logWarn *log.Logger
  // logger for error messages
  logError *log.Logger
  // logger for debug messages
  logDebug *log.Logger
  // log level definition
  // INFO: 0
  // WARN: 1
  // ERROR: 2
  // DEBUG: 3
  logLevel int

  // channel to write messages
  mChan chan []interface{}
  // channel to quit from print machine
  quitChan chan bool
}


//
// generate a new Message object
func New(l int, w io.Writer, f int) *Message {

  if ( msg == nil ){
    // The info loglevel is set if an incorrect value would be configured
    if l < 0 || l > 3 { l = 0 }
    if w == nil { w = os.Stdout }

    msg = &Message {
      Writer: w,
      logInfo: log.New(w, "[INFO] ",f),
      logWarn: log.New(w, "[WARN] ",f),
      logError: log.New(w, "[ERROR] ",f),
      logDebug: log.New(w, "[DEBUG] ",f),
      mChan: make(chan []interface{}),
      quitChan: make(chan bool),
      logLevel: l,
    }
  }
  return msg
}


// GetInstance: return an instance of the object Message. If no instance has been created, a new one is created
func GetInstance(l int) *Message {
  if msg == nil {
    // The info loglevel is set if an incorrect value would be configured
    if l < 0 || l > 3 { l = 0 }

    msg = new(Message)
    c := make(chan []interface{})
    q := make(chan bool)

    msg = &Message{
      mChan: c,
      quitChan: q,
      logLevel: l,
    }
    // Starting the print machine
    go msg.printMachine()

  }else{
    // change loglevel if instance is already running
    msg.SetLogLevel(l)
  }

  return msg
}

//
// DestroyInstance method stops the printMachine
func (m *Message) DestroyInstance(){
  msg.quitChan <- true
}
//
// SetLogLevel method set the loglevel to the gived one
func (m *Message) SetLogLevel(l int){
  if l < 0 || l > 3 { l = 0 }
  m.logLevel = l
}

//
// printMachine method waits for messages to write till the instance is destroyed
func (m *Message) printMachine(){

  fi := false
  for ;!fi; {
  select{
    case msg := <-m.mChan:
      fmt.Println(msg)
    case <-m.quitChan:
      fi = true
    }
  }
  // close channels
  defer close(m.mChan)
  defer close(m.quitChan)
}

//
// Info write info messages to logger
func (m *Message) Info(msg... interface{}){
  if m.logLevel >= INFO {
    m.logInfo.Println(msg)
  }
}

//
// Warn write warning messages to logger
func (m *Message) Warn(msg... interface{}){
  if m.logLevel >= WARN {
    m.logWarn.Println(msg)
  }
}

//
// Error write error messages to logger
func (m *Message) Error(msg... interface{}){
  if m.logLevel >= ERROR {
    m.logError.Println(msg)
  }
}

//
// Debug write debug messages to logger
func (m *Message) Debug(msg... interface{}){
  if m.logLevel >= DEBUG {
    m.logDebug.Println(msg)
  }
}

//
// WriteCh always send a message to be written
func (m *Message) WriteCh(msg... interface{}){
  m.mChan <- msg
}
//
// WriteChInfo send a message to be written by printMachine if the loglevel is greater or equal to info
func (m *Message) WriteChInfo(msg... interface{}){
  if m.logLevel >= INFO {
      t := time.Now()
      m.mChan <- append([]interface{}{t.Format(layout)," INFO:"},msg...)
  }
}
//
// WriteChWarn send a message to be written by printMachine if the loglevel is greater or equal to warn
func (m *Message) WriteChWarn(msg... interface{}){
  if m.logLevel >= WARN {
      t := time.Now()
      m.mChan <- append([]interface{}{t.Format(layout)," WARN:"},msg...)
  }
}
//
// WriteChError send a message to be written by printMachine if the loglevel is greater or equal to error
func (m *Message) WriteChError(msg... interface{}){
  if m.logLevel >= ERROR {
      t := time.Now()
      m.mChan <- append([]interface{}{t.Format(layout)," ERROR:"},msg...)
  }
}
//
// WriteChDebug send a message to be written by printMachine if the loglevel is greater or equal to debug
func (m *Message) WriteChDebug(msg... interface{}){
  if m.logLevel >= DEBUG {
      t := time.Now()
      m.mChan <- append([]interface{}{t.Format(layout)," DEBUG:"},msg...)
  }
}

//
// Write a message wrapping fmt
func Write(message string) {
	fmt.Println(message)
}
//
// Write a message wrapping fmt
func WriteInfo(message interface{}) {
  fmt.Println("INFO", message)
}
//
// Write a message wrapping fmt
func WriteWarn(message interface{}) {
  fmt.Println("WARN", message)
}
//
// Write a message wrapping fmt
func WriteError(message interface{}) {
  fmt.Println("ERROR", message)
}
//
// Write a message wrapping fmt
func WriteDebug(message interface{}) {
  fmt.Println("DEBUG", message)
}
