package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

func Get(url string) {
	client := resty.New()
	client.OnError(func(req *resty.Request, err error) {
		if v, ok := err.(*resty.ResponseError); ok {
			fmt.Println("resty.ResponseError.Response", v.Response)
			fmt.Println("resty.ResponseError.Err", v.Err)
			// v.Response contains the last response from the server
			// v.Err contains the original error
		}
		// Log the error, increment a metric, etc...
	})

	resp, err := client.R().
		EnableTrace().
		Get(url)

	fmt.Println("response info")
	fmt.Println("err:", err)
	fmt.Println("status code", resp.StatusCode())
	fmt.Println("status", resp.Status())
	fmt.Println("proto", resp.Proto())
	fmt.Println("time", resp.Time())
	fmt.Println("received at", resp.ReceivedAt())
	//fmt.Println("body", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())
}

func main() {
	Get("https://yandex.com")
}
