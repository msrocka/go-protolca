@echo off

set olcaproto=..\..\olca-proto\proto

if exist %olcaproto% (
    
    xcopy /y %olcaproto% .
)

protoc --proto_path=. --go_out=../. --go-grpc_out=../. *.proto
