module kk.com/redis

go 1.19

require (
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/glog v1.1.0
	kk.com/time_tool v0.0.0-00010101000000-000000000000
)

require (
	github.com/onsi/ginkgo v1.16.5 // indirect
	github.com/onsi/gomega v1.27.4 // indirect
)

replace kk.com/time_tool => ../time_tool
