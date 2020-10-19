package forward_proxy

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
)

const (
	strHttpPort  = "80"
	strHttpsPort = "443"
	local8081    = "127.0.0.1:8081"
)

// ForwardPxy 是正向代理，通过发送给代理服务器，
// 由代理服务器代为请求，并将响应回传
func ForwardPxy() {
	l, err := net.Listen("tcp", local8081)
	if err != nil {
		log.Panic(err)
	}
	for {
		client, err := l.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handleClientReq(client)
	}
}

func handleClientReq(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	var method, host, addr string
	fmt.Sscanf(string(b[:bytes.IndexByte(b[:], '\n')]), "%s%s", &method, &host)
	hostPortUrl, err := url.Parse(host)
	if err != nil {
		log.Println(err)
		return
	}

	if hostPortUrl.Opaque == strHttpsPort {
		addr = hostPortUrl.Scheme + ":" + strHttpsPort
	} else {
		if strings.Index(hostPortUrl.Host, ":") == -1 {
			addr = hostPortUrl.Host + ":" + strHttpPort
		} else {
			addr = hostPortUrl.Host
		}
	}

	server, err := net.Dial("tcp", addr)
	if err != nil {
		log.Println(err)
		return
	}
	if method == "CONNECT" {
		fmt.Fprint(client, "HTTP/1.1 200 Connection established\n")
	} else {
		server.Write(b[:n])
	}

	// 为什么做两次copy，第一次需要另起Goroutine处理
	//go io.Copy(server, client)
	io.Copy(client, server)
}
