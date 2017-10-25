// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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

	"github.com/spf13/cobra"
)

// delmCmd represents the delm command
var delmCmd = &cobra.Command{
	Use:   "delm",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		meeting,_:=cmd.Flags().GetString("title")
		fmt.Println("meeting " + meeting + " will be deleted")
	},
}

func init() {
	RootCmd.AddCommand(delmCmd)

	// Here you will define your flags and configuration settings.

	delmCmd.Flags().StringP("title", "t", "Anonymous", "Help message for meeting")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
