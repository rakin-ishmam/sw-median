// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
)

var fileLoc string
var wsize int

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "find median",
	Long:  `find median`,
	Run: func(cmd *cobra.Command, args []string) {
		calcMedian(fileLoc, wsize)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&fileLoc, "file", "f", "", "file location")
	runCmd.MarkFlagRequired("file")

	runCmd.Flags().IntVarP(&wsize, "window", "w", 3, "sliding window size")
}
