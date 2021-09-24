package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/beevik/ntp"
)

var (
	Version string
)

func getOffset(host string) (offset float64, err error) {
	response, err := ntp.Query(host)
	if err != nil {
		return offset, err
	}
	offset = response.ClockOffset.Seconds() * 1000 * 1000
	return offset, err
}

func putval(collectdIdentifier string, interval int, now time.Time, offset float64) {
	fmt.Printf("PUTVAL \"%s\" interval=%d %d:%f\n", collectdIdentifier, interval, now.Unix(), offset)
}

func main() {
	defaultHostname, _ := os.Hostname()
	if os.Getenv("COLLECTD_HOSTNAME") != "" {
		defaultHostname = os.Getenv("COLLECTD_HOSTNAME")
	}
	defaultInterval := 60
	if os.Getenv("COLLECTD_INTERVAL") != "" {
		defaultInterval, _ = strconv.Atoi(os.Getenv("COLLECTD_INTERVAL"))
	}
	var (
		host             string
		identifier       string
		interval         int
		showVersion      bool
		showVersionShort bool
	)
	flag.StringVar(&host, "host", "169.254.169.123", "destination host.")
	flag.StringVar(&identifier, "identifier", fmt.Sprintf("%s/exec-timesync/gauge-time_offset", defaultHostname), "collectd identifier. first tier is replaced to hostname. respect COLLECTD_HOSTNAME environment variable.")
	flag.IntVar(&interval, "interval", defaultInterval, "interval(sec). respect COLLECTD_INTERVAL environment variable.")
	flag.BoolVar(&showVersion, "version", false, "show version.")
	flag.BoolVar(&showVersionShort, "v", false, "show version.")
	flag.Parse()

	if showVersion || showVersionShort {
		fmt.Printf("Version: %s\n", Version)
		os.Exit(0)
	}

	for {
		now := time.Now()
		offset, err := getOffset(host)
		if err != nil {
			log.Printf("ERROR: %v\n", err)
		}
		putval(identifier, interval, now, offset)

		time.Sleep(time.Duration(interval) * time.Second)
	}

}
