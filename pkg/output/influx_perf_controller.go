package output

import (
	"fmt"
	"log"
	"time"

	"github.com/like2k1/eternus-collector/pkg/types"
	"github.com/influxdata/influxdb/client/v2"
)

func InfluxPerfController(cfg *struct {
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
			"busy_rate":           elem.BusyRate,
			"copy_reminder_count": fmt.Sprintf("%.6f", elem.CopyReminderCount),
		}

		pt, err := client.NewPoint(
			"eternus-perf-controller",
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
