package testCase

import (
	"fmt"

	"github.com/EasyGolang/goTools/mEmail"
	"github.com/EasyGolang/goTools/mEncrypt"
	"github.com/EasyGolang/goTools/mTime"
)

func TestEmail() {
	TStr := "测试邮件" + mEncrypt.GetUUID()

	err := mEmail.New(mEmail.Opt{
		Account:  "trade@mo7.cc",
		Password: "svaeJgMraNmdVFJu",
		To: []string{
			"meichangliang@outlook.com",
		},
		From:        "Hunter 测试 服务",
		Subject:     "这里是Subject",
		Port:        "587",
		Host:        "smtp.exmail.qq.com",
		TemplateStr: TStr,
		SendData: struct {
			SysTime string
			Message string
		}{
			SysTime: mTime.UnixFormat(mTime.GetUnix()),
			Message: "这里是一封邮件",
		},
	}).Send()

	fmt.Println("发送邮件", err)
}
