protoc -I ./ --go_out=../ --go-grpc_out=../ ./in/* && \
rm -rf ../protos && mv protos ../