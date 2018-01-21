package cron

import (
	"fmt"

	"strings"

	"github.com/open-falcon/alarm/g"
	"github.com/open-falcon/common/model"
	"github.com/open-falcon/common/utils"
)

func BuildCommonSMSContent(event *model.Event) string {
	link := g.Link(event)
	var status string
	if strings.ToUpper(event.Status) == "OK" {
		status = "恢复"
	} else {
		status = "故障"
	}

	return fmt.Sprintf(
		"状态：%s[P%d]\n"+
			"节点：%s\n"+
			"指标：%s\n"+
			"标签：%s\n"+
			"监控函数：%s %s%s%s\n"+
			"监控备注：%s\n"+
			"告警次数：%d/%d\n"+
			"告警时间：%s\n"+
			"告警链接：%s\n",
		status,
		event.Priority(),
		event.Endpoint,
		event.Metric(),
		utils.SortedTags(event.PushedTags),
		event.Func(),
		utils.ReadableFloat(event.LeftValue),
		event.Operator(),
		utils.ReadableFloat(event.RightValue()),
		event.Note(),
		event.CurrentStep,
		event.MaxStep(),
		event.FormattedTime(),
		link,
	)
}

func BuildCommonMailContent(event *model.Event) string {
	link := g.Link(event)
	var status string
	if strings.ToUpper(event.Status) == "OK" {
		status = "恢复"
	} else {
		status = "故障"
	}
	return fmt.Sprintf(
		"状态：%s[P%d]<br>"+
			"节点：%s<br>"+
			"指标：%s<br>"+
			"标签：%s<br>"+
			"监控函数：%s %s%s%s<br>"+
			"监控备注：%s<br>"+
			"告警次数：%d/%d\n"+
			"告警时间：%s<br>"+
			"告警链接：%s<br>",
		status,
		event.Priority(),
		event.Endpoint,
		event.Metric(),
		utils.SortedTags(event.PushedTags),
		event.Func(),
		utils.ReadableFloat(event.LeftValue),
		event.Operator(),
		utils.ReadableFloat(event.RightValue()),
		event.Note(),
		event.CurrentStep,
		event.MaxStep(),
		event.FormattedTime(),
		link,
	)
}

func GenerateSmsContent(event *model.Event) string {
	return BuildCommonSMSContent(event)
}

func GenerateMailContent(event *model.Event) string {
	return BuildCommonMailContent(event)
}
