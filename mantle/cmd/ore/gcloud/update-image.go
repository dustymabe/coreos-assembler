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
	"fmt"

	"github.com/coreos/mantle/platform/api/gcloud"
	"github.com/spf13/cobra"
)

var (
	cmdUpdateImage = &cobra.Command{
		Use:   "update-image",
		Short: "Update os image",
		Long:  "Update os image attributes in GCP.",
		Run:   runUpdateImage,
	}

	updateImageDescription string
	updateImageFamily      string
	updateImageName        string
	updateImageReplacement string
	updateImageState       string
)

func init() {
	cmdUpdateImage.Flags().StringVar(
		&updateImageName, "image", "", "GCP image name")
	cmdUpdateImage.Flags().StringVar(
		&updateImageFamily, "family", "",
		"Updated GCP image family to attach image to")
	cmdUpdateImage.Flags().StringVar(
		&updateImageDescription, "description", "",
		"The updated description for the image")
	cmdUpdateImage.Flags().StringVar(&updateImageState, "state", "",
		fmt.Sprintf("Deprecation state must be one of: %s,%s,%s,%s",
			gcloud.DeprecationStateActive,
			gcloud.DeprecationStateDeprecated,
			gcloud.DeprecationStateObsolete,
			gcloud.DeprecationStateDeleted))
	cmdUpdateImage.Flags().StringVar(&updateImageReplacement,
		"replacement", "", "optional: link to replacement for the deprecated image")
	GCloud.AddCommand(cmdUpdateImage)
}

func runUpdateImage(cmd *cobra.Command, args []string) {
	// Check that the user provided an image
	if updateImageName == "" {
		plog.Fatal("Must provide an image name via --image")
	}
	// Check that the user provided at least one thing to change
	if updateImageFamily == "" && updateImageDescription == "" && updateImageState == "" {
		plog.Fatal("Must provide one of --family/--description/--state")
	}

    // Make call to UpdateImage. Don't worry about passing ""
    // If "" is passed no update will happen.
	pending, err := api.UpdateImage(
        updateImageName,
        updateImageFamily,
        updateImageDescription,
        updateImageState,
        updateImageReplacement,
    )
    if err == nil {
        err = pending.Wait()
    }
	if err != nil {
		plog.Fatalf("Updating image failed: %v\n", err)
	}
}
