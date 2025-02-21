PREFIX=github.com/FelizYK/gomall

.PHONY: frontend-gen
frontend-gen: # make frontend-gen page=common
	cd app && \
	protoc -I ../idl/frontend --go_out=. --go-grpc_out=. ../idl/frontend/${page}.proto

.PHONY: frontend-gen-import
frontend-gen-import: # make frontend-gen-import page=home import=common
	cd app && \
	protoc -I ../idl/frontend \
		--go_out=. --go_opt=M${import}.proto=${PREFIX}/frontend/rpc/${import} \
		--go-grpc_out=. --go-grpc_opt=M${import}.proto=${PREFIX}/frontend/rpc/${import} \
		../idl/frontend/${page}.proto
