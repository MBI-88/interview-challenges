package main

import (
	"fmt"
	"gitreports/github"
	"gitreports/editor"
	"log"
	"os"
	"text/template"
	"time"
)

const usage = `usage:
  create OWNER REPO
  get    OWNER REPO ISSUE_NUMBER
  edit   OWNER REPO ISSUE_NUMBER
  close  OWNER REPO ISSUE_NUMBER
  reopen OWNER REPO ISSUE_NUMBER
`

var templ = template.Must(template.New("issue").Funcs(template.FuncMap{"formatTime": formatTime}).Parse(`
	Number:   {{.Number}}
	URL:      {{.HTMLURL}}
	User:     {{.User.Login}}
	Title:    {{.Title | printf "%.64s"}}
	State:    {{.State}}
	Comments: {{.Comments}}
	Created:  {{.CreatedAt | formatTime}}
	Updated:  {{.UpdatedAt | formatTime}}

	{{if ne (len .Body) 0}}{{.Body}}{{else}}(no body){{end}}
`))

func formatTime(t time.Time) string {
	return t.Format("2023-01-01 00:00:00")
}

func get(owner, repo, number string) {
	result,err := github.Get(owner,repo,number)
	if err != nil {log.Fatalf("Error in get %s",err)}
	templ.Execute(os.Stdout,result)

}

func create(owner, repo string) {
	fields := map[string]string{
		"title": "",
		"body":"",
	}
	// Editor
   	if 	er := editor.Editor(fields); er != nil {log.Fatalf("error in editor %s",er)}

    if err := github.Post(owner,repo,fields); err != nil {
		log.Fatalf("Error in create %s",err)
	}
}

func edit(owner, repo, number string) {
	resp,err := github.Get(owner,repo,number)
	if err != nil {log.Fatalf("error in edit %s",err)}
	fields := map[string]string{
		"title":resp.Title,
		"body":resp.Body,
	}
	// Editor
	if 	er := editor.Editor(fields); er != nil {log.Fatalf("error in editor %s",er)}

	if err := github.Update(owner,repo,number,fields);err != nil {log.Fatalf("error in update %s",err)}

}
func close(owner, repo, number string) {
	if err := github.Close(owner,repo,number); err != nil {log.Fatalf("error in close %s",err)}

}
func reopen(owner, repo, number string) {
	if err := github.Reopen(owner,repo,number); err != nil { log.Fatalf("error in reopen %s",err)}

}

func main() {
	if len(os.Args) == 4 {
		command, owner, repo := os.Args[1], os.Args[2], os.Args[3]
		switch command {
		case "create":
			create(owner, repo)
		default:
			fmt.Fprintf(os.Stderr, usage)
			os.Exit(1)
		}
	} else if len(os.Args) == 5 {
		conmand, owner, repo, number := os.Args[1], os.Args[2], os.Args[3], os.Args[4]
		switch conmand {
		case "get":
			get(owner, repo, number)
		case "edit":
			edit(owner, repo, number)
		case "close":
			close(owner, repo, number)
		case "reopen":
			reopen(owner, repo, number)
		default:
			fmt.Fprintf(os.Stderr, usage)
			os.Exit(1)
		}
	} else {
		fmt.Fprintf(os.Stderr, usage)
		os.Exit(1)
	}

}
