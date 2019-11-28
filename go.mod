module github.com/micro-in-cn/config-server

go 1.13

replace github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v0.0.0-20190723190241-65acae22fc9d

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/jinzhu/gorm v1.9.11
	github.com/magiconair/properties v1.8.1
	github.com/micro/go-micro v1.14.0
	github.com/prometheus/common v0.6.0
)
