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

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "--un=UserName --pw=password --email=a@xxx.com --phone=xxxxxx",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("un")
		password, _ := cmd.Flags().GetString("pw")
		email, _ := cmd.Flags().GetString("email")
		phone, _ := cmd.Flags().GetString("phone")
		fmt.Println("register args : ", username, password, email, phone)
		if username == "" {
			fmt.Println("please input username")
		} else if password == "" {
			fmt.Println("please input password")
		} else if email == "" {
			fmt.Println("please input email")
		} else if phone == "" {
			fmt.Println("please input phone")
		} else {
			tmp := entity.Register(username, password, email, phone)
			if tmp == 0 {
				fmt.Println("register successfully")
			} else {
				fmt.Println("the username is not available")
			}
		}
		logger.Log(username + "register")
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)

	// Here you will define your flags and configuration settings.
	registerCmd.Flags().StringP("un", "u", "", "an unregisted username")
	registerCmd.Flags().StringP("pw", "p", "", "your password must have both number and character")
	registerCmd.Flags().StringP("email", "e", "", "a valid email address")
	registerCmd.Flags().StringP("phone", "c", "", "a valid phone number")
	
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// registerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// registerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
