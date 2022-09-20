// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

type GlobalOptionsType struct {
	TrustServerCertificate bool
	DatabaseName           string
	UseTrustedConnection   bool
	UserName               string
	Endpoint               string
	AuthenticationMethod   string
	UseAad                 bool
	PacketSize             int
	LoginTimeout           int
	WorkstationName        string
	ApplicationIntent      string
	Encrypt                string
	DriverLogLevel         int
}

var GlobalOptions = &GlobalOptionsType{}

func init() {
	//rootCmd.Flags().Bool()
	rootCmd.PersistentFlags().BoolVarP(&GlobalOptions.TrustServerCertificate, "trust-server-certificate", "C", false, "Whether to trust the certificate presented by the endpoint for encryptoin")
	rootCmd.PersistentFlags().StringVarP(&GlobalOptions.DatabaseName, "database-name", "d", "", "The initial database for the connection")
	rootCmd.PersistentFlags().BoolVarP(&GlobalOptions.UseTrustedConnection, "use-trusted-connection", "E", false, "Whether to use integrated security")
}