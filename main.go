package main

import (
	"flag"
	"fmt"
	"github.com/beevik/ntp"
	"log"
	"os"
	"time"
)

var (
    Version string
)

func getOffset(host string) (offset int64, err error) {
	response, err := ntp.Query(host)
	if err != nil {
		return offset, err
	}
	offset = response.ClockOffset.Milliseconds()
	return offset, err
}

func putval(collectdIdentifier string, now time.Time, offset int64) {
	fmt.Printf(`PUTVAL "%s" %d:%d`, collectdIdentifier, now.Unix(), offset)
}

func main() {
	hostname, _ := os.Hostname()
	var (
		host       string
		identifier string
		interval   int
        showVersion bool
	)
	flag.StringVar(&host, "host", "169.254.169.123", "destination host.")
	flag.StringVar(&identifier, "identifier", fmt.Sprintf("%s/time/offset", hostname), "collectd identifier. first tier is replaced to hostname.")
	flag.IntVar(&interval, "interval", 60, "interval(sec).")
	flag.BoolVar(&showVersion, "version", false, "show version.")
	flag.Parse()

    if showVersion {
        fmt.Printf("Version: %s\n", Version)
        os.Exit(0)
    }

	for {
		now := time.Now()
		offset, err := getOffset(host)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		putval(identifier, now, offset)

		time.Sleep(time.Duration(interval) * time.Second)
	}

}
