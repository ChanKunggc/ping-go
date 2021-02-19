package internal

import (
	"fmt"
	"github.com/ChanKunggc/ping-go/pkg/logger"
)

type IPPayload struct {
	VersionHeadLen uint8
	Length         uint16
	TTL            uint8
	Protocol       uint8
	SrcAddr        [4]byte
	DesAddr        [4]byte
	Data           []byte
}

func parseIPPacket(respb []byte) (string, []byte) {
	version := respb[0] >> 4
	headerLength := respb[0] & 0x1f
	length := int(respb[2]*0xff + respb[3])
	TTL := respb[8] & 0xff
	var protocol string
	switch respb[9] {
	case 1:
		protocol = "ICMP"
		break
	case 6:
		protocol = "TCP"
	case 17:
		protocol = "UDP"
	}
	var srcIP = fmt.Sprintf("%d.%d.%d.%d", respb[12], respb[13], respb[14], respb[15])
	var destIP = fmt.Sprintf("%d.%d.%d.%d", respb[16], respb[17], respb[18], respb[19])
	logger.Debugf("版本: ipv%d ,ip报文头长度: %d ,总长度: %d ,存活时间: %d ,源地址: %s ,目标地址: %s ", version, headerLength*4, length, TTL, srcIP, destIP)
	return fmt.Sprintf("协议: %s ,存活时间: %d ", protocol, TTL), respb[headerLength*4:]
}
func (ip IPPayload) UnmarshalBinary() []byte {
	header := make([]byte, 20)
	header[0] = ip.VersionHeadLen
	header[2] = byte(ip.Length >> 8)
	header[3] = byte(ip.Length)
	header[8] = ip.TTL
	header[9] = ip.Protocol
	copy(header[11:15], ip.SrcAddr[:])
	copy(header[15:19], ip.SrcAddr[:])
	data := append(header, ip.Data...)
	checkSum(data)
	return data
}
func (ip IPPayload) checkSum(msg []byte) uint16 {
	sum, length := 0, len(msg)
	for i := 0; i < length-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if length%2 == 1 {
		sum += int(msg[length-1]) * 256 // notice here, why *256?
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16
	var answer = uint16(^sum)
	msg[10] = byte(answer >> 8)
	msg[11] = byte(answer & 0xff)
	return answer
}
