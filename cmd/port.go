package cmd

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
)

func ScanPort(ctx context.Context, cmd *cli.Command) error {
	domain := ""
	port := cmd.String("port")

	if cmd.NArg() > 0 {
		domain = cmd.Args().First()
	}

	if port == "all" {
		const maxPort = 9999
		ports := make(chan int, 100)
		results := make(chan int, 100)

		for i := 0; i < 100; i++ {
			go worker(ports, results, domain)
		}

		go func() {
			for i := 1; i <= maxPort; i++ {
				ports <- i
			}
			close(ports)
		}()

		for i := 1; i <= maxPort; i++ {
			p := <-results
			if p != 0 {
				fmt.Println("Port open:", p)
			}
		}

	} else {
		if portIsOpen(domain, port) {
			fmt.Println("Port open:", port)
		} else {
			fmt.Println("Port closed:", port)
		}
	}

	return nil
}

func worker(ports <-chan int, results chan<- int, domain string) {
	for p := range ports {
		if portIsOpen(domain, strconv.Itoa(p)) {
			results <- p
		} else {
			results <- 0
		}
	}
}

func portIsOpen(domain, port string) bool {
	address := net.JoinHostPort(domain, port)
	conn, err := net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		if strings.Contains(err.Error(), "timeout") {
			return false
		}
		return false
	}
	defer conn.Close()
	return true
}
