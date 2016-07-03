package main

import (
	"fmt"
	"os"

	"github.com/cj123/goprowl"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var buildCommit string

type ProwlConfig struct {
	ProviderKey string `json:"provider_key"`
	APIKey      string `json:"apikey"`
}

func main() {
	fmt.Printf("Drone ProwlApp Plugin built from %s\n", buildCommit)

	var (
		repo  = new(drone.Repo)
		build = new(drone.Build)
		sys   = new(drone.System)
		vargs = new(ProwlConfig)
	)

	plugin.Param("build", build)
	plugin.Param("repo", repo)
	plugin.Param("system", sys)
	plugin.Param("vargs", vargs)

	err := plugin.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prowl := goprowl.NewProwlClient(vargs.ProviderKey)

	notification := goprowl.Notification{
		Application: "Drone CI Server",
		Description: fmt.Sprintf("%s: %s/%s#%s (%s) by %s",
			build.Status,
			repo.Owner,
			repo.Name,
			build.Commit[:8],
			build.Branch,
			build.Author,
		),
		URL:   fmt.Sprintf("%s/%s/%v", sys.Link, repo.FullName, build.Number),
		Event: "Build " + build.Status + " - " + repo.FullName,
	}

	notification.AddKey(vargs.APIKey)
	err = prowl.Push(notification)

	if err != nil {
		fmt.Printf("Unable to send prowl notification: %s\n", err.Error())
		os.Exit(1)
	}
}
