// Copyright 2020 Platform9 Systems Inc.

package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	getKeystoneCmd.PersistentFlags().StringVar(&serviceName, "service-name", "", "Name of the service (example qbert)")
	getKeystoneCmd.MarkPersistentFlagRequired("service-name")
	rootCmd.AddCommand(getKeystoneCmd)
}

var getKeystoneCmd = &cobra.Command{
	Use:   "get-keystone",
	Short: "Get Keystone related attributes",
	Long:  `Get Keystone related attributes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgMgr := GetCfg()
		zap.L().Debug("Get keystone password")
		_, err := cfgMgr.GetKeystonePassword(serviceName)
		if err != nil {
			zap.L().Info("Error getting keystone password")
			return err
		}
		// Do not print password
		// It causes a security alert
		// Clear-text logging of sensitive information
		// fmt.Println(password)
		return nil
	},
}
