package main

import (
	"fmt"
	"time"
	"web-scrapper/scrawle"
)

var data  = make(map[string][]string)

func main() {
		currentTime:= time.Now()

	 resp := scrawle.GetHtml("https://devbox.tools/")	

	 headers := scrawle.GetTitles(resp)
	 descriptions := scrawle.GetDescription(resp)
	 alternates := scrawle.GetAlternates(resp)
	 robots := scrawle.GetRobots(resp)
	 keywords := scrawle.GetKeyWords(resp) 

	 data["title"] = headers
	 data["descriptions"] = descriptions
	 data["alternates"] = alternates
	 data["robots"] = robots 
	 data["keywords"] = keywords

	 endTime := time.Since(currentTime)
	 fmt.Println(endTime)
	 fmt.Println(data)
 }
