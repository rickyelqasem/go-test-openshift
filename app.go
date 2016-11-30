package main

import (
	"fmt"
	owm "github.com/briandowns/openweathermap"
	"log"
	"net/http"
	"os"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {

	we, err := owm.NewCurrent("C", "en")
	if err != nil {
		log.Fatalln(err)
	}

	we.CurrentByName("London")

	s := fmt.Sprint(we.Weather)

	strArray := strings.Fields(s)

	var weather string = string(strArray[2])
	weather = weather + " " + string(strArray[3])
	//var weather string = getWeather()
	fmt.Fprintf(w, "Hi there Ricky, the weather in London is "+ weather +" %s!", r.URL.Path[1:])
}

func main() {
	os.Setenv("OWM_API_KEY", "c0ce4a2db13e8061f06fc7a77db850f1")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}


