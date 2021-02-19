package internal

import (
	"fmt"
	"net"
)

type TCPPayload struct {
	SrcAddr    [4]byte
	DesAddr    [4]byte
	SrcPort    uint16
	DesPort    uint16
	SeqNumber  uint32
	AckNumber  uint32
	Flags      uint8
	WindowSize uint16
}


func (t TCPPayload) UnmarshalBinary() []byte {

	data := make([]byte, 20)
	data[0] = byte(t.SrcPort >> 8)
	data[1] = byte(t.SrcPort)
	data[2] = byte(t.DesPort >> 8)
	data[3] = byte(t.DesPort)
	data[4] = byte(t.SeqNumber >> 24)
	data[5] = byte(t.SeqNumber >> 16)
	data[6] = byte(t.SeqNumber >> 8)
	data[7] = byte(t.SeqNumber)
	data[8] = byte(t.AckNumber >> 24)
	data[9] = byte(t.AckNumber >> 24)
	data[10] = byte(t.AckNumber >> 24)
	data[11] = byte(t.AckNumber >> 24)
	data[12] = 5 << 4
	/*标志位*/
	data[13] = t.Flags
	data[14] = byte(t.WindowSize >> 8)
	data[15] = byte(t.WindowSize)
	t.checkSum(data)

	return data
}
func (t TCPPayload) checkSum(msg []byte) uint16 {
	sum, length := 0, len(msg)
	sum += int(t.SrcAddr[0])*256 + int(t.SrcAddr[1]) + int(t.SrcAddr[2])*256 + int(t.SrcAddr[3])
	sum += int(t.DesAddr[0])*256 + int(t.DesAddr[1]) + int(t.DesAddr[2])*256 + int(t.DesAddr[3])
	for i := 0; i < length-1; i += 2 {
		sum += int(msg[i])*256 + int(msg[i+1])
	}
	if length%2 == 1 {
		sum += int(msg[length-1]) * 256 // notice here, why *256?
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16
	var answer = uint16(^sum)
	msg[16] = byte(answer >> 8)
	msg[17] = byte(answer & 0xff)
	return answer
}
func asdasd() {
	srcAddr := &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: 8888,
	}
	destAddr := &net.TCPAddr{
		IP:   []byte{192, 168, 110, 227},
		Port: 8080,
	}
	conn, err := net.DialTCP("tcp", srcAddr, destAddr)
	if err != nil {
		fmt.Println(err)
	}
	conn.Read([]byte("asdasdasd"))
}
