PREFIX=github.com/FelizYK/gomall

.PHONY: frontend-gen
frontend-gen: # make frontend-gen page=home
	cd app && \
	protoc -I ../idl/frontend --go_out=. ../idl/frontend/${page}.proto

.PHONY: service-gen
service-gen: # make service-gen module=user
	protoc -I idl --go_out=. --go-grpc_out=. idl/${module}.proto
