package main

import (
	"fmt"
	"sync"
	"github.com/miekg/dns"
	"os"
	"bufio"

)

func worker(dnsServers chan string, wg *sync.WaitGroup){
	for nsServer:= range dnsServers {
		var msg dns.Msg
		fqdn := dns.Fqdn("as.com")
		msg.SetQuestion(fqdn,dns.TypeA)
		_,err := dns.Exchange(&msg,nsServer+":53")
		if err != nil {
		}else {
			fmt.Println(nsServer)
		}
		wg.Done()
	}
}
//

func main() {
	dnsServers := make(chan string, 1000)
	var wg sync.WaitGroup

	sc := bufio.NewScanner(os.Stdin)
	for i := 0; i < cap(dnsServers); i++ {
		go worker(dnsServers,&wg)
	}
	for sc.Scan() {
		wg.Add(1)
		dnsServers <- sc.Text()
	}
	wg.Wait()
	close(dnsServers)
}
