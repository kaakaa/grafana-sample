package main

import (
    //"net/url"
    "log"
    
    "github.com/influxdata/influxdb/client/v2"
    "github.com/google/go-github/github"
)

const (
    MYDB = "github_sample"
    username = "root"
    password = "root"
)

func main() {
    // Make client
    c, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: "http://docker:8086",
        Username: username,
        Password: password,
    })
    
    if err != nil {
        log.Fatalln("Error: ", err)
    }
    
    // Write the batch
    writePoints(c)
}

func writePoints(clnt client.Client) {
    owner := "kaakaa"
    repo := "ppt-museum"

    bp, _ := client.NewBatchPoints(client.BatchPointsConfig{
        Database: MYDB,
        Precision: "us",
    })
    
    commits, _ := get(owner, repo)
    for i := 0; i < len(commits); i++ {
        tags := map[string]string{
            "owner":    owner,
            "repo":     repo,
        }
        
        c := commits[i]
        dat := map[string]interface{}{
            "sha": *c.SHA,
            "committer_id": *c.Committer.ID,
            "committer_name": *c.Committer.Login,
            "committer_email": *c.Commit.Committer.Email,
            "author_id": *c.Author.ID,
            "author_name": *c.Author.Login,
            "author_email": *c.Commit.Author.Email,
        }
        pt, _ := client.NewPoint(owner + "/" + repo, tags, dat, *c.Commit.Committer.Date)
        bp.AddPoint(pt)
    }
    
    err := clnt.Write(bp)
    if err != nil {
        log.Fatal(err)
    }
}


func get(owner string, repo string) ([]github.RepositoryCommit, error) {
    client := github.NewClient(nil)
    c, _, err := client.Repositories.ListCommits(owner, repo, nil)
    
    return c, err
}

