# qpack-parser
QPACK-Parser is a tool that can be used to parse HTTP/3 headers compressed with QPACK.

# How to using it
1. build
```shell
go build main.go
```
2. use
```shell
main.exe -s whatyouwanttocovert
```
for example
```shell
main.exe -s 0000508a089d5c0b8170dc70220fd151886242d276126a0a8fd75f10839bd9ab5f508bed6988b4c7531efdfad867
```
output:
```shell
:authority: 127.0.0.1:6121
:method: GET
:path: /demo/tiles
:scheme: https
accept-encoding: gzip
user-agent: quic-go HTTP/3
```
# How to get hex data?
When using packet capture tools like Wireshark, you can locate the header packet within the HTTP/3 packet, copy the Frame Payload from the HTTP/3 headers, and directly utilize it.
