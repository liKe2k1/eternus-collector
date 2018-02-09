package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/gcfg.v1"

	"github.com/like2k1/eternus-collector/pkg/input"
	"github.com/like2k1/eternus-collector/pkg/output"
	"github.com/like2k1/eternus-collector/pkg/types"
	"time"
)

var fConfig = flag.String("config", "/etc/eternus-collector/eternus-collector.conf", "Config file path (Default: /etc/eternus-collector/eternus-collector.conf)")
var fVersion = flag.Bool("version", false, "display the version")
var fDaemonize = flag.Bool("daemon", false, "Override config - Runs in foreground and pull data every X seconds (Default: 60 - See config)")

const usage = `Eternus collector

Usage:
  eternus-collector [commands|flags]

The commands & flags are:

  version             	print the version to stdout
  volume				fetch volume performance data
  controller			fetch controller performance data
  disk					fetch disk performance data

  --config <config>		Config file path (Default: /etc/eternus-collector/eternus-collector.conf)
  --daemon				Override config - Runs in foreground and pull data every X seconds (Default: 60 - See config)

Examples:

  # fetch host io performance data (volumes)
  eternus-collector volumes 

  # fetch controller performance data
  eternus-collector controller

  # fetch disk performance data
  eternus-collector disk
`

func usageExit(rc int) {
	fmt.Println(usage)
	os.Exit(rc)
}

func main() {
	flag.Usage = func() { usageExit(0) }
	flag.Parse()
	args := flag.Args()

	if len(args) > 0 {
		switch args[0] {
		case "version":
			fmt.Printf("eternus-collector %s\n", "1.0")
			return
		}
	}

	var cfg types.ConfigLayout
	err := gcfg.ReadFileInto(&cfg, *fConfig)
	if err != nil {
		log.Fatalf("Failed to parse gcfg data: %s", err)
		os.Exit(1)
	}

	if (*fDaemonize == true) || (cfg.Global.Daemon == true) {
		if cfg.Global.Interval == 0 {
			cfg.Global.Interval = 60
		}
		log.Println("Pulling data every", cfg.Global.Interval, "seconds")
		for {
			log.Println("Processing Volumes")
			for k := range cfg.Storage {
				log.Println("+",k)
				p := input.Performance(cfg.Storage[k])
				output.InfluxPerfHostIo(cfg.Storage[k], cfg.Influx, p.GetHostIO())
			}
			log.Println("Processing Controller")
			for k := range cfg.Storage {
				log.Println("+",k)
				p := input.Performance(cfg.Storage[k])
				output.InfluxPerfController(cfg.Storage[k], cfg.Influx, p.GetController())
			}

			log.Println("Processing Disk")
			for k := range cfg.Storage {
				log.Println("+",k)
				p := input.Performance(cfg.Storage[k])
				output.InfluxPerfDisk(cfg.Storage[k], cfg.Influx, p.GetDisk())
			}
			time.Sleep(time.Second * time.Duration(cfg.Global.Interval))
		}
	}


	if len(args) > 0 {
		switch args[0] {
		case "version":
			fmt.Printf("eternus-collector %s\n", "1.0")
			return

		case "volume":
			for k := range cfg.Storage {
				p := input.Performance(cfg.Storage[k])
				output.InfluxPerfHostIo(cfg.Storage[k], cfg.Influx, p.GetHostIO())
			}
			return
		case "controllers":
			for k := range cfg.Storage {
				p := input.Performance(cfg.Storage[k])
				output.InfluxPerfController(cfg.Storage[k], cfg.Influx, p.GetController())
			}
			return

		case "disks":
			for k := range cfg.Storage {
				p := input.Performance(cfg.Storage[k])
				output.InfluxPerfDisk(cfg.Storage[k], cfg.Influx, p.GetDisk())
			}
			return
		}
	}
}
