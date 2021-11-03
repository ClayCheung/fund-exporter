package fund

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	labelCode            = "code"
	labelName            = "name"
	labelNetWorthDate    = "networth_date"
	labelExpectWorthDate = "expect_worth_date"
	labelPeriod          = "period"
)

var (
	fundGrowthRate = promauto.NewGaugeVec(prometheus.GaugeOpts{
		//Namespace: "clay",
		Name: "fund_growth",
		Help: "Fund growth rate",
	}, []string{labelCode, labelName, labelPeriod})

	fundNetWorth = promauto.NewGaugeVec(prometheus.GaugeOpts{
		//Namespace: "clay",
		Name: "fund_netWorth",
		Help: "Fund netWorth",
	}, []string{labelCode, labelName, labelNetWorthDate})

	fundExpectWorth = promauto.NewGaugeVec(prometheus.GaugeOpts{
		//Namespace: "clay",
		Name: "fund_expect_worth",
		Help: "Fund expect worth",
	}, []string{labelCode, labelName, labelExpectWorthDate})
)
