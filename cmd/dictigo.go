package cmd

import (
	"fmt"
	"os"

	"github.com/Metudu/dictigo/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "dictigo",
	Short: "dictigo is a translated powered by DeepL",
	Long: `dictigo translates texts using DeepL API endpoint. You have to specify the target language first and the text latter.`,
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		language := args[0]
		text := args[1:]
		client.SendRequest(language, text)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}