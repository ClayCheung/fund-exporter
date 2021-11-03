package fund

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "k8s.io/klog"
)

func ContinueFetchFundsRealTimeData(fundCodes []string, fetchInterval int) {
	tick := time.Tick(time.Duration(fetchInterval) * time.Second)

	for {
		select {
		case <-tick:
			log.Infof("fetch funds %s real-time data per %d seconds", fundCodes, fetchInterval)

			funds := fetchFundsRealTimeData(fundCodes).Data
			for _, fd := range funds {
				labels := make(map[string]string)
				labels[labelCode] = fd.Code
				labels[labelName] = fd.Name
				labels[labelPeriod] = "day"
				fundGrowthRate.With(labels).Set(strToDouble(fd.DayGrowth))
				labels[labelPeriod] = "lastWeek"
				fundGrowthRate.With(labels).Set(strToDouble(fd.LastWeekGrowth))
				labels[labelPeriod] = "lastMonth"
				fundGrowthRate.With(labels).Set(strToDouble(fd.LastMonthGrowth))
				labels[labelPeriod] = "lastThreeMonths"
				fundGrowthRate.With(labels).Set(strToDouble(fd.LastThreeMonthsGrowth))
				labels[labelPeriod] = "lastSixMonths"
				fundGrowthRate.With(labels).Set(strToDouble(fd.LastSixMonthsGrowth))
				labels[labelPeriod] = "lastYear"
				fundGrowthRate.With(labels).Set(strToDouble(fd.LastYearGrowth))

				labels = make(map[string]string)
				labels[labelCode] = fd.Code
				labels[labelName] = fd.Name
				labels[labelNetWorthDate] = fd.NetWorthDate
				fundNetWorth.With(labels).Set(fd.NetWorth)

				labels = make(map[string]string)
				labels[labelCode] = fd.Code
				labels[labelName] = fd.Name
				labels[labelExpectWorthDate] = fd.ExpectWorthDate
				fundExpectWorth.With(labels).Set(fd.ExpectWorth)
			}

		}
	}
}

type FundData struct {
	// 基金代码
	Code string `json:"code"`
	// 基金名称
	Name string `json:"name"`
	// 净值
	NetWorth float64 `json:"netWorth"`
	// 估算净值
	ExpectWorth float64 `json:"expectWorth"`
	// 估算涨幅
	ExpectGrowth string `json:"expectGrowth"`
	// 日涨幅
	DayGrowth string `json:"dayGrowth"`
	// 周涨幅
	LastWeekGrowth string `json:"lastWeekGrowth"`
	// 月涨幅
	LastMonthGrowth string `json:"lastMonthGrowth"`
	// 季度涨幅
	LastThreeMonthsGrowth string `json:"lastThreeMonthsGrowth"`
	// 半年涨幅
	LastSixMonthsGrowth string `json:"lastSixMonthsGrowth"`
	// 年涨幅
	LastYearGrowth string `json:"lastYearGrowth"`
	// 净值日期
	NetWorthDate string `json:"netWorthDate"`
	// 估值日期
	ExpectWorthDate string `json:"expectWorthDate"`
}

type FundDataResp struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []FundData `json:"data"`
	Meta    string     `json:"meta"`
}

func fetchFundsRealTimeData(fundCodes []string) *FundDataResp {
	resp, err := http.Get(fmt.Sprintf("https://api.doctorxiong.club/v1/fund?code=%s", strings.Join(fundCodes, ",")))
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	data := string(bytes)
	fdr := &FundDataResp{}
	if err := json.Unmarshal([]byte(data), fdr); err != nil {
		log.Error(err)
	}
	return fdr
}

func strToDouble(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return f
}
