package main

import (
	"fmt"
	"log"
	"net/http"
	utils "web-api-student/utils/handlers"
)

func main() {

	// HandleFunc
	http.HandleFunc("/", utils.HomeHandler)
	http.HandleFunc("/add", utils.AddHandler)
	http.HandleFunc("/view", utils.ViewHandler)
	http.HandleFunc("/stats", utils.StatsHandler)

	//Info in console
	fmt.Println(`
🚀 The server is running on http://localhost:8080
📍 Available routes:
   • GET / (List subjects)
   • GET /stats (statistic subject)
   • POST /add (add new subject)
   • GET /view?id=1 (view subject)
⏹️ To stop pressing a combination of buttons Ctrl+C
	`)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
