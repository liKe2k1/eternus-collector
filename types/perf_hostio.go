package types

type PerfHostIO struct {
	Idx                 int
	Name                string
	IopsRead            int
	IopsWrite           int
	ThroughputRead      int
	ThroughputWrite     int
	ResponseTimeRead    int
	ResponseTimeWrite   int
	ProcessingTimeRead  int
	ProcessingTimeWrite int
	CacheHitRateRead    int
	CacheHitRateWrite   int
	Prefech             int
}
