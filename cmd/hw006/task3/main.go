package main

import (
	"github.com/rodkevich/go-course/homework/hw006/task03"
)

func main() {
	s := task03.NewWebServer()
	print("127.0.0.1:5050")
	s.Run()
}

	// go s.Run()
	// // check with use of curl
	// args := []string{
	// 	"127.0.0.1:5050",
	// 	"127.0.0.1:5050/call404",
	// }
	// for _, arg := range args {
	//
	// 	out, err := exec.Command("curl","-X", "POST", arg).Output()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Printf("Response:\n %s", out)
	// }
