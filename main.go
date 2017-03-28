package main

import (
	"fmt"
	"os"

	"github.com/cj123/goprowl"
)

var buildCommit string

func main() {
	fmt.Printf("Drone ProwlApp Plugin built from %s\n", buildCommit)

	var (
		providerKey = os.Getenv("PLUGIN_PROVIDER_KEY")
		apiKey      = os.Getenv("PLUGIN_APIKEY")

		buildStatus  = os.Getenv("DRONE_BUILD_STATUS")
		repoOwner    = os.Getenv("DRONE_REPO_OWNER")
		repoName     = os.Getenv("DRONE_REPO_NAME")
		commit       = os.Getenv("DRONE_COMMIT_SHA")
		branch       = os.Getenv("DRONE_COMMIT_BRANCH")
		author       = os.Getenv("DRONE_COMMIT_AUTHOR")
		url          = os.Getenv("DRONE_BUILD_LINK")
		repoFullName = os.Getenv("DRONE_REPO")
	)

	prowl := goprowl.NewProwlClient(providerKey)

	notification := goprowl.Notification{
		Application: "Drone CI Server",
		Description: fmt.Sprintf("%s: %s/%s#%s (%s) by %s",
			buildStatus,
			repoOwner,
			repoName,
			commit[:8],
			branch,
			author,
		),
		URL:   url,
		Event: "Build " + buildStatus + " - " + repoFullName,
	}

	notification.AddKey(apiKey)
	err := prowl.Push(notification)

	if err != nil {
		fmt.Printf("Unable to send prowl notification: %s\n", err.Error())
		os.Exit(1)
	}
}
