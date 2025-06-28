package main

import (
	"fmt"
	"web-scrapper/scrawle"
)

var data []string


func main() {
	 resp := scrawle.GetHtml("https://devbox.tools/")	
	 headers := scrawle.GetTitles(resp)
	 descriptions := scrawle.GetDescription(resp)
	 alternates := scrawle.GetAlternates(resp)
	
	
	 data = append(data, headers...)
	 data = append(data, descriptions...)
	 data = append(data, alternates...)
	 fmt.Println(alternates)
}
