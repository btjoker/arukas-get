# arukas-get

一个arukas的辅助小工具。

不依赖其他包，克隆到本地后修改ID，Token，Secret(或端口)

构建后，丢到 Shadowsocks 根目录，然后为这个执行文件创建一个桌面快捷方式。

你先需要获取arukas的 Token和Secret [传送门](https://app.arukas.io/settings/api-keys)

***
* 安装有 `golang` 
`git clone https://github.com/btjoker/arukas-get`
然后将安装了 `Shadowsocks` 服务器的 `AppID`，申请的 `Token` 和 `Secret`。
需要填写的地方:

    var (
	    // ID arukas的AppID
	    ID = ""
	    // Token 
	    Token = ""
	    // Secret
	    Secret = ""
		// Port 如果未修改过不要改动
		Port = 8989.0
	)
保存后直接 `go build` 编译好，丢到 Shadowsocks 根目录，然后为这个执行文件创建一个桌面快捷方式。运行。 


*  无编译环境：
	下载 `releases` 中的压缩包， 解压缩到 `Shadowsocks` 根目录。填写`apikey.txt` 的文件，	保存后。
	为 `arukas-get.exe` 创建一个桌面快捷方式。运行。 

填写格式：

    ID:11111111111111

    Token:1111111111111111111

    Secret:111111111111111111111111111111111111111


Note：
* 可自行修改，随意使用。免费测试结束后估计会砍不少免费功能。
* 使用本程序前， 请先备份原有的 `gui-config.json` 文件，本程序生成的文件会覆盖原有内容。
* http://acjoker.tk
***
LICENSE GPL3.0