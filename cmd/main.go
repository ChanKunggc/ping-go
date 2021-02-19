package main

import "github.com/ChanKunggc/ping-go/cmd/ping"

func main() {

	//tcpPat.UnmarshalBinary()
	ping.RootCmd.Execute()
	//
	//srcAddr := &net.TCPAddr{
	//	IP:   []byte{192, 168, 57, 102},
	//	Port: 8888,
	//}
	//destAddr := &net.TCPAddr{
	//	IP:   []byte{192, 168, 110, 106},
	//	Port: 27017,
	//}
	//conn, err := net.DialTCP("tcp", srcAddr, destAddr)
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(0)
	//}
	//defer conn.Close()
	//fmt.Println("asdds")
	////_, err = conn.Write([]byte("asdasdasd"))
	////if err != nil {
	////	fmt.Println(err)
	////}
	//asd := make([]byte, 512)
	//length, err := conn.Read(asd)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(asd[:length])
}
