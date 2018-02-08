package output

import (
	"log"
	"time"

	"github.com/like2k1/eternus-collector/pkg/types"
	"github.com/influxdata/influxdb/client/v2"
)

func InfluxPerfHostIo(cfg *struct {
	Host  string;
	Label string;
	User  string;
	Pass  string
}, influx struct {
	Address     string;
	User        string;
	Pass        string;
	Database    string;
	Sslnoverify bool;
	Precision   string

}, data []types.PerfHostIO) {

	c, bp := InfluxDb(influx)

	for _, elem := range data {

		tags := map[string]string{
			"host": cfg.Label,
			"name": elem.Name,
		}
		fields := map[string]interface{}{
			"idx":                   elem.Idx,
			"iops_read":             elem.IopsRead,
			"iops_write":            elem.IopsWrite,
			"throughput_read":       elem.ThroughputRead,
			"throughput_write":      elem.ThroughputWrite,
			"response_time_read":    elem.ResponseTimeRead,
			"response_time_write":   elem.ResponseTimeWrite,
			"processing_time_read":  elem.ProcessingTimeRead,
			"processing_time_write": elem.ProcessingTimeWrite,
			"cache_hit_rate_read":   elem.CacheHitRateRead,
			"cache_hit_rate_write":  elem.CacheHitRateWrite,
			"cache_prefetch":        elem.Prefech,
		}

		pt, err := client.NewPoint(
			"eternus-perf-volume-host-io",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			log.Fatal(err)
		}
		bp.AddPoint(pt)
	}

	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}

}
