@echo off

if exist ..\..\olca-modules\olca-proto\src\main\proto (
    xcopy /y ..\..\olca-modules\olca-proto\src\main\proto .
)
protoc --go_out=../. --go-grpc_out=../. *.proto
