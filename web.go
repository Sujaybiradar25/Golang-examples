package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("In Hello")
	fmt.Fprintf(w, "In hello")
}

func newPageHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("In new Page")
	log.Print(r.URL.Host)
	fmt.Fprintf(w, "In new Page")
}
func mainPage(w http.ResponseWriter, r *http.Request) {
	log.Print("On homepage")
	log.Print(r.URL.Host)
	fmt.Fprintf(w, "In home page")
}

/**
func main() {
	//handle specific URL path
	//http.HandleFunc("/hello", helloHandler)
	//http.HandleFunc("/newpage", newPageHandler)
	r := mux.NewRouter()
	r.HandleFunc("/", mainPage)
	r.HandleFunc("/hello", helloHandler)
	r.HandleFunc("/newpage", newPageHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
**/
func main() {
	task := []bool{false, false}

	temp := 0
	access := 1
	for true {
		go func(temp int) {
			for i := 0; i < 10; {
				access = 0
				task[1] = true
				fmt.Println("in section 2")
				for access == 0 && task[0] == true {
				}
				temp++
				task[1] = false
				access = 1
				fmt.Println(temp)
			}
		}(temp)

		for i := 0; i < 10; {
			task[0] = true
			access = 1
			fmt.Println("in section 1")
			for access == 1 && task[1] == true {
			}
			temp++
			access = 0
			task[0] = false
			fmt.Println(temp)
		}
	}
	//return 0

}
