package scrawle

import (
	"io"
	"net/http"
	"regexp"
)

func GetHtml(url string) string {
    resp, err := http.Get(url)
		if err != nil {
        return err.Error()
		}	
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
    if err != nil {
        return err.Error()
    }
    return string(body)
}


func GetTitles(html string) []string {
	var re = regexp.MustCompile(`<h1[^>]*>(.*?)</h1>`)
	
	headers := re.FindStringSubmatch(html)
	return headers
}

func GetDescription(html string) []string {
	var re = regexp.MustCompile(`<meta\s+name=["']description["'][^>]*content=["']([^"']+)["'][^>]*>`)

	descriptions := re.FindStringSubmatch(html)

	return  descriptions
}

func GetAlternates(html string) []string {
	var re = regexp.MustCompile(`https?://[^\s'"]/[(.*)]/[(.*)]`)

	urls := re.FindStringSubmatch(html)

	return urls
}

func GetRobots(html string) []string {
	var re = regexp.MustCompile(`<meta\s+[^>]*\bname\s*=\s*["']?robots["']?[^>]*>`)

	urls := re.FindStringSubmatch(html)
return urls	
}



