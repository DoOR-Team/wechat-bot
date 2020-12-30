package wechat_bot

import (
	"fmt"
	"log"
	"testing"
)

func TestSendMsg(t *testing.T) {
	// url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=bb18509f-2324-47dc-8672-64736b11ab8f"
	url := ""

	baseUrl := "http://cdbot_be.hz-xuelang.xyz/"
	app := "test_deploy"
	env := "daily"
	dockerImage := "git.5c82513824cab00ef17467c4b6fd95375cdd87bb"

	msgContent := MsgContent{
		Msgtype: "markdown",
		Markdown: Markdown{
			Content: fmt.Sprintf("[发布测试环境](%s/deploy?app=%s&env=%s&dockerImage=%s)", baseUrl, app, env, dockerImage),
			// Content: "http://cdbot-be.hz-xuelang.xyz/deploy?app=test_deploy&env=daily&dockerImage=git.5c82513824cab00ef17467c4b6fd95375cdd87bb",
		},
	}
	log.Println(fmt.Sprintf("%s/deploy?app=%s&env=%s&dockerImage=%s", baseUrl, app, env, dockerImage))
	err := SendMsg(url, msgContent)
	log.Println(err)
}
