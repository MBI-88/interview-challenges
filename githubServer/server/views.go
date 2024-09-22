package server

import (
	"encoding/json"
	"log"
	"net/http"
	"html/template"
)

//Api
const api = "https://api.github.com/search/issues?q=repo:golang/go"



// Get method. Get makes a query to Github repo once
func (g *GitModel) Get(w http.ResponseWriter, r *http.Request) {
	req, _:= http.NewRequest("GET",api,nil)
	req.Header.Set("Accept","application/json")
	resp,err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Error in get request %s",err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error in status code with code %d",resp.StatusCode)
	}
	if err := json.NewDecoder(resp.Body).Decode(g); err != nil {
		log.Fatalf("Error in json decoder %s",err)
	}
	if err := g.Layout(w); err != nil {
		log.Fatalf("Error in layout parser %s",err)
	}
	
	defer resp.Body.Close()
}

// Layout method. Make layout tamplate
func (g GitModel) Layout(w http.ResponseWriter) error{
	issuesTemp, err := template.ParseFiles("server/GithubIssues.html")
	if err != nil {return err}
   if er :=	issuesTemp.Execute(w,g); er != nil {
	return er
   }
	return nil
}