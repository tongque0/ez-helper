// 描述： 智能命令实现
// 作用： 调用大模型，生成EZ各个功能命令
package commands

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/tongque0/ez-helper/internel/model"
)

func ExecuteGenCommand(args []string) {

	userInput := strings.Join(args, " ")

	openaiStream := model.OpenAI(userInput)
	defer openaiStream.Close()
	for {
		response, err := openaiStream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println()
			return
		}

		if err != nil {
			fmt.Printf("\n生成EZ命令出错: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
