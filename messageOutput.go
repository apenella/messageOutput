package message 

import "fmt"

/*
  mode
    0: info
    1: warn
    2: error
    3: debug
*/
// Define the how to messaging to user
/*
struct message type {
  stdEnable bool
  stdMode int
  logEnable bool
  logMode int
}*/

func Write(message string) {
	WriteInfo(message)
}

func WriteInfo(message string) {
  fmt.Println("INFO", message)
}
func WriteWarn(message string) {
  fmt.Println("WARN", message)
}
func WriteError(message string) {
  fmt.Println("ERROR", message)
}
func WriteDebug(message string) {
  fmt.Println("DEBUG", message)
}
