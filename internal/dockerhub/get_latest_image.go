package dockerhub

import (
	"fmt"
	"regexp"
	"sort"
	"time"

	"github.com/samber/lo"
)

var (
	ErrNoImages = fmt.Errorf("no images found")

	nilDockerHubImage = DockerHubImage{}
	semverRegex       = regexp.MustCompile(`^v?\d+\.\d+\.\d+$`)
)

func GetLatestImage(repository string) (DockerHubImage, error) {
	if err := validateRepository(repository); err != nil {
		return nilDockerHubImage, err
	}

	url := getTagsUrl(repository, 25) // try to find both latest and semver tags

	rawResults, err := getRawResults(url)
	if err != nil {
		return nilDockerHubImage, err
	}

	dhrr, err := UnmarshalDockerHubGetTagsResponse(rawResults)
	if err != nil {
		return nilDockerHubImage, err
	}

	if len(dhrr.Results) == 0 {
		return nilDockerHubImage, ErrNoImages
	}

	sort.Slice(dhrr.Results, func(i, j int) bool {
		iTime, err := time.Parse(time.RFC3339, dhrr.Results[i].TagLastPushed)

		if err != nil {
			return false
		}

		jTime, err := time.Parse(time.RFC3339, dhrr.Results[j].TagLastPushed)

		if err != nil {
			return false
		}

		return iTime.After(jTime)
	})

	latest, okLatest := lo.Find(dhrr.Results, func(res Result) bool {
		return res.Name == "latest"
	})

	if !okLatest {
		latest := dhrr.Results[0]
		for _, res := range dhrr.Results {
			resTime, err := time.Parse(time.RFC3339, res.TagLastPushed)

			if err != nil {
				continue
			}

			latestTime, err := time.Parse(time.RFC3339, latest.TagLastPushed)

			if err != nil {
				continue
			}

			if resTime.After(latestTime) {
				latest = res
			}
		}
	}

	// find the latest semver that matches the latest tag
	semver, okSemver := lo.Find(dhrr.Results, func(res Result) bool {
		return semverRegex.MatchString(res.Name) && latest.Digest == res.Digest
	})

	if okSemver {
		return DockerHubImage{
			Name:       repository,
			Tag:        semver.Name,
			LastPushed: semver.TagLastPushed,
		}, nil
	}

	return nilDockerHubImage, ErrNoImages
}
