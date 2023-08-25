package main

import (
	"math/rand"
	"testing"
)

func BenchmarkDfaManager_Check(b *testing.B) {
	_ = NewDfa()
	strs := []string{
		"你好你好你好你好,你!好你好你--好你好你好你好你好你好你好你好你////好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好你好***你好你好",
		"正红!@#@#$@$@$丹红丹红丹红丹常",
		"粿聊",
		"粿聊",
		"哈萨",
		"铪粉",
		"蛤蟆",
		"還權",
		"含屌",
		"含~~~~~~~~~~~~~~进含~~~~~~~~~~~~~~进含~~~~~~~~~~~~~~进含~~~~~~~~~~~~~~进含~~~~~~~~~~~~~~进",
		"含着",
		"含住",
		"韩正",
		"韓正",
		"汉刀",
		"漢奸",
		"豪乳",
		"好12333333333332313123紧",
		"好嫩",
		"河殇",
		"河殤",
		"荷治",
		"贺龙",
		"賀龍",
		"黑1111111111113343434$$$$$***1111111屄",
		"黑逼",
		"黑彩红丹红丹红丹",
		"黑卡红丹红丹红丹红丹",
		"嗨壶红丹红丹红丹",
		"狠肏红丹红丹红丹红丹",
		"红丹红丹红丹红丹红丹红丹",
		"红会",
		"红@#@#@#$$%#@$磷",
		"洪傳",
		"洪法",
		"洪兴",
		"洪興",
		"洪23/---@#RTYU***#$%^&*()吟",
		"洪志",
		"后庭",
		"后握",
		"豞粮",
		"狐@#@@$$fsdfsdfsf媚",
		"胡瘟",
		"虎门",
		"虎骑",
		"互fsdfsf12#$%^&*()淫",
		"护照",
	}
	n := len(strs)
	for i := 0; i < b.N; i++ {
		dfaManager.Check(strs[rand.Intn(n)])
	}
}
