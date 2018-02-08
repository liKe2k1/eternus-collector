package types

type ConfigStorage struct {
	Host  string
	Label string
	User  string
	Pass  string
}

type ConfigLayout struct {
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
