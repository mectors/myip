package main

import (
	"fmt"
	"os/exec"
	"os"
	"log"
	"net"
	"time"
)

var startTime time.Time

func uptime() time.Duration {
	return time.Since(startTime)
}

func init() {
	startTime = time.Now()
}

func Scroll(text string) {
	fmt.Println("received:" + text)
	go func() {
		snap := os.Getenv("SNAP")
		cmd := exec.Command(snap+"/dcled", "-m", text)
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
  }()
}

func main() {
	var msg string
	for (uptime() < (time.Second * 120)) {
		  msg = ""
			fmt.Printf("uptime %s\n", uptime())
			ifaces, _ := net.Interfaces()
			// handle err
			for _, i := range ifaces {
			    addrs, _ := i.Addrs()
			    // handle err
			    for _, addr := range addrs {
			        var ip net.IP

							switch v := addr.(type) {
			        case *net.IPNet:
			                ip = v.IP
											if(!ip.IsLoopback() && ip.To4() != nil) {
												msg = msg + "  " + i.Name +" - "+ip.String()
											}
			        }
			    }
			}
      Scroll(msg)
			time.Sleep(time.Second * 20)
		}
}
