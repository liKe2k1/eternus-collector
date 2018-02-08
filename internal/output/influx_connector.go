package output

import "log"
import "github.com/influxdata/influxdb/client/v2"

func InfluxDb(address, user, pass, database, precision string) (client.Client, client.BatchPoints) {
	if precision == "" {
		precision = "s"
	}

	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:               address,
		Username:           user,
		Password:           pass,
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database,
		Precision: precision,
	})
	if err != nil {
		log.Fatal(err)
	}

	return c, bp
}
