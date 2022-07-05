package services

import (
	"context"
	"dashboard/app/internal/cli/docker"
	"io"
	"os"

	"fmt"

	"github.com/docker/docker/api/types"
)

var cli = docker.Cli()

func ImagesID() []string {
	ctx := context.Background()
	imagesList, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		fmt.Println(error.Error(err))
	}
	imagesID := make([]string, 0)
	for _, image := range imagesList {
		imagesID = append(imagesID, image.ID)
	}
	return imagesID
}

func PullImages(imageName string) string {
	ctx := context.Background()
	images, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		fmt.Println(err)
	}
	defer images.Close()
	_, err = io.Copy(os.Stdout, images)

	return "创建成功"
}

func DelteImages(imageID string) []types.ImageDeleteResponseItem {
	ctx := context.Background()
	images := docker.Cli()

	deletedImage, err := images.ImageRemove(ctx, imageID, types.ImageRemoveOptions{})
	if err != nil {
		fmt.Println(error.Error(err))
	}
	return deletedImage
}
