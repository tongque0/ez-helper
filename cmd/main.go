package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/tongque0/ez-helper/conf"
	"github.com/tongque0/ez-helper/internel/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	// 实例化一个cli应用
	conf.GetConf()

	app := &cli.App{
		Name:  "EZ-helper",
		Usage: "一款ez智能终端助手",
		Commands: []*cli.Command{
			{
				Name:    "install",
				Aliases: []string{"i"},
				Usage:   "下载并安装所需工具",
				Action: func(c *cli.Context) error {
					commands.ExecuteDownloadCommand(c)
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			commands.ExecuteGenCommand(c.Args().Slice())
			return nil
		},
	}
	// 读取用户输入
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\033[1;32mez-helper>\033[0m ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "exit" {
			break
		}

		args := strings.Split(input, " ")
		err := app.Run(append([]string{os.Args[0]}, args...))
		if err != nil {
			fmt.Println(err)
		}
	}
}
