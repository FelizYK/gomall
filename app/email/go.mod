module github.com/FelizYK/gomall/app/email

go 1.23.5

replace github.com/FelizYK/gomall/rpc => ../../rpc

require (
	github.com/FelizYK/gomall/rpc v0.0.0-00010101000000-000000000000
	github.com/kr/pretty v0.3.1
	github.com/nats-io/nats.go v1.39.1
	google.golang.org/protobuf v1.36.5
)

require (
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/nats-io/nkeys v0.4.9 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241202173237-19429a94021a // indirect
	google.golang.org/grpc v1.70.0 // indirect
)
