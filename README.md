#Fund-exporter

##快速开始

### 容器启动
1. 启动
```shell
 docker run -p 8080:8080 clayz95/fund-exporter:v0.1
```
2. 访问 http://localhost:8080/metrics

### 本地运行

1. 编译 
```shell
go build -o fund-exporter ./main.go 
```
2. 本地运行
```shell
➜  ./fund-exporter --help                                                                                                                                  
Usage of ./fund-exporter:
  -address string
        The address to listen on for HTTP requests. (default ":8080")
  -fetch-interval int
        获取间隔（会定期获取基金实时数据） (default 15)
  -fund-codes string
        基金代码（逗号分隔） (default "000001,400015")
```

3. 访问 http://localhost:8080/metrics

## 预期结果
```shell
# HELP fund_expect_worth Fund expect worth
# TYPE fund_expect_worth gauge
fund_expect_worth{code="000001",expect_worth_date="2021-11-03 15:00:00",name="华夏成长混合"} 1.2274
fund_expect_worth{code="400015",expect_worth_date="2021-11-03 15:00:00",name="东方新能源汽车混合"} 4.5893
# HELP fund_growth Fund growth rate
# TYPE fund_growth gauge
fund_growth{code="000001",name="华夏成长混合",period="day"} -0.57
fund_growth{code="000001",name="华夏成长混合",period="lastMonth"} -4.2
fund_growth{code="000001",name="华夏成长混合",period="lastSixMonths"} -12.13
fund_growth{code="000001",name="华夏成长混合",period="lastThreeMonths"} -6.67
fund_growth{code="000001",name="华夏成长混合",period="lastWeek"} -2.3016
fund_growth{code="000001",name="华夏成长混合",period="lastYear"} -9.01
fund_growth{code="400015",name="东方新能源汽车混合",period="day"} -0.5
fund_growth{code="400015",name="东方新能源汽车混合",period="lastMonth"} 4.32
fund_growth{code="400015",name="东方新能源汽车混合",period="lastSixMonths"} 67.46
fund_growth{code="400015",name="东方新能源汽车混合",period="lastThreeMonths"} 5.84
fund_growth{code="400015",name="东方新能源汽车混合",period="lastWeek"} -2.7999
fund_growth{code="400015",name="东方新能源汽车混合",period="lastYear"} 114.08
# HELP fund_netWorth Fund netWorth
# TYPE fund_netWorth gauge
fund_netWorth{code="000001",name="华夏成长混合",networth_date="2021-11-02"} 1.231
fund_netWorth{code="400015",name="东方新能源汽车混合",networth_date="2021-11-02"} 4.6518
```