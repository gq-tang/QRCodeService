# QRCode Service
 fork from github.com/chuliam/QRCodeService
 
# Pre condition
- vendor go tool
- Go environment
- Redis server

# How to use?
1. go get github.com/gq-tang/QRCodeService
2. modify the conf/app.ini, enter your config info
3. go build
4. run

# About app.ini
[QRCODE]
AppName=QRCode
AppVer=0.1.1.0814
AppPort=8080
AppUri=http://localhost:8080
#0:debug 1:info 2:warn 3:error 4:panic 5:fatal 6:none
AppLogLevel=0
#ByDay:"2006-01-02",ByHour:"2006-01-02-15",ByMonth:"2006-01",
AppLogFileFormat=ByHour
IndexFileUri=./index.html
QRCodeSaveDir=./public/
StaticDir=./static/

[REDIS]
Addr=localhost:6379
Password=
DB=0
PoolSize=5