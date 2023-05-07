/*
Copyright © 2023 Iryna Nazaruk <ira.nazaruk2011@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// TeleToken bot
	TeleToken = os.Getenv("TELE_TOKEN")
)

// kbotCmd represents the kbot command
var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started. ", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatal("Please check your token, it might be wrong ", err)
			return
		}

		kbot.Handle(telebot.OnText, func(telebotContext telebot.Context) error {
			logger(telebotContext)

			payload := telebotContext.Message().Payload

			switch payload {
			case "Hello":
				err = telebotContext.Send(fmt.Sprintf("Hi there, I'm your Kbot %s!", appVersion))
			case "GoodBye":
				err = telebotContext.Send("Thanks for using Kbot %s! See you soon!")
			case "How are u?":
				err = telebotContext.Send("I'm fine, and u?")
			case "Can you sing for me?":
				err = telebotContext.Send("No, I'm not a human and cant do it...")				
			}

			return err
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// kbotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// kbotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func logger(m telebot.Context) {
	payload := m.Message().Payload
	messageText := m.Text()

	log.Printf("[payload]: %s; [text]: %s;", payload, messageText)
}