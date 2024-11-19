# 项目介绍
ez-helper 是一款用于查找 ez [https://docs.ezreal.cool/docs/intro/](https://docs.ezreal.cool/docs/intro/) 命令的智能程序。它调用了 OpenAI 的 API，能够方便地下载 ez。

# 演示视频
![演示视频](assets/show.mp4)


# 使用方法
1. 运行 `ez-helper.exe`。
2. 输入你想查找的 ez 命令。
3. 程序将通过 OpenAI 的 API 返回相关的 ez 命令信息。

# 依赖
- Go 语言
- OpenAI API
- UPX (用于压缩可执行文件)(非必须)

# 安装
请确保已安装 Go 语言和 UPX，然后运行以下指令进行打包：
```sh
go build -o ez-helper.exe .\cmd\main.go

upx.exe --best --lzma ez-helper.exe
```

# 许可证
本项目使用 MIT 许可证。详情请参阅 LICENSE 文件。
























