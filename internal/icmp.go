package internal

import (
	"github.com/ChanKunggc/ping-go/pkg/logger"
	"os"
)

type ICMPPayload struct {
	DestAddr string
	Data     string
}

func (i ICMPPayload) UnmarshalBinary() []byte {
	data := make([]byte, 8)
	//	set ping type
	data[0] = 8
	data = append(data, []byte(i.Data)...)
	_, err := checkSum(data)
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}
	return data
}
func checkSum(msg []byte) (uint16, error) {
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
	msg[2] = byte(answer >> 8)
	msg[3] = byte(answer & 0xff)
	return answer, nil
}
