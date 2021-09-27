@echo off

rem this script tries to update the proto-definitions and then generates the
rem Protocol Buffers and gRPC code.

set olcaproto=..\..\olca-proto\proto

if exist %olcaproto% (
    del *.proto
    xcopy /y %olcaproto% .
)

del ..\*.pb.go
protoc --proto_path=. --go_out=../. --go-grpc_out=../. *.proto
