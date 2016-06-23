// Copyright © 2016 John Nguyen <jtnguyen236@gmail.com>
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

package cmd

import (
	"fmt"

	"github.com/nii236/gopher-resume/models"
	"github.com/nii236/gopher-resume/process"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// processCmd represents the process command
var processCmd = &cobra.Command{
	Use:   "process",
	Short: "The CLI tool for gopher-resume",
	Long: `The CLI tool for gopher-resume

Pass in the various flags in order to generate your CV.
`,
	Run: func(cmd *cobra.Command, args []string) {
		resume := &models.Resume{BasicInformation: models.Basic{
			Name:    viper.GetString("Name"),
			Label:   viper.GetString("Label"),
			Picture: viper.GetString("Picture"),
			Email:   viper.GetString("Email"),
			Phone:   viper.GetString("Phone"),
			Website: viper.GetString("Website"),
			Summary: viper.GetString("Summary"),
		},
		}
		res := process.ProcessCV(resume)
		fmt.Println(res.String())
	},
}

func init() {
	RootCmd.AddCommand(processCmd)
	processCmd.Flags().StringP("name", "n", "John Citizen", "Name of the candidate")
	processCmd.Flags().StringP("Label", "l", "A label", "The Label of the candidate")
	processCmd.Flags().StringP("Picture", "a", "A great avatar", "The Picture of the candidate")
	processCmd.Flags().StringP("Email", "e", "john@citizen.com", "The Email of the candidate")
	processCmd.Flags().StringP("Phone", "p", "123 456 789", "The Phone of the candidate")
	processCmd.Flags().StringP("Website", "w", "www.google.com", "The Website of the candidate")
	processCmd.Flags().StringP("Summary", "s", "John is a great guy!", "The Summary of the candidate")
	viper.BindPFlags(processCmd.Flags())
}
