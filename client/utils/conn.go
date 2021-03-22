package utils

import (
	"log"
	"net"
)

var (
	// 全局连接和全局读写数据接口
	Conn net.Conn
)

func init() {
	Conn = Dial("tcp", "127.0.0.1:8888")
}

// 客户端与服务器建立连接
func Dial(network string, address string) net.Conn {
	conn, err := net.Dial(network, address)
	if err != nil {
		log.Fatalln(err)
	}
	return conn
}