/*
*/

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
