# Flux
## 查询Influxdb
1.定义数据源，指定查询的bucket
```
from(bucket:"bucket_name") 
```
2.指定时间范围，采用管道符“|>”
```
from(bucket:"bucket_name")
    |> range(start: -1h, stop: -10m)
OR
from(bucket:"bucket_name")
    |> range(start: -1h)
```
3.筛选数据filter()
```
(r) => (r.recordProperty comparisonOperator comparisonExpression)
```