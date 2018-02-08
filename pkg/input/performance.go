package input

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"

	"github.com/like2k1/eternus-collector/pkg/remote"
	"github.com/like2k1/eternus-collector/pkg/types"
)

type performance struct {
	t *remote.Telnet
	cfg *types.ConfigLayout
}

func NewPerformance(connector *remote.Telnet) *performance {
	p := new(performance)
	p.t = connector
	return p
}

func Performance(cfg types.ConfigLayout) *performance {
	p := new(performance)
	p.t = remote.NewTelnet(cfg.Storage[k].Host, cfg.Storage[k].User, cfg.Storage[k].Pass)
	return p
}

func (p *performance) GetHostIO() []types.PerfHostIO {
	data, err := p.t.Send("show performance -type host-io")
	p.t.CheckErr(err)
	s := p.t.BytesToString(data)
	scanner := bufio.NewScanner(strings.NewReader(s))

	var result = []types.PerfHostIO{}

	for scanner.Scan() {
		var volumeData = regexp.MustCompile(`\s+(?P<id>[0-9]+)\s+(?P<name>[\w\s-]+)\s+(?P<read_iops>[\d]+)\s+(?P<write_iops>[\d]+)\s+(?P<throughput_read>[\d]+)\s+(?P<throughput_write>[\d]+)\s+(?P<response_time_read>[\d]+)\s+(?P<response_time_write>[\d]+)\s+(?P<processing_time_read>[\d]+)\s+(?P<processing_time_write>[\d]+)\s+(?P<cache_hit_rate_read>[\d]+)\s+(?P<cache_hit_rate_write>[\d]+)\s+(?P<prefech>[\d]+)`)
		if volumeData.MatchString(scanner.Text()) {
			res := volumeData.FindStringSubmatch(scanner.Text())
			item := types.PerfHostIO{}
			item.Idx, err = strconv.Atoi(res[1])
			item.Name = strings.TrimSpace(res[2])
			item.IopsRead, err = strconv.Atoi(res[3])
			item.IopsWrite, err = strconv.Atoi(res[4])
			item.ThroughputRead, err = strconv.Atoi(res[5])
			item.ThroughputWrite, err = strconv.Atoi(res[6])
			item.ResponseTimeRead, err = strconv.Atoi(res[7])
			item.ResponseTimeWrite, err = strconv.Atoi(res[8])
			item.ProcessingTimeRead, err = strconv.Atoi(res[9])
			item.ProcessingTimeWrite, err = strconv.Atoi(res[10])
			item.CacheHitRateRead, err = strconv.Atoi(res[11])
			item.CacheHitRateWrite, err = strconv.Atoi(res[12])
			item.Prefech, err = strconv.Atoi(res[12])
			result = append(result, item)
		}
	}
	return result
}

func (p *performance) GetController() []types.PerfController {

	data, err := p.t.Send("show performance -type cm")
	p.t.CheckErr(err)
	s := p.t.BytesToString(data)
	scanner := bufio.NewScanner(strings.NewReader(s))
	var result = []types.PerfController{}

	for scanner.Scan() {
		var data = regexp.MustCompile(`(?P<name>[\w# ]+)\s+(?P<busy_rate>[\d]+)\s+(?P<copy>[\d-]+)`)
		if data.MatchString(scanner.Text()) {
			res := data.FindStringSubmatch(scanner.Text())
			item := types.PerfController{}
			item.Name = strings.TrimSpace(res[1])
			item.BusyRate, err = strconv.Atoi(res[2])
			item.CopyReminderCount, err = strconv.ParseFloat(res[3], 32)

			result = append(result, item)
		}
	}
	return result
}

func (p *performance) GetDisk() []types.PerfDisk {
	data, err := p.t.Send("show performance -type disks")
	p.t.CheckErr(err)
	s := p.t.BytesToString(data)
	scanner := bufio.NewScanner(strings.NewReader(s))
	var result = []types.PerfDisk{}

	i := 0
	for scanner.Scan() {
		var data = regexp.MustCompile(`(?P<name>[\w#-]+)\s+(?P<busy_rate>[\d]+)`)
		if data.MatchString(scanner.Text()) {
			res := data.FindStringSubmatch(scanner.Text())
			item := types.PerfDisk{}
			item.Idx = i
			item.Name = strings.TrimSpace(res[1])
			item.BusyRate, err = strconv.Atoi(res[2])
			result = append(result, item)
			i++
		}
	}
	return result
}
