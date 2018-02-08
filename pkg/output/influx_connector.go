package output

import "log"
import "github.com/influxdata/influxdb/client/v2"

func InfluxDb(cfg struct {
	Address     string;
	User        string;
	Pass        string;
	Database    string;
	Sslnoverify bool;
	Precision   string
}) (client.Client, client.BatchPoints) {
	if cfg.Precision == "" {
		cfg.Precision = "s"
	}

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:               cfg.Address,
		Username:           cfg.User,
		Password:           cfg.Pass,
		InsecureSkipVerify: cfg.Sslnoverify,
	})
	if err != nil {
		log.Fatal(err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  cfg.Database,
		Precision: cfg.Precision,
	})

	if err != nil {
		log.Fatal(err)
	}

	return c, bp
}
