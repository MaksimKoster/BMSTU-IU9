package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"log"
	"runtime"
	"time"
)

type chs chan string

func main() {
	runtime.GOMAXPROCS(2)
	result := make(chs)
	for i := 0; i < 5; i++ {
		go func(i int) {
			fmt.Println(i)
			result.StartPinging(i, "www.google.com")
		}(i)
	}
	for i := 0; i < 5; i++ {
		select {
		case res := <-result:
			fmt.Println(res)
		case <-time.After(10 * time.Second):
			fmt.Println("Timed out")
			return
		}
	}
}

func (result *chs) StartPinging(i int, addr string) {
	fmt.Println("pinging...")
	pinger, err := ping.NewPinger(addr)
	if err != nil {
		log.Println(err)
	}

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %v%% 				packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	pinger.Count = 1
	pinger.SetPrivileged(true)
	pinger.Run()
	*result <- fmt.Sprint(i, ": ", pinger.Statistics())
}
