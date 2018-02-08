package output

import "log"
import "github.com/like2k1/eternus-collector/pkg/types"
import "github.com/influxdata/influxdb/client/v2"

func InfluxDb(cfg types.ConfigLayout) (client.Client, client.BatchPoints) {
	if cfg.Influx.Precision == "" {
		cfg.Influx.Precision = "s"
	}

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:               cfg.Influx.Host,
		Username:           cfg.Influx.User,
		Password:           cfg.Influx.Pass,
		InsecureSkipVerify: cfg.Influx.Sslnoverify,
	})
	if err != nil {
		log.Fatal(err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  cfg.Influx.Database,
		Precision: cfg.Influx.Precision,
	})

	if err != nil {
		log.Fatal(err)
	}

	return c, bp
}
