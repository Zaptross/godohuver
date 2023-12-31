// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    dockerHubGetTagsResponse, err := UnmarshalDockerHubGetTagsResponse(bytes)
//    bytes, err = dockerHubGetTagsResponse.Marshal()

package dockerhub

import "encoding/json"

func UnmarshalDockerHubGetTagsResponse(data []byte) (DockerHubGetTagsResponse, error) {
	var r DockerHubGetTagsResponse
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *DockerHubGetTagsResponse) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type DockerHubGetTagsResponse struct {
	Count    int64       `json:"count"`
	Next     string      `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []Result    `json:"results"`
}

type Result struct {
	Creator             int64   `json:"creator"`
	ID                  int64   `json:"id"`
	Images              []Image `json:"images"`
	LastUpdated         string  `json:"last_updated"`
	LastUpdater         int64   `json:"last_updater"`
	LastUpdaterUsername string  `json:"last_updater_username"`
	Name                string  `json:"name"`
	Repository          int64   `json:"repository"`
	FullSize            int64   `json:"full_size"`
	V2                  bool    `json:"v2"`
	TagStatus           string  `json:"tag_status"`
	TagLastPulled       string  `json:"tag_last_pulled"`
	TagLastPushed       string  `json:"tag_last_pushed"`
	MediaType           string  `json:"media_type"`
	ContentType         string  `json:"content_type"`
	Digest              string  `json:"digest"`
}

type Image struct {
	Architecture string      `json:"architecture"`
	Features     string      `json:"features"`
	Variant      interface{} `json:"variant"`
	Digest       string      `json:"digest"`
	OS           string      `json:"os"`
	OSFeatures   string      `json:"os_features"`
	OSVersion    interface{} `json:"os_version"`
	Size         int64       `json:"size"`
	Status       string      `json:"status"`
	LastPulled   string      `json:"last_pulled"`
	LastPushed   string      `json:"last_pushed"`
}
