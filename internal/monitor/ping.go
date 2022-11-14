package monitor

import (
	"context"
	"fmt"
	"github.com/glinton/ping"
)

func PingAdvReqest(host string, timeout int, retries int) (bool, string) {
	res, err := ping.IPv4(context.Background(), host)
	if err != nil {
		panic(err)
		return false, "err"
	} else {
		fmt.Printf("Completed one ping to google.com with %d bytes in %v\n", res.TotalLength, res.RTT)
		return true, string(res.RTT)
	}

}
