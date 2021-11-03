package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/ClayCheung/funds-exporter/fund"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	metricsPath = "/metrics"
)

var (
	addr      = flag.String("address", ":8080", "The address to listen on for HTTP requests.")
	fundCodes = flag.String("fund-codes", "000001,400015", "基金代码（逗号分隔）")
	fetchInterval = flag.Int("fetch-interval", 15, "获取间隔（会定期获取基金实时数据）")
)

func parseCodes(codes string) []string {
	codes = strings.TrimSpace(codes)
	codeList := strings.Split(codes, ",")

	codeMap := make(map[string]struct{})
	for _, code := range codeList {
		code = strings.TrimSpace(code)
		codeMap[code] = struct{}{}
	}

	var ret []string
	for k, _ := range codeMap {
		ret = append(ret, k)
	}
	return ret
}

func main() {
	go fund.ContinueFetchFundsRealTimeData(parseCodes(*fundCodes), *fetchInterval)
	flag.Parse()
	http.Handle(metricsPath, promhttp.Handler())
	log.Fatal(http.ListenAndServe(*addr, nil))
}
