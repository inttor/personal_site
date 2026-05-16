.PHONY: proto

proto:
	@echo "Generating Go and Python pb files..."
	@mkdir -p server-go/internal/pb
	@mkdir -p worker-python/pb
	protoc -I ./proto \
		--go_out=./server-go/internal/pb --go_opt=paths=source_relative \
		--go-grpc_out=./server-go/internal/pb --go-grpc_opt=paths=source_relative \
		./proto/api.proto
	python -m grpc_tools.protoc -I ./proto --python_out=./worker-python/pb --grpc_python_out=./worker-python/pb ./proto/api.proto
	@echo "Done!"
