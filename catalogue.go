package main

import (
    "fmt"
    "os"
    "io"
    "log"
    "strings"
    "flag"
    "github.com/gocolly/colly"
)

func getCourse() {
    // Checking a command-line argument was passed
    if len(os.Args) < 2 {
        fmt.Println("Please Enter Course Number")
        os.Exit(1)
    }
    courseNumber := os.Args[1]
    courseName := "cmput"
    defaultFile, err := os.Open("default.txt")

    if err != nil {
        fmt.Println("Default file not found, default will be set to cmput")
    } else {
        defer defaultFile.Close()
        courseNameBytes, err := io.ReadAll(defaultFile)
        courseName = strings.TrimSpace(string(courseNameBytes))
        if err != nil {
            log.Fatal(err)
        }
    }
    

    if len(os.Args) > 2 {
        courseName = os.Args[2]    
    }
    courseName = courseName + "/"

    dashes := strings.Repeat("-", 100)

    c := colly.NewCollector()
    /*
    c.OnResponse(func(r *colly.Response) { 
        fmt.Println("Page visited: ", r.Request.URL) 
    }) 
    */
    c.OnHTML("h1", func(e *colly.HTMLElement) {
        fmt.Println(dashes + "\n")
        fmt.Println(e.Text)
    })
    c.OnHTML("div.mb-3 > div.container > p", func(e *colly.HTMLElement) {
        text := e.Text
        //text = strings.Split(text, "\n")[0]
        text = strings.TrimSpace(text)
        fmt.Println(text + "\n")
    })
    c.Visit("https://apps.ualberta.ca/catalogue/course/" + courseName + courseNumber)
    fmt.Println(dashes)
}

func setCourse() {
    if len(os.Args) < 3 {
        fmt.Println("Please enter default course name")
        os.Exit(1)
    }


    defaultFile, err := os.Create("default.txt")

    if err != nil {
        fmt.Println("Default file not found")
    } 
    defer defaultFile.Close()
    defaultFile.WriteString(os.Args[2])
}

func main() {
    setFlag := flag.Bool("s", false, "-s to set the value of default course")    
    flag.Parse()

    if *setFlag {
        setCourse()
    } else {
        getCourse()
    }       
}
