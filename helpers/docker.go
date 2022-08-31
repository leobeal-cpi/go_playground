package helpers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"os/exec"
	"time"
)

func Docker() {
	clientWithOpts, _ := client.NewClientWithOpts(client.FromEnv)

	err := getContainers(clientWithOpts)

	// if docker is running, return
	if err != nil {
		openDocker(clientWithOpts)
		return
	}

	fmt.Println("Docker is already running")

	return
}

func getContainers(client *client.Client) error {

	// attempt to start docker
	_, err := client.ContainerList(context.Background(), types.ContainerListOptions{})

	if err != nil {
		fmt.Println("Docker is not running")
		return err
	}

	return nil
}

func openDocker(client *client.Client) {
	// attempt to start docker
	_, _ = exec.Command("open", "/Applications/Docker.app").Output()

	open := false

	fmt.Println("Waiting for Docker to start")
	for open == false {
		time.Sleep(1 * time.Second)
		fmt.Print(".")
		_, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
		if err == nil {
			open = true
			fmt.Println(".")
			fmt.Println("Docker has started")
		}
	}
}
