package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var ldapURL string
var peopleDN string
var groupDN string
var ldapAdminUser string
var ldapAdminPassword string
var currentUser string // e.g., john.xu@example.com

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ldap-client",
	Short: "a simple ldap client to create account and reset account password",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	viper.AutomaticEnv() // read in environment variables that match
	ldapURL = viper.GetString("LDAP_URL")
	ldapAdminUser = viper.GetString("LDAP_USER")
	ldapAdminPassword = viper.GetString("LDAP_PASSWORD")
	peopleDN = viper.GetString("PEOPLE_DN")
	groupDN = viper.GetString("GROUP_DN")
	currentUser = viper.GetString("CURRENT_USER")
	//rootCmd.PersistentFlags().StringVar(&ldapURL, "ldap-url", viper.GetString("LDAP_URL"), "ldap url")
	//rootCmd.PersistentFlags().StringVar(&ldapAdminUser, "ldap-user", viper.GetString("LDAP_USER"), "ldap user")
	//rootCmd.PersistentFlags().StringVar(&ldapAdminPassword, "ldap-password", viper.GetString("LDAP_PASSWORD"), "ldap password")
	//rootCmd.PersistentFlags().StringVar(&peopleDN, "people-dn", viper.GetString("PEOPLE_DN"), "ldap people dn")
}

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}
