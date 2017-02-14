# arukas-get

一个arukas的辅助小工具。

不依赖其他包，克隆到本地后修改ID，Token，Secret(或端口)

构建后，丢到 Shadowsocks 根目录，然后为这个执行文件创建一个桌面快捷方式。

`git clone https://github.com/btjoker/arukas-get`

你需要获取arukas的apikey [传送门](https://app.arukas.io/settings/api-keys)

然后将安装了 Shadowsocks 的 `APPID`，申请的 `Token` 和 `Secret`。

替换掉 `main.go` 里的该字段，对于梯子的默认寻找端口是 `8989` 端口，如果你自己设定的不一样请修改 `main.go` 中第 `85` 行

`if i.(map[string]interface{})["container_port"] == 8989.0`

其中 `8989.0` 就是你要修改的地方，因为获取到的数值是 `float64` 类型的数值，
所以不要丢失后面的 `.0`.

需要修改的地方:

    const (
	    // ID arukas的APPID
	    ID = "11111111111111"
	    // Token 
	    Token = "1111111111111111111"
	    // Secret
	    Secret = "111111111111111111111111111111111111111"
    )


Node：
* 可自行修改，随意使用。免费测试结束后估计会砍不少免费功能。