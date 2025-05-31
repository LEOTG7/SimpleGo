package main

import (
       "fmt"
       "net/http"
       "runtime/debug"
       "time"

       "github.com/LEOTG7/SimpleGo/plugins"

)

func main() {
	defer func() {
		if r := recover(); r != nil {
			// Print reason for panic + stack for some sort of helpful log output
			fmt.Println(r)
			fmt.Println(string(debug.Stack()))
		}
	}()

	
       go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(w, "Hi Hi")
		})

		http.ListenAndServe(":"+config.Port, nil)
	}()
