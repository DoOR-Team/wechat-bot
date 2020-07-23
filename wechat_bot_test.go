package wechat_bot

import "testing"

func TestSendMsg(t *testing.T) {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=bb18509f-2324-47dc-8672-64736b11ab8f"

	msgContent := MsgContent{
		Msgtype: "markdown",
		Markdown: Markdown{
			Content: `实时新增用户反馈<font color="warning">132例</font>，请相关同事注意。
			>类型:<font color="comment">用户反馈</font>
			>普通用户反馈:<font color="comment">117例</font>
			>VIP用户反馈:<font color="comment">15例</font>`,
		},
	}
	SendMsg(url, msgContent)
}
