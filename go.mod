module fix-redis

go 1.17

require github.com/go-redis/redis v6.15.9+incompatible

require (
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.15.0 // indirect
)

// replace github.com/go-redis/redis => ./../go_oss/redis/
