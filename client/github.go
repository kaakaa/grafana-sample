package main

import (
    "log"
    "os"
    "encoding/json"
    "github.com/google/go-github/github"
)

func main() {
    client := github.NewClient(nil)
    issues, _, err := client.Issues.ListRepositoryEvents("kaakaa", "ppt-museum", nil)
    if err != nil {
        log.Fatalln(err)
    }
    d, _ := json.MarshalIndent(issues[0], "", "  ")
    os.Stdout.Write(d)
    log.Println(len(issues))
}
