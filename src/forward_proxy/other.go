package forward_proxy

// in main.go
// http.Handle("/forward-proxy", &forward_proxy.Pxy{})
// http.ListenAndServe("127.0.0.1:8080", nil)

// Pxy is a empty struct, just for converting to http.Request
// type Pxy struct{}

// ServeHTTP get a local req and transport to local by another
// port, get response then transmit it to old port
// func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
// 	fmt.Printf("received req %s %s %s\n", req.Method, req.Host, req.RemoteAddr)

// 	transport := http.DefaultTransport

// 	outReq := new(http.Request)
// 	*outReq = *req

// 	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
// 		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
// 			clientIP = strings.Join(prior, ",") + "," + clientIP
// 		}
// 		outReq.Header.Set("X-Forward-For", clientIP)
// 	}

// 	res, err := transport.RoundTrip(outReq)
// 	if err != nil {
// 		log.Fatalln(err)
// 		rw.WriteHeader(http.StatusBadGateway)
// 		return
// 	}

// 	for k, v := range res.Header {
// 		for _, v2 := range v {
// 			rw.Header().Add(k, v2)
// 		}
// 	}

// 	rw.WriteHeader(res.StatusCode)
// 	io.Copy(rw, res.Body)
// 	res.Body.Close()
// }
