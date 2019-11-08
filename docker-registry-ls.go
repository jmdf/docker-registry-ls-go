package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RepoLists struct {
	RepoLists string `json:"repositories"`
}

func main() {

	// variables
	dockerRegistry := ""

	// Grab variable from user environment
	dockerRegistryEnv := os.Getenv("DOCKER_REGISTRY")

	// Command line options
	serverPtr := flag.String("s", "", "Server to query")
	repoPtr := flag.String("r", "", "List tags from specific repo")
	flag.Parse()

	// Do we have a server?
	if dockerRegistryEnv == "" && *serverPtr == "" {
		println("No server specified, use variable DOCKER_REGISTRY or '-s'")
		os.Exit(1)
	}

	// Prioritize command line input
	if *serverPtr != "" {
		dockerRegistry = *serverPtr
	} else {
		dockerRegistry = dockerRegistryEnv
	}

	if *repoPtr != "" {
		listRepoTags(dockerRegistry, *repoPtr)
	} else {
		listRepos(dockerRegistry)
	}
}

func listRepos(url string) {
	// Query registry
	response, err := http.Get(url + "/v2/_catalog")
	// Handle error, if needed
	if err != nil {
		panic(err)
	}
	// Process query response
	processedResponse, _ := ioutil.ReadAll(response.Body)

	// Print info
	fmt.Println(string(processedResponse))
}

func listTags(url string) {
	// Query registry
	response, err := http.Get(url + "/v2/_catalog")
	// Handle error, if needed
	if err != nil {
		panic(err)
	}
	// Process query response
	processedResponse, _ := ioutil.ReadAll(response.Body)

	var result RepoLists
	json.Unmarshal(processedResponse, &result)
	fmt.Println(result.RepoLists)
}

// Function to list tags from specific repo
func listRepoTags(url string, repoName string) {
	// Query registy
	response, err := http.Get(url + "/v2/" + repoName + "/tags/list")
	// Handle error, if needed
	if err != nil {
		panic(err)
	}
	// Process query response
	processedResponse, _ := ioutil.ReadAll(response.Body)

	// Print info
	fmt.Println(string(processedResponse))
}
