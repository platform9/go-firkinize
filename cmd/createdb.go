// Copyright 2021 Platform9 Systems Inc.

package cmd

import (
    "github.com/spf13/cobra"
    "go.uber.org/zap"
)

func init() {
    createDBCmd.PersistentFlags().StringVar(&serviceName, "service-name", "", "Name of the service (example qbert)")
    createDBCmd.MarkFlagRequired("service-name")
    rootCmd.AddCommand(createDBCmd)
}

var createDBCmd = &cobra.Command{
    Use:   "create-db",
    Short: "Create DB for service",
    Long: `Create DB for service`,
    RunE: func(cmd *cobra.Command, args []string) error {
        cfgMgr := GetCfg()
        zap.L().Debug("Creating DB for service")
        err := cfgMgr.CreateDB(serviceName)
        if err != nil {
            zap.L().Info("Error creating DB")
            return err
        }
        return nil
    },
}
