package main

import (
    "fmt"
    "os"
    "strings"
    "github.com/gocolly/colly"
)

func main() {
    // Checking a command-line argument was passed
    if len(os.Args) < 2 {
        fmt.Println("Please Enter Course Number")
        os.Exit(1)
    }
    courseNumber := os.Args[1]
    courseName := "cmput"

    if len(os.Args) > 2 {
        courseName = os.Args[2]    
    }
    courseName = courseName + "/"

    dashes := strings.Repeat("-", 100)
    fmt.Println(dashes)

    c := colly.NewCollector()
    /*
    c.OnResponse(func(r *colly.Response) { 
        fmt.Println("Page visited: ", r.Request.URL) 
    }) 
    */
    c.OnHTML("h1", func(e *colly.HTMLElement) {
        fmt.Println(e.Text)
    })
    c.OnHTML("div.mb-3 > div.container > p", func(e *colly.HTMLElement) {
        text := e.Text
        //text = strings.Split(text, "\n")[0]
        text = strings.TrimSpace(text)
        fmt.Println(text + "\n")
    })
    c.Visit("https://apps.ualberta.ca/catalogue/course/" + courseName + courseNumber)
}
