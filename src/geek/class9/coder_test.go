package class9

import (
	"fmt"
	"testing"
)

func TestCoder(t *testing.T) {
	msg := "hello world!"
	p := &Pkg{
		PLen: uint32(len(msg)) + 16,
		HLen: 16,
		Ver:  3,
		Op:   4,
		Seq:  100,
		Body: msg,
	}
	fmt.Printf("包数据： %+v\n", p)
	enmsg := Encoder(p)
	fmt.Printf("编码数据： %+v\n", enmsg)

	depkg, err := Decoder(enmsg)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("解码包： %+v\n", depkg)
	if msg != depkg.Body {
		t.Errorf("测试失败")
	}
}
