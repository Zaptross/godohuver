package godohuver

import "github.com/zaptross/godohuver/internal/dockerhub"

type DockerHubImage = dockerhub.DockerHubImage

func GetTagsFromRepository(repository string, count int) ([]DockerHubImage, error) {
	return dockerhub.GetTagsFromRepository(repository, count)
}

func GetLatestImage(repository string) (DockerHubImage, error) {
	return dockerhub.GetLatestImage(repository)
}
