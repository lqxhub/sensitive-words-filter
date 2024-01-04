package sensitive_word_filter_filter

import (
	"math/rand"
	"testing"
)

const words_rule_file = "./example/rule.txt"

func TestSensitiveWordManager_Check(t *testing.T) {
	dfaManager, err := InitSensitiveWordWithFile(words_rule_file)
	if err != nil {
		panic(err)
	}

	//要屏蔽的词
	strs := []string{
		"粿聊",
		"含屌",
		"含~~~~~~~~~~~~~~进",
		"含--着",
		"漢奸",
		"豪乳",
		"豪&&乳",
		"好---紧",
		"好嫩",
		"好!!嫩",
		"黑111111111111$$$$$***1111111屄",
		"黑逼",
		"后庭",
		"后--庭",
		"狐@#@@$$媚",
		"互fsdfsf12#$%^&*()淫",
		"bz176.com",
	}

	for _, str := range strs {
		if !dfaManager.HasSensitiveWords(str) {
			t.Fatalf("dfaManager.Check(%s) failed", str)
		}
	}

	//不屏蔽的词
	strs = []string{
		"你好",
		"早上好",
		"打副本",
		"加我好友",
		"文档",
		"百度",
		"特斯拉",
		"马斯克",
	}

	for _, str := range strs {
		if dfaManager.HasSensitiveWords(str) {
			t.Fatalf("dfaManager.Check(%s) failed", str)
		}
	}

}

func BenchmarkDfaManager_Check(b *testing.B) {
	dfaManager, err := InitSensitiveWordWithFile(words_rule_file)
	if err != nil {
		panic(err)
	}
	strs := []string{
		"粿聊",
		"含屌",
		"含~~~~~~~~~~~~~~进",
		"含--着",
		"漢奸",
		"豪乳",
		"豪&&乳",
		"好---紧",
		"好嫩",
		"好!!嫩",
		"黑111111111111$$$$$***1111111屄",
		"黑逼",
		"后庭",
		"后--庭",
		"狐@#@@$$媚",
		"互fsdfsf12#$%^&*()淫",
		"bz176.com",
	}
	n := len(strs)
	for i := 0; i < b.N; i++ {
		dfaManager.HasSensitiveWords(strs[rand.Intn(n)])
	}
}
