package utils

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
)

const (
	// 存放消息体长度的固定长度字节切片
	defaultMsgHeadLen = 4
	defaultBufSize = 8096
)

// transfer 将 Msg信息从用户提交的数据中读取出来，并将服务器端封装的MSG结构返回给客户端
type Transfer struct {
	// 连接
	Conn net.Conn
	// 读取数据用的缓冲
	buf []byte
	// 消息体长度转换成[]byte固定多长
	msgHeadLen int
}

// NewTransfer 获取一个控制
func NewTransfer() *Transfer {
	return &Transfer{
		Conn: Conn,
		buf: make([]byte, defaultBufSize),
		msgHeadLen: defaultMsgHeadLen,
	}
}

// ByteSliceToUint 字节切片转unit数据
func (tf *Transfer)ByteSliceToUint(data []byte) uint32 {
	return binary.BigEndian.Uint32(data)
}

// UintToByteSlice 将一个uint整数存放到固定长度的[]byte中
func (tf *Transfer)UintToByteSlice(num int) []byte{
	buf := tf.buf[:tf.msgHeadLen]
	binary.BigEndian.PutUint32(buf, uint32(num))
	return buf
}

// ReadData 从用户提交的数据中提取出MSG信息
func (tf *Transfer)ReadData() (msg *message.Msg, err error) {
	// 1. 获取消息体的长度
	n, err := tf.Conn.Read(tf.buf[:tf.msgHeadLen])
	if n != tf.msgHeadLen || err != nil {
		log.Println(n, tf.msgHeadLen)
		err = errors.New(fmt.Sprintf("消息长度丢包或err=%v", err))
		return
	}
	msgLen := int(tf.ByteSliceToUint(tf.buf[:tf.msgHeadLen]))

	// 2. 获取消息体
	buf := tf.buf[:msgLen]
	n, err = tf.Conn.Read(buf)
	if err != nil || n != msgLen{
		err = errors.New(fmt.Sprintf("消息体丢包或err=%v", err))
		return
	}

	// 3. 反序化消息体
	err = json.Unmarshal(buf, &msg)
	if err != nil {
		return
	}
	return
}

// WriteData 将服务器端返回的MSG结构返回给客户端
func (tf *Transfer)WriteData(msg *message.Msg) (err error) {
	// 1. 序列化消息体
	msgData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	// 2. 计算消息的长度，将长度存放入一个字节固定长度的字节切片中
	msgLen := len(msgData)
	msgLenSlice := tf.UintToByteSlice(msgLen)

	// 3. 发送消息长度并确认长度是否成功发送
	n, err := tf.Conn.Write(msgLenSlice)
	if err != nil || n != tf.msgHeadLen {
		err = errors.New(fmt.Sprintf("消息长度发包失败或err=%v", err))
		return
	}

	fmt.Println(string(msgData))
	// 4. 发送真正的消息体并确认消息是否成功发送
	n, err = tf.Conn.Write(msgData)
	if err != nil || n != msgLen{
		err = errors.New(fmt.Sprintf("消息内容发包失败或err=%v", err))
		return
	}
	return
}

// sendData 发送消息
func (tf *Transfer) SendData(data interface{}, dataType message.Kind) (err error) {
	d, err := json.Marshal(data)
	if err != nil {
		return
	}
	var msg message.Msg
	msg.Data = string(d)
	msg.Type = dataType
	err = tf.WriteData(&msg)
	if err != nil {
		return
	}
	return
}

// recvData 接收消息
func (tf *Transfer) RecvData() (data string, dataType message.Kind, err error) {
	recvData, err := tf.ReadData()
	if err != nil {
		return
	}
	data = recvData.Data
	dataType = recvData.Type
	return
}
