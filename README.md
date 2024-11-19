
# ez-helper

**ez-helper** 是一款智能工具，专为高效查找 [ez](https://docs.ezreal.cool/docs/intro/) 命令设计。基于 **chatgpt-40-mini** 实现，支持快速查询命令并提供 ez 最新版本的下载功能。

## 功能特点

- 智能查找：快速定位您需要的 ez 命令。
- 版本更新：轻松获取 ez 的最新版本下载地址。
- 经济高效：API 调用成本低，每次调用仅 ¥0.0014，初始预存余额为 ¥2（请勿恶意使用）。
- 能力增强: 使用更为强大的api，实现能力增强。
## 演示视频

https://github.com/user-attachments/assets/f96eaa99-6997-4664-b44f-cf24a5dd1ebf

## 使用方法

1. 下载并运行 `ez-helper.exe`。
2. 在命令行界面中输入您想要查找的 ez 命令。
3. 程序将调用 OpenAI 的 API，返回相关命令的详细信息。

## 安装

直接下载发布版本的可执行文件，解压后即可使用，无需额外配置。

## 自行编译

如果需要自行编译程序，可以按照以下步骤进行操作：

### 环境要求

- 安装 [Go](https://golang.org/) 编译环境。
- 安装 [UPX](https://upx.github.io/) 压缩工具（可选）。

### 编译命令

在 Windows 系统中运行以下命令：

```sh
go build -o ez-helper.exe .\cmd\main.go
upx.exe --best --lzma ez-helper.exe
```

## 特别说明

- 本程序调用的 API 为低价转接服务，每次调用费用为 **¥0.0014**。
- 默认预存余额为 **¥2**，仅供正常使用。**请勿进行恶意调用**。

## 许可证

本项目使用 [MIT 许可证](LICENSE)。有关更多信息，请参阅 LICENSE 文件。

