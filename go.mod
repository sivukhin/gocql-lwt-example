module gocql-lwt

go 1.20

require (
	github.com/gocql/gocql v1.3.1
	github.com/scylladb/gocqlx/v2 v2.8.0
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/klauspost/cpuid/v2 v2.0.3 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.10.0
