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
	"Agenda/entity"
	"github.com/spf13/cobra"
	"Agenda/logger"
	"io/ioutil"
)

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "--title=meeting --pr=participator --st=start_time --et=end_time",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		pr,_:= cmd.Flags().GetString("pr")
		st,_:=cmd.Flags().GetString("st")
		et,_:=cmd.Flags().GetString("et")
		entity.CreateMeeting(title, pr, st, et)
		bytes,_ := ioutil.ReadFile("./CurUser")
  		curuser := string(bytes)
		logger.Log("'" + curuser + "' called: cm, title: " + title + ", pr: " + pr + ", starttime: " + st + ", endtime: " + et)
	},
}

func init() {
	RootCmd.AddCommand(cmCmd)

	// Here you will define your flags and configuration settings.
	cmCmd.Flags().StringP("title", "t", "Anonymous", "Help message for title")
	cmCmd.Flags().StringP("pr", "p", "Anonymous", "Help message for participator")
	cmCmd.Flags().StringP("st", "s", "Anonymous", "Help message for starttime")
	cmCmd.Flags().StringP("et", "e", "Anonymous", "Help message for endtime")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cmCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cmCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}