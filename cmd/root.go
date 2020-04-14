package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/2fk/qybot/bot"
)

const (
	ENVWebHook = "WEIXIN_BOT_WEBHOOK"
)

var (
	hook  string
	debug bool
)

func Execute() {
	root := NewRoot()
	root.AddCommand(newTextCommand())
	root.AddCommand(newMarkdownCommand())
	root.AddCommand(newVersionCommand())

	if err := root.Execute(); err != nil {
		if debug {
			log.Fatalf("%+v\n", err)
		}
		log.Fatalf("%v\n", err)
	}
}

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:           bot.BuildAppName,
		Short:         "QiYe WeiXin Bot",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.PersistentFlags().StringVarP(&hook, "hook", "", "", "weixin webhook address")
	cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug")
	return cmd
}

func newTextCommand() *cobra.Command {

	var opt struct {
		mentionedList       []string
		mentionedMobileList []string
	}

	cmd := &cobra.Command{
		Use:   "text",
		Short: "Send text message",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("miss content")
			}

			content := args[0]
			a := bot.New(buildHook())
			return a.SendText(content, opt.mentionedList, opt.mentionedMobileList)
		},
	}

	flag := cmd.Flags()
	flag.StringSliceVarP(&opt.mentionedList, "list", "", nil, "mentioned list")
	flag.StringSliceVarP(&opt.mentionedMobileList, "mobile-list", "", nil, "mentioned mobile list")

	return cmd
}

func newMarkdownCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "markdown",
		Short: "Send markdown message",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("miss content")
			}

			content := args[0]
			a := bot.New(buildHook())
			return a.SendMarkdown(content)
		},
	}
	return cmd
}

func newVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show Version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("%s %s(%s %s)\n%s\n",
				bot.BuildAppName,
				bot.BuildVersion,
				bot.BuildCommitHash,
				bot.BuildTime,
				bot.BuildGoVersion,
			)
			return nil
		},
	}
	return cmd
}

func buildHook() string {
	if hook == "" {
		return os.Getenv(ENVWebHook)
	}
	return hook
}
