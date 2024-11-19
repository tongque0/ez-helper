package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
)

var (
	conf *Config
	once sync.Once
)

type Config struct {
	Model     Model     `yaml:"model"`
	EZ_Helper EZ_Helper `yaml:"ez_helper"`
}

type Model struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

type EZ_Helper struct {
	OS               string `yaml:"os"`
	EZHelp           string `yaml:"ez_help"`
	EZ_Helper_System string `yaml:"ez_helper_system"`
}

// GetConf returns a singleton configuration instance.
func GetConf() *Config {
	once.Do(func() {
		err := initConf()
		if err != nil {
			panic(fmt.Sprintf("初始化配置文件失败: %v", err))
		}
	})
	return conf
}

// initConf initializes the configuration by loading from a YAML file.
func initConf() error {
	welcome()
	wd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("获取当前工作目录失败: %w", err)
	}

	// 构建配置文件路径
	confFileRelPath := filepath.Join(wd, "ez_helper.yaml")
	content, err := os.ReadFile(confFileRelPath)
	if err != nil {
		if os.IsNotExist(err) {
			// 如果配置文件不存在，则创建默认配置文件
			fmt.Println("加载配置文件失败，创建默认配置文件...")
			err = createDefaultConfig(confFileRelPath)
			if err != nil {
				return fmt.Errorf("创建默认配置文件失败: %w", err)
			}
			fmt.Println("默认配置文件创建成功！")
			// 重新读取配置文件
			content, err = os.ReadFile(confFileRelPath)
			if err != nil {
				return fmt.Errorf("读取配置文件失败: %w", err)
			}
		} else {
			return fmt.Errorf("读取配置文件失败: %w", err)
		}
	}

	conf = new(Config)
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		return fmt.Errorf("解析 YAML 失败: %w", err)
	}

	if err := validator.Validate(conf); err != nil {
		return fmt.Errorf("验证配置文件失败: %w", err)
	}

	return nil
}

// createDefaultConfig 创建默认配置文件
func createDefaultConfig(filePath string) error {
	defaultConfig := Config{
		Model: Model{
			APIKey:  "sk-5sH3SdNxuzQRS4sHZZU5JlCEuL6bW14pC1LkIPFMFYjo1nRX",
			BaseURL: "https://prime.zetatechs.com/v1",
		},
		EZ_Helper: EZ_Helper{
			OS:               runtime.GOOS,
			EZHelp:           ez_help,
			EZ_Helper_System: ez_helper_system,
		},
	}

	content, err := yaml.Marshal(&defaultConfig)
	if err != nil {
		return fmt.Errorf("序列化默认配置失败: %w", err)
	}

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		return fmt.Errorf("写入默认配置文件失败: %w", err)
	}

	return nil
}

func welcome() {
	asciiArt := `


███████╗███████╗     ██╗  ██╗███████╗██╗     ██████╗ ███████╗██████╗
██╔════╝╚══███╔╝     ██║  ██║██╔════╝██║     ██╔══██╗██╔════╝██╔══██╗
█████╗    ███╔╝█████╗███████║█████╗  ██║     ██████╔╝█████╗  ██████╔╝
██╔══╝   ███╔╝ ╚════╝██╔══██║██╔══╝  ██║     ██╔═══╝ ██╔══╝  ██╔══██╗
███████╗███████╗     ██║  ██║███████╗███████╗██║     ███████╗██║  ██║
╚══════╝╚══════╝     ╚═╝  ╚═╝╚══════╝╚══════╝╚═╝     ╚══════╝╚═╝  ╚═╝

`
	fmt.Print(asciiArt)
	fmt.Print(`
Usage:
	ez-helper> [options]

Options:
	-config <file>    指定配置文件，提供自定义参数用于任务的执行。例如：
	                  ez-helper -config config.yaml

	-install          智能安装模式，根据检测的环境自动安装所需工具。例如：
	                  ez-helper -install

	-help             显示帮助信息，列出所有可用选项。例如：
	                  ez-helper -help

	-version          显示当前版本号。例如：
	                  ez-helper -version

Default Mode:
	智能模式：无需任何选项，根据用户输入的内容自动生成适合的 ez 命令。例如：

Examples:
	1. 自动智能模式：
	    ez-helper> 扫描一下ocybers.com的子域名

	2. 使用配置文件运行：
	   ez-helper> config set apikey xxx

	3. 启动智能安装：
	   ez-helper> -install


`)
}

