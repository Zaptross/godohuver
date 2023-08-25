package godohuver

import (
	"github.com/zaptross/godohuver/internal/dockerhub"
	"github.com/zaptross/godohuver/internal/utils"
)

type DockerHubImage = dockerhub.DockerHubImage

func GetTagsFromRepository(repository string, count int) ([]DockerHubImage, error) {
	return dockerhub.GetTagsFromRepository(repository, count)
}

func GetLatestImage(repository string) (DockerHubImage, error) {
	return dockerhub.GetLatestImage(repository)
}

func ExtractSemver(version string) (string, error) {
	return utils.ExtractSemver(version)
}
