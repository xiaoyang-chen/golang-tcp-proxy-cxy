package reverse_proxy

import (
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// in main.go
// fmt.Println("start reverse-proxy")
// reversePxy := reverse_proxy.NewMultiHostsReverseProxy([]*url.URL{
// 	{
// 		Scheme: "http",
// 		Host:   "localhost:9091",
// 	},
// 	{
// 		Scheme: "http",
// 		Host:   "localhost:9092",
// 	},
// })
// log.Fatal(http.ListenAndServe(":9090", reversePxy))

func NewMultiHostsReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
	}
	return &httputil.ReverseProxy{Director: director}
}
