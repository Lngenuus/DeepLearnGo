package class9

import (
	"encoding/binary"
	"errors"
)

// 注意大小端问题 计算机存储中是小端序 而网络传输中是符合人类识别的大端序

type Pkg struct {
	PLen uint32 // 4 bytes
	HLen uint16 // 2 bytes
	Ver  uint16 // 2bytes
	Op   uint32 // 4bytes
	Seq  uint32 // 4bytes
	Body string // body
}

// Decoder 解码
func Decoder(data []byte) (*Pkg, error) {
	if len(data) <= 16 {
		return nil, errors.New("data format error")
	}

	p := &Pkg{}
	p.PLen = binary.BigEndian.Uint32(data[:4])

	p.HLen = binary.BigEndian.Uint16(data[4:6])

	p.Ver = binary.BigEndian.Uint16(data[6:8])

	p.Op = binary.BigEndian.Uint32(data[8:12])

	p.Seq = binary.BigEndian.Uint32(data[12:16])

	p.Body = string(data[16:])
	return p, nil
}

// Encoder 编码
func Encoder(p *Pkg) []byte {
	ret := make([]byte, p.PLen)
	binary.BigEndian.PutUint32(ret[:4], p.PLen)
	binary.BigEndian.PutUint16(ret[4:6], p.HLen)
	binary.BigEndian.PutUint16(ret[6:8], p.Ver)
	binary.BigEndian.PutUint32(ret[8:12], p.Op)
	binary.BigEndian.PutUint32(ret[12:16], uint32(p.Seq))

	byteBody := []byte(p.Body)
	copy(ret[16:], byteBody)
	return ret
}
