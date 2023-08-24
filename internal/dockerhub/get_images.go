package dockerhub

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type DockerHubImage struct {
	Name       string
	Tag        string
	LastPushed string
}

var (
	ErrInvalidCount      = fmt.Errorf("count must be greater than 0 and less than 100")
	ErrEmptyRepository   = fmt.Errorf("repository must not be empty")
	ErrInvalidRepository = fmt.Errorf("repository must be in the format of <user>/repo, <org>/<repo>, _/<repo>, or <repo> (for official images)")
)

func GetTagsFromRepository(repository string, count int) ([]DockerHubImage, error) {
	if err := validateRepository(repository); err != nil {
		return nil, err
	}
	if err := validateCount(count); err != nil {
		return nil, err
	}

	url := getTagsUrl(repository, count)

	rawResults, err := getRawResults(url)
	if err != nil {
		return nil, err
	}

	dhrr, err := UnmarshalDockerHubGetTagsResponse(rawResults)
	if err != nil {
		return nil, err
	}

	var images []DockerHubImage
	for _, result := range dhrr.Results {
		images = append(images, DockerHubImage{
			Name:       repository,
			Tag:        result.Name,
			LastPushed: result.TagLastPushed,
		})
	}

	return images, nil
}

func getTagsUrl(repository string, count int) string {
	split := strings.Split(repository, "/")

	// Official images don't have a user or org
	if len(split) == 1 {
		split = append([]string{"library"}, split...)
	}

	// Official images can be prefixed with an underscore
	if len(split) > 1 && split[0] == "_" {
		split[0] = "library"
	}

	repo := strings.Join(split, "/")

	return fmt.Sprintf("https://hub.docker.com/v2/repositories/%s/tags/?page_size=%d", repo, count)
}

func getRawResults(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
