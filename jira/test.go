package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"net/http"
	"time"
)

func main() {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: false,
	}
	client := &http.Client{Transport: tr}
	jiraClient, err := jira.NewClient(client, "http://10.1.2.220:8080/")
	if err != nil {
		panic(err)
	}
	//res, err := jiraClient.Authentication.AcquireSessionCookie
	res, err := jiraClient.Authentication.AcquireSessionCookie("weizhen.wang", "wwz921203")
	if err != nil || res == false {
		fmt.Printf("Result: %v\n", res)
		panic(err)
	}
	fmt.Println(jiraClient.Authentication.Authenticated())

	issue, _, err := jiraClient.Issue.Get("DTSA-40", nil)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
	fmt.Printf("Description: %s\n", issue.Fields.Description)
	for IssueLink := range(issue.Fields.IssueLinks){
		fmt.Printf("IssueLinks: %s\n",IssueLink)	
	}
	

	//	issue, _, err := jiraClient.Issue.Get("DTSA-45", nil)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)

	// MESOS-3325: Running mesos-slave@0.23 in a container causes slave to be lost after a restart
	// Type: Bug
	// Priority: Critical
}
