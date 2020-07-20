package unpack

import (
	"encoding/binary"
	"errors"
	"io"
)

const Msg_Header = "12345678"

// 编码
func Encode(bytesBuffer io.Writer, content string) error {
	// 消息格式 msg_header + content_length + content
	// 对应字节长度 header_length = 8 + content_length = 4 + 实际content_length
	// msg_header总共8个字节，我们就写入8个字节的数据，写到bytesBuffer中
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(Msg_Header)); err != nil {
		return err
	}
	// content_length是根据content实际大小去写入的，最长不让超过int32的长度
	clen := int32(len([]byte(content)))
	if err := binary.Write(bytesBuffer, binary.BigEndian, clen); err != nil {
		return err
	}
	// 最后写入content大小
	if err := binary.Write(bytesBuffer, binary.BigEndian, []byte(content)); err != nil {
		return err
	}
	return nil
}

// 解码
func Decode(bytesBuffer io.Reader) (bodyBuf []byte, err error) {
	// 先读取出header大小的bytes
	MagicBuf := make([]byte, len(Msg_Header))
	if _, err := io.ReadFull(bytesBuffer, MagicBuf); err != nil {
		return nil, err
	}
	// 比较读到的header和实际的header是否相同
	if string(MagicBuf) != Msg_Header {
		return nil, errors.New("msg_header error")
	}

	// 读取对应的4个字节的数据长度，到lengthBuf
	lengthBuf := make([]byte, 4)
	if _, err := io.ReadFull(bytesBuffer, lengthBuf); err != nil {
		return nil, err
	}
	// 读取完数据长度后，将lengthBuf转换为实际length
	// lengthBuf拿到的是二进制，经过大端字节器解码操作
	// 实际数据传输的时候，都是用大端去编码的
	length := binary.BigEndian.Uint32(lengthBuf)
	// 拿到实际数据的length后就可以去读取实际数据长度
	bodyBuf = make([]byte, length)
	if _, err := io.ReadFull(bytesBuffer, bodyBuf); err != nil {
		return nil, err
	}
	return bodyBuf, nil
}
