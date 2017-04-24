# arukas-get

一个arukas的辅助小工具。

不依赖其他包，克隆到本地后填好 ss密码，Token，Secret

构建后，丢到 Shadowsocks 根目录，然后为这个执行文件创建一个桌面快捷方式。

你先需要获取arukas的 Token 和 Secret [传送门](https://app.arukas.io/settings/api-keys)

***
* 安装有 `golang` 
`git clone https://github.com/btjoker/arukas-get`
然后将申请的 `Token` 和 `Secret` 和ss服务器中设置的 `密码`
需要填写的地方:

	const (
		// Token Token 1
		Token = ""
		// Secret Secret 1
		Secret = ""
		// Port ss设置的端口
		Port = 8989
		// Password ss的密码,
		Password = ``
	)


保存后直接 `go build` 编译好，丢到 Shadowsocks 根目录，然后为这个执行文件创建一个桌面快捷方式。运行。 


Note：
* 可自行修改，随意使用。免费测试结束后估计会砍不少免费功能。
* 使用本程序前， 请先备份原有的 `gui-config.json` 文件，本程序生成的文件会覆盖原有内容。
* http://acjoker.tk
***
LICENSE GPL3.0
