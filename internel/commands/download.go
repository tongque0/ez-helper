package commands

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/urfave/cli/v2"
)

// ExecuteDownloadCommand 执行下载命令的函数
func ExecuteDownloadCommand(c *cli.Context) {
	if c.NArg() > 0 {
		downloadTarget := c.Args().Get(0)
		fmt.Printf("下载EZ: %s\n", downloadTarget)
		executeDownloadScript(downloadTarget)
	} else {
		fmt.Println("获取最新EZ版本信息...")
		var cmd *exec.Cmd

		if runtime.GOOS == "windows" {
			// Windows: 执行 PowerShell 脚本
			cmd = exec.Command("powershell", "-Command", `
				$OutputEncoding = [Console]::OutputEncoding = [Text.UTF8Encoding]::UTF8
				$owner = "m-sec-org"
				$repo = "EZ"
				$url = "https://api.github.com/repos/$owner/$repo/releases/latest"
				$response = Invoke-RestMethod -Uri $url -UseBasicParsing
				if ($response -ne $null) {
					Write-Host "最新版本信息："
					Write-Host "版本 Tag：" $response.tag_name
					Write-Host "版本名称：" $response.name
					Write-Host "发布时间：" $response.published_at
					Write-Host "下载链接："
					foreach ($asset in $response.assets) {
						Write-Host "文件名称：" $asset.name
						Write-Host "下载地址：" $asset.browser_download_url
					}
				} else {
					Write-Host "未能获取最新版本数据。"
				}
			`)
		} else {
			// Linux: 执行 Bash 脚本
			cmd = exec.Command("bash", "-c", `
				export LC_ALL=en_US.UTF-8
				owner="m-sec-org"
				repo="EZ"
				url="https://api.github.com/repos/$owner/$repo/releases/latest"
				response=$(curl -s $url)
				if [ -n "$response" ]; then
					echo "最新版本信息："
					echo "版本 Tag：" $(echo $response | jq -r '.tag_name')
					echo "版本名称：" $(echo $response | jq -r '.name')
					echo "发布时间：" $(echo $response | jq -r '.published_at')
					echo "下载链接："
					echo $response | jq -r '.assets[] | "文件名称：" + .name, "下载地址：" + .browser_download_url'
				else
					echo "未能获取最新版本数据。"
				fi
			`)
		}

		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("执行获取最新版本信息脚本时出错: %v\n", err)
			return
		}

		fmt.Printf("%s\n", output)
	}
}

// executeDownloadScript 使用 grab 库下载文件
func executeDownloadScript(target string) {
	// 从下载链接中提取文件名
	fileName := getFileNameFromURL(target)
	if fileName == "" {
		fmt.Println("无法解析文件名，请检查下载链接。")
		return
	}

	// 创建下载请求
	client := grab.NewClient()
	req, err := grab.NewRequest(fileName, target)
	if err != nil {
		fmt.Printf("创建下载请求失败: %v\n", err)
		return
	}

	fmt.Printf("开始下载: %s", target)

	// 开始下载
	resp := client.Do(req)

	// 打印下载进度
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Printf("\r  已下载 %.2f%% (%d/%d)",
				100*resp.Progress(),
				resp.BytesComplete(),
				resp.Size())
		case <-resp.Done:
			// 下载完成或出错
			if err := resp.Err(); err != nil {
				fmt.Printf("下载失败: %v\n", err)
			} else {
				fmt.Printf("下载完成: %s\n", resp.Filename)
			}
			return
		}
	}
}

// getFileNameFromURL 从 URL 中提取文件名
func getFileNameFromURL(url string) string {
	parts := strings.Split(url, "/")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}
