package main

import (
	"fmt"
	"log"
	"net/http"
	"syscall"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ulimit", getUlimit)

	http.ListenAndServe("0.0.0.0:8000", nil)
}

// Default domain handler
func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("request root URL")
	fmt.Fprint(w, "Server is Alive")
}

func getUlimit(w http.ResponseWriter, r *http.Request) {
	log.Println("check ulimit")
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Fprint(w, "open file : \t")
	fmt.Fprintln(w, rLimit)
	err = syscall.Getrlimit(syscall.RLIMIT_DATA, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Fprint(w, "data seg size : \t")
	fmt.Fprintln(w, rLimit)
	err = syscall.Getrlimit(syscall.RLIMIT_FSIZE, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Fprint(w, "file size : \t")
	fmt.Fprintln(w, rLimit)
	err = syscall.Getrlimit(syscall.RLIMIT_STACK, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Fprint(w, "stack size : \t")
	fmt.Fprintln(w, rLimit)
	err = syscall.Getrlimit(syscall.RLIMIT_CORE, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Fprint(w, "core file size : \t")
	fmt.Fprintln(w, rLimit)
	err = syscall.Getrlimit(syscall.RLIMIT_CPU, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Fprint(w, "cpu size : \t")
	fmt.Fprintln(w, rLimit)
	err = syscall.Getrlimit(syscall.RLIMIT_AS, &rLimit)
	if err != nil {
		fmt.Println("Error Getting Rlimit ", err)
	}
	fmt.Fprint(w, "as size : \t")
	fmt.Fprintln(w, rLimit)
}