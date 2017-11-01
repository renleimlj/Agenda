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
	"Agenda/entity"
	"github.com/spf13/cobra"
	"Agenda/logger"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "--un=UserName --pw=password",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("un")
		password, _ := cmd.Flags().GetString("pw")
		fmt.Println("login args : ", username, password)
		if username == "" {
			fmt.Println("please input username")
		} else if password == "" {
			fmt.Println("please input password")
		} else {
			tmp := entity.Login(username, password)
			if tmp == 0 {
				fmt.Println("login successfully")
			} else {
				fmt.Println("please input correct username and password")
			}
		}
		logger.Log("'" + username + "' log in")
	},
}

var (
	username, password *string
)

func init() {
	RootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.
	username = loginCmd.Flags().StringP("un", "u", "", "agenda username")
	password = loginCmd.Flags().StringP("pw", "p","","agenda password")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}