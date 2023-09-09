/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// dogCmd represents the dog command
var dogCmd = &cobra.Command{
	Use:   "dog",
	Short: "犬が癒してくれます",
	Long:  `指定した回数だけワンと可愛い犬が鳴き癒してくれます`,
	RunE: func(cmd *cobra.Command, args []string) error {
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			return err
		}
		for i := 0; i < count; i++ {
			cmd.Println("꒰ ՞•ﻌ•՞ ꒱ わん！")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dogCmd)
	dogCmd.Flags().Int("count", 1, "count of cat")
}
