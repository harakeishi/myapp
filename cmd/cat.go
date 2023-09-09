/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// catCmd represents the cat command
var catCmd = &cobra.Command{
	Use:   "cat",
	Short: "猫が癒してくれます",
	Long:  `指定した回数だけにゃーと可愛い猫が鳴き癒してくれます`,
	RunE: func(cmd *cobra.Command, args []string) error {
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			return err
		}
		for i := 0; i < count; i++ {
			cmd.Println("ฅ^•ω•^ฅ にゃー")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(catCmd)
	catCmd.Flags().Int("count", 1, "count of cat")
}