var ez_helper_system = `

	你是ez-helper,能够根据用户需求，通过ez的使用说明来生成命令合适命令，你的回答应该只包含推荐的命令，且仅推荐一个命令。

`

var ez_help = `
下面是ez的使用说明：
   NAME:
   ez - A powerful scanner engine

USAGE:
   ez [global options] command [command options] [arguments...]

DESCRIPTION:
   A powerful scanner engine

COMMANDS:
   webscan         Run a webscan task
   servicescan     Run a service scan task
   dnsscan         Run a dns scan task,gather subdomain
   brute           Run a brute service scan task
   reverse         Run a standalone reverse server
   web             Run a web server
   crawler         Run a crawler task
   machineid, mid  generate machineid
   exploit         exploit tool
   help, h         Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level value, -l value  Log level, choices are debug, info, warn, warning, success (default: "info")
   --config value, -c value     Load ez configuration from file (default: "config.yaml")
   --lic value                  ez license file (default: "ez.lic")
   --check-reverse              check reverse service is online,finish it will exit program (default: false)
   --help, -h                   show help

下面是ez各模块的使用说明:
NAME:
   ez dnsscan - Run a dns scan task,gather subdomain

USAGE:
   ez dnsscan [command options] [arguments...]

OPTIONS:
   --domain value, -d value             provide domain , separated by ',',like baidu.com,qq.com,163.com
   --disable-domain-api, --dda          disable gather subdomain by domain-api (default: false)
   --domain-file value, --df value      provide file path ,per line should be a domain on file
   --subdomain-file value, --sdf value  provide file path ,per line should be a subdomain on file, will resolver subdomain
   --brute-dict value, --bd value       brute domain dict path (default: built-in dict,around 9w+)
   --slow                               will analysis subdomain by all dns server,will more exactly (default: false, means analysis subdomain by a random dns server)
   --output value, -o value             subdomain result file (default: "subdomain.txt")
   --output-json value, --oj value      subdomain result to json file
   --help, -h                           show help

NAME:
   ez web - Run a web server

USAGE:
   you also can do passive web scan in this module, web-listen argument act on web,other arguments act on webscan,

OPTIONS:
   --web-listen value                  web listen ip:port  (default: "0.0.0.0:8888")
   --reset-web-password, --rwp         will reset web password (default: false)
   --safe-path value, --sp value       must visit the safe path first, unless set --no-safe-path (default: random path)
   --no-safe-path, --nsp               no need visit safe path if set (default: false)
   --pocs value                        specify the webpoc to run, separated by ',' (default: all pocs)
   --disable-pocs value                specify the webpoc to not run, separated by ',' (default: no pocs)
   --hosts value                       white hosts. specify the webpoc to run, separated by ',' like edu.cn,qq.com, if you use this ,will cover config.yaml's http.white_host (default: all hosts)
   --disable-hosts value               black hosts. specify the webpoc to run, separated by ',' like edu,gov, if you use this ,will cover config.yaml's http.black_host (default: no hosts)
   --level value                       0:low,1:medium,2:high,3:critical,the allow webpoc level >= you put level (default: 0,mean low,medium,high,critical)
   --listen value                      use proxy resource collector, value is proxy addr (default: "127.0.0.1:2222")
   --listen-auth value                 http listen auth,like user:pass
   --proxy value                       support http,socks5 ,like http://127.0.0.1:8080 or socks5://127.0.0.1:1080 (default: no proxy)
   --html-output value, --ho value     output result to FILE in HTML format (default: "result.html")
   --crawler-headless                  crawler with chrome headless, only act on --url and --urls-file, auto set around 5 min timeout  (default: false)
   --json-output value, --jo value     output result to FILE in json format
   --webhook-output value, --wo value  post ez result to url in json format
   --exploit                           run webshell payload (default: false)
   --no-force-finger, --nff            if don't find the website finger,will request the poc (default: false)
   --help, -h                          show help
`
