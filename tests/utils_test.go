package tests

import (
	"fmt"
	"testing"

	"github.com/BitTraceProject/BitTrace-Types/pkg/common"
)

func TestScanFileLines(t *testing.T) {
	var (
		fp = "./.bittrace/test_file/test.txt"
		n  = int64(0)
	)
	lines, eof, err := common.ScanFileLines(fp, n)
	for _, line := range lines {
		t.Log(len(line))
	}
	t.Log(eof, err)
}

func TestStructToJsonStr(t *testing.T) {

	//编程语言结构体
	type programLang struct {
		Name  string `json:"name"`            //转json的时候自定义字段的标签，可用json标签定义
		Other string `json:"tools,omitempty"` //转json的时候，自定义标签加omitempty，在此值为空的时候可忽略
	}

	//定义一个web项目结构体
	type webItem struct {
		Name string
		Db   string
		Env  string
		Lang []programLang
	}

	item := &webItem{
		Name: "Taobao",
		Db:   "MySQL",
		Env:  "Centos7",
		Lang: []programLang{
			{Name: "PHP", Other: "swoole,laravel"},
			{Name: "GO", Other: "docker,beego"},
			{Name: "python", Other: ""},
			{Name: "JAVA", Other: "Spring Boot"},
		},
	}

	fmt.Println(item)
	res := common.StructToJsonStr(item)
	fmt.Println(res)
}
