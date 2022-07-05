package services

import (
	"context"
	"dashboard/app/internal/cli/docker"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/docker/docker/api/types"
)

func getContainer() []types.Container {
	ctx := context.Background()
	container := docker.Cli()
	containerLists, err := container.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		fmt.Println(error.Error(err))
	}

	return containerLists
}

func GetContainerNameList() []string {
	containerLists := getContainer()
	containerNameLists := make([]string, 0)
	for _, containerList := range containerLists {
		containerNameLists = append(containerNameLists, containerList.Names...)
	}
	return containerNameLists
}

func RunContainer(imageName string) {

}

func RunAllContainer() {

}

func RunContainerDaemon() {

}

func StopAllContainers() {

}

func RemoveContainer(containerID string) string {
	if len(containerID) == 0 {
		fmt.Println("容器ID不能为空!")
	}
	ctx := context.Background()
	containerCLI := docker.Cli()
	removedContainerMsg := containerCLI.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{})
	if removedContainerMsg != nil {
		fmt.Println(error.Error(removedContainerMsg))
	}
	return "删除成功!"
}

func RestartContainer(containerID string, timeout *time.Duration) string {
	if len(containerID) == 0 {
		fmt.Println("容器ID不能为空!")
	}
	ctx := context.Background()
	containerCLI := docker.Cli()
	restartContainerMsg := containerCLI.ContainerRestart(ctx, containerID, timeout)
	if restartContainerMsg != nil {
		fmt.Println(error.Error(restartContainerMsg))
	}
	return "重启成功!"
}

func LogsOnContainers(containersID string) {
	ctx := context.Background()
	containerCLI := docker.Cli()
	options := types.ContainerLogsOptions{ShowStdout: true}
	outfile, err := containerCLI.ContainerLogs(ctx, containersID, options)
	if err != nil {
		fmt.Println(error.Error(err))
	}
	io.Copy(os.Stdout, outfile)
}
