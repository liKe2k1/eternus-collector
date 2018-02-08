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
	cfg *struct { Host string; Label string; User string; Pass string }
}

func Performance(cfg *struct { Host string; Label string; User string; Pass string }) *performance {
	p := new(performance)
	p.cfg = cfg
	return p
}

func (p *performance) GetHostIO() []types.PerfHostIO {

	t := remote.NewTelnet(p.cfg.Host, p.cfg.User, p.cfg.Pass)
	defer t.Close()

	data, err := t.Send("show performance -type host-io")
	t.CheckErr(err)

	s := t.BytesToString(data)

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

	t := remote.NewTelnet(p.cfg.Host, p.cfg.User, p.cfg.Pass)
	defer t.Close()

	data, err := t.Send("show performance -type cm")
	t.CheckErr(err)

	s := t.BytesToString(data)

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

	t := remote.NewTelnet(p.cfg.Host, p.cfg.User, p.cfg.Pass)
	defer t.Close()

	data, err := t.Send("show performance -type disks")
	t.CheckErr(err)

	s := t.BytesToString(data)


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
