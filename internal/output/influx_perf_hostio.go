package output

import (
	"log"
	"time"
	"go/types"

	"github.com/like2k1/eternus-collector/internal/types"
)

func InfluxPerfHostIo(hostname string, data []types.PerfHostIO) {

	c, bp := output.InfluxDb("", "", "", "", "s")
	
	for _, elem := range data {

		tags := map[string]string{
			"host": hostname,
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
