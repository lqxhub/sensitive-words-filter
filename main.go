package main

import (
	"bufio"
	"container/list"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"unicode"
)

const HttpAddress = ":8081"

var dfaManager *DfaManager

func NewDfa() error {
	dfaManager = NewDfaManager()

	//read rule from txt file
	file, err := os.OpenFile("./rule.txt", os.O_RDWR, 0666)
	if err != nil {
		return fmt.Errorf("open file error:%w", err)
	}

	readFile := bufio.NewReader(file)

	l := list.New()
	for {
		if line, _, err := readFile.ReadLine(); err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("read file error:", err)
		} else {
			l.PushBack(strings.TrimFunc(string(line), func(r rune) bool {
				return !unicode.IsPrint(r)
			}))
		}
	}

	dfaManager.InitWordTree(l)

	return nil
}

func main() {
	err := NewDfa()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/check_word", checkWord)
	http.HandleFunc("/", http.FileServer(http.Dir("./static")).ServeHTTP)

	err = http.ListenAndServe(HttpAddress, nil)
	if err != nil {
		panic(err)
	}
}

type Resp struct {
	Code int    `json:"code"` //1.ok, 2.敏感词, 3.参数错误
	Msg  string `json:"msg"`
}

func checkWord(w http.ResponseWriter, r *http.Request) {
	word := r.URL.Query().Get("word")
	resp := Resp{
		Code: 1,
	}
	if len(word) == 0 {
		resp.Code = 3
		resp.Msg = "参数错误"
	} else {
		if dfaManager.Check(word) {
			resp.Code = 2
			resp.Msg = fmt.Sprintf("`%s` 有敏感词", word)
		} else {
			resp.Code = 1
			resp.Msg = fmt.Sprintf("`%s` OK", word)
		}
	}

	bytes, _ := json.Marshal(resp)
	_, _ = w.Write(bytes)
}
