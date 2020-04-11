// Copyright 2020 Red Hat
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gcloud

import (
	"github.com/spf13/cobra"

	"github.com/coreos/mantle/platform/api/gcloud"
)

var (
	cmdUpdate = &cobra.Command{
		Use:   "update-image --image=ImageName [--family=ImageFamily] [--description=ImageDescription]",
		Short: "Update os image",
		Long:  "Update os image attributes in GCP.",
		Run:   runUpdateImage,
	}

	updateImageName        string
	updateImageFamily      string
	updateImageDescription string
)

func init() {
	cmdUpdate.Flags().StringVar(
		&updateImageName, "image", "", "GCP image name")
	cmdUpdate.Flags().StringVar(
		&updateImageFamily, "family", "",
		"Updated GCP image family to attach image to")
	cmdUpdate.Flags().StringVar(
		&updateImageDescription, "description", "",
		"The updated description for the image")
	GCloud.AddCommand(cmdUpdate)
}

func runUpdateImage(cmd *cobra.Command, args []string) {
	// Check that the user provided an image
	if updateImageName == "" {
		plog.Fatal("Must provide an image name via --image")
	}
	// Check that the user provided at least one thing to change
	if updateImageFamily == "" && updateImageDescription == "" {
		plog.Fatal("Must provide one of --family or --description")
	}

	// Create the ImageSpec. Don't worry about passing "" for
	// family or description. If "" is passed no update will happen.
	spec := &gcloud.ImageSpec{
		Name:        updateImageName,
		Family:      updateImageFamily,
		Description: updateImageDescription,
	}

	_, err := api.UpdateImage(spec)
	if err != nil {
		plog.Fatalf("Updating image failed: %v\n", err)
	}
}
