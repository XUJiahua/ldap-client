/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xujiahua/ldap-client/pkg/ldap"
)

// addUserCmd represents the addUser command
var addUserCmd = &cobra.Command{
	Use:   "addUser",
	Short: "add a user to a group, user email and group name are required",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("group name and user email are required")
			os.Exit(1)
		}
		client, err := ldap.NewClient(ldapURL, ldapAdminUser, ldapAdminPassword, peopleDN, groupDN)
		handleErr(err)
		err = client.AddUserToGroupEasy(args[0], args[1])
		handleErr(err)
		fmt.Printf("%s added to group %s\n", args[0], args[1])
	},
}

func init() {
	groupCmd.AddCommand(addUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
