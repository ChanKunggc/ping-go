package internal

import (
	"fmt"
	"github.com/ChanKunggc/ping-go/pkg/logger"
	"net"
	"os"
	"time"
)

var Timeout time.Duration = 15
var ReConDelay time.Duration = 1
var ReConCont  = 10

type Payload interface {
	UnmarshalBinary() []byte
}

func Ping(tcpUse bool, desAddr, data string) {
	res := make([]string, 0)
	if tcpUse {
		go pingTCP(desAddr, res)
	} else {
		go pingICMP(desAddr, res)
	}
	time.Sleep(Timeout*time.Second)
	pl := 100
	if len(res) > 0 {
		pl -= len(res) * 10
	}
	fmt.Printf("发送数量: %d ,接收数量: %d ,丢包率: %d%s\n", ReConCont, len(res), pl, "%")
}
func pingTCP(desAddr string, res []string) {
	for i := 0; i < ReConCont; i++ {
		start := time.Now()
		conn, err := net.DialTimeout("tcp4", desAddr, Timeout*time.Second)
		if err != nil {
			logger.Error(err)
			continue
		}
		end := time.Now()
		delay := end.Sub(start)
		str := fmt.Sprintf("协议: %s ,序号=%d, 延迟: %s ", "TCP", i, delay)
		logger.Infof(str)
		res = append(res, str)
		conn.Close()
		time.Sleep(ReConDelay * time.Second)
	}
}
func pingICMP(desAddr string, res []string) {
	conn, err := net.DialTimeout("ip4:icmp", desAddr, Timeout*time.Second)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	defer conn.Close()
	icmpPkt := ICMPPayload{
		DestAddr: desAddr,
	}
	for i := 0; i < ReConCont; i++ {
		start := time.Now()
		_, err = conn.Write(icmpPkt.UnmarshalBinary())
		if err != nil {
			logger.Error(err)
		}
		resp := make([]byte, 512)
		respLen, err := conn.Read(resp)
		if err != nil {
			logger.Error(err)
			continue
		}
		end := time.Now()
		str, _ := parseIPPacket(resp[:respLen])
		fmt.Println()
		logger.Infof(str+",延迟: %s", end.Sub(start))
		res = append(res, str)
		time.Sleep(ReConDelay * time.Second)
	}
	return
}
