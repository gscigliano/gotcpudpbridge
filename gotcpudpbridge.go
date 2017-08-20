package main

import (
 "fmt"
 "net"
 "os"
 "strconv"
)

func main() {
// argv basic checks
 argsWithoutProg := os.Args[1:]
 if len(argsWithoutProg) !=5 { printUsage() }

 udp_buf,_ := strconv.Atoi(argsWithoutProg[3])  ;  tcp_buf,_ := strconv.Atoi(argsWithoutProg[4])

 switch argsWithoutProg[0] {
  case "udp2tcp": udp2tcp(argsWithoutProg[1],argsWithoutProg[2],udp_buf,tcp_buf)
  case "tcp2udp": tcp2udp(argsWithoutProg[1],argsWithoutProg[2],udp_buf,tcp_buf)
  default: printUsage()
 }

}

func udp2tcp(src_conn string,dst_conn string, udp_buf_size int , tcp_buf_size int) {

// start listening on udp and setup tcp conn to the dest

 udp_add,err := net.ResolveUDPAddr("udp",src_conn)

 udp_sock, err := net.ListenUDP("udp", udp_add)
 if err != nil { fmt.Println("Error listening:", err) ; os.Exit(1) }

 tcp_sock, err := net.Dial("tcp", dst_conn)
 if err != nil {  fmt.Println("Error calling dest:", err.Error()) ; os.Exit(1) }

 udp_buf := make([]byte, udp_buf_size) ;  tcp_buf := make([]byte, tcp_buf_size)

for {
// loops over and bridge the data

 // udp into tcp
 udpReqLen,remoteaddr,_ := udp_sock.ReadFromUDP(udp_buf)

 fmt.Println("udp2tcp: got from UDP :", udpReqLen)                              // for debugging
 tcp_sock.Write([]byte(udp_buf[:udpReqLen]))                                            // dump data into tcp

 // tcp into udp
 tcpReqLen, _ := tcp_sock.Read(tcp_buf)
 fmt.Println("udp2tcp: got from TCP :", tcpReqLen)
 _,err = udp_sock.WriteToUDP(tcp_buf[:tcpReqLen], remoteaddr)

}

// killing sockets when script is killed
 defer tcp_sock.Close()
 defer udp_sock.Close()

}

func tcp2udp(src_conn string,dst_conn string, udp_buf_size int , tcp_buf_size int) {

// start listening on tcp and setup udp conn to the dest

 udp_add,err := net.ResolveUDPAddr("udp",dst_conn)

 tcp_list, err := net.Listen("tcp", src_conn)
 if err != nil { fmt.Println("Error listening:", err.Error()) ; os.Exit(1) }

 tcp_sock, err := tcp_list.Accept()
 if err != nil { fmt.Println("Error accepting: ", err.Error()) ; os.Exit(1) }

 udp_sock, err := net.DialUDP("udp", nil,udp_add)

 udp_buf := make([]byte, udp_buf_size) ; tcp_buf := make([]byte, tcp_buf_size)

for {
// loops over and bridge the data

 // tcp into udp
 tcpReqLen, _ := tcp_sock.Read(tcp_buf)
 fmt.Println("tcp2udp: got from TCP :", tcpReqLen)
 udp_sock.Write([]byte (tcp_buf[:tcpReqLen]))

 // udp into tcp
 udpReqLen, _, _ := udp_sock.ReadFromUDP(udp_buf)
 fmt.Println("tcp2udp: got from UDP :", udpReqLen)
 tcp_sock.Write ([]byte(udp_buf[:udpReqLen]))

}

// killing sockets when script is killed
 defer tcp_sock.Close()
 defer udp_sock.Close()

}

func printUsage() {
     fmt.Printf("\nUsage: gotcpudpbridge TYPE SRC_SOCKET DST_SOCKET UDP_BUF_SIZE TCP_BUF_SIZE\n\nExample: to bridge udp localhost:161 to tcp google.com:80\n gotcpudpbridge udp2tcp 127.0.0.1:161 google.com:80 1024 1024\n\nExample: to bridge tcp localhost:80 to udp google.com:53\n gotcpudpbridge tcp2udp 127.0.0.1:80 google.com:53 1024 1024\n\n")
     os.Exit(1)
}
