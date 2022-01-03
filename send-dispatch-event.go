package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/actions-go/toolkit/core"
)

var organization, repo, githubToken string
var httpClient = http.Client{
	Transport: &http.Transport{MaxIdleConnsPerHost: 4},
	Timeout:   time.Second * 60,
}

func main() {
	organization, _ := core.GetInput("Organization")
	repo, _ = core.GetInput("Repository")
	githubToken, _ = core.GetInput("github_token")

	fmt.Println("org", organization)
	fmt.Println("repo", repo)
	fmt.Print("token ---", githubToken)
	makeDispatchRequest()

}

func makeDispatchRequest() {

	url := fmt.Sprintf("https://api.github.kyndryl.net/repos/%s/%s/dispatches", organization, repo)

	request, err := http.NewRequest("POST", url, nil)
	if err != nil {
		core.SetFailed("Error making request :" + err.Error())
	}

	request.Header.Add("Authorization", "Bearer "+githubToken)

	response, err := httpClient.Do(request)
	if err != nil {
		core.SetFailed("Error making request :" + err.Error())
	}

	log.Println(response.Status)
	if response.StatusCode != http.StatusOK {
		core.SetFailedf("Error dispatching event status [%s] ", response.Status)
	}
}
