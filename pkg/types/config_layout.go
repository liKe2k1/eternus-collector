package types

type ConfigLayout struct {

	Global struct {
		Daemon bool
		Interval int
	}

	Storage map[string] *struct {
		Host  string
		Label string
		User  string
		Pass  string
	}
	Influx struct {
		Address     string
		User        string
		Pass        string
		Database    string

		Sslnoverify bool
		Precision 	string
	}
}
