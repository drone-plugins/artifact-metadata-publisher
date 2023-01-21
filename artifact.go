package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	fileArtifactV1 string = "fileUpload/v1"
)

type (
	File struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	Data struct {
		FileArtifacts []File `json:"fileArtifacts"`
	}
	FileArtifact struct {
		Kind string `json:"kind"`
		Data Data   `json:"data"`
	}
)

func writeArtifactFile(files []File, artifactFilePath string) error {
	artifact := FileArtifact{
		Kind: fileArtifactV1,
		Data: Data{
			FileArtifacts: files,
		},
	}

	b, err := json.MarshalIndent(artifact, "", "\t")
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to marshal output %+v", artifact))
	}

	dir := filepath.Dir(artifactFilePath)
	err = os.MkdirAll(dir, 0644)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to create %s directory for artifact file", dir))
	}

	err = ioutil.WriteFile(artifactFilePath, b, 0644)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to write artifact to artifact file %s", artifactFilePath))
	}
	return nil
}
