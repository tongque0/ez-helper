// 描述: 大模型功能封装
// 作用: 调用openai的大模型功能
package model

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
	"github.com/tongque0/ez-helper/conf"
)

func OpenAI(userinput string) *openai.ChatCompletionStream {
	config := openai.DefaultConfig(conf.GetConf().Model.APIKey)
	config.BaseURL = conf.GetConf().Model.BaseURL

	c := openai.NewClientWithConfig(config)
	ctx := context.Background()

	req := openai.ChatCompletionRequest{
		Model:     openai.GPT4oMini,
		MaxTokens: 188,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: conf.GetConf().EZ_Helper.EZ_Helper_System + conf.GetConf().EZ_Helper.EZHelp,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: userinput,
			},
		},
		Stream: true,
	}
	stream, err := c.CreateChatCompletionStream(ctx, req)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return nil
	}
	return stream
}
