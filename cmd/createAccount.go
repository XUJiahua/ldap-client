package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/xujiahua/ldap-client/pkg/ldap"
)

// createAccountCmd represents the createAccount command
var createAccountCmd = &cobra.Command{
	Use:   "createAccount",
	Short: "create account in ldap",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("email is required")
			os.Exit(1)
		}
		client, err := ldap.NewClient(ldapURL, ldapAdminUser, ldapAdminPassword, peopleDN, "")
		handleErr(err)
		err = client.CreateAccount(args[0])
		handleErr(err)

		newPassword, err := client.ResetPasswordByEmail(args[0])
		handleErr(err)

		cn, _ := ldap.EmailToCN(args[0])
		fmt.Println("login user: " + cn)
		fmt.Println("login password: " + newPassword)
	},
}

func init() {
	rootCmd.AddCommand(createAccountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createAccountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createAccountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
