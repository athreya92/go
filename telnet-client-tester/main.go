package main

import (
	"fmt"
	"github.com/athreya92/telnet"
	"log"
	"time"
)

const timeout = 10 * time.Second

func checkErr(err error) {
	if err != nil {
		log.Fatalln("Error:", err)
	}
}

func expect(t *telnet.Conn, d ...string) {
	checkErr(t.SetReadDeadline(time.Now().Add(timeout)))
	checkErr(t.SkipUntil(d...))
}

func sendln(t *telnet.Conn, s string) {
	checkErr(t.SetWriteDeadline(time.Now().Add(timeout)))
	buf := make([]byte, len(s)+1)
	copy(buf, s)
	buf[len(s)] = '\n'
	_, err := t.Write(buf)
	checkErr(err)
}

func main() {
	t, err := telnet.Dial("tcp", "localhost:23")
	checkErr(err)
	t.SetUnixWriteMode(true)
	onelinestatus(t)
	getInfoList(t)
	getMacros(t)
}

func getInfoList(t *telnet.Conn) {
	sendln(t, "getinfoList")
	data, err := t.ReadUntil("End of info list")
	checkErr(err)
	fmt.Println(string(data[:]))
}

func getMacros(t *telnet.Conn) {
	sendln(t, "getmacros")
	data, err := t.ReadUntil("End of macros")
	checkErr(err)
	fmt.Println(string(data[:]))
}

func onelinestatus(t *telnet.Conn) {
	var data []byte
	var err error
	var prefix bool
	sendln(t, "onelinestatus")
	for {
		data, prefix,  err = t.ReadLine()
		if prefix == false {
			break
		}
	}
	checkErr(err)
	fmt.Println(string(data[:]))
}
