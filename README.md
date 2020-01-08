# initialise dependencies
```shell script
go get ./
```

# start server
```shell script
go run main.go
```

# start client
```shell script
cd client
go run client.go
```

# how to use protoc
1. install protoc  
    first, go to [protobuf](https://github.com/protocolbuffers/protobuf/releases) release page, download a suitable 
platform, then add the execution file to $PATH

2. run command to generate the template file to current dir
```shell script
cd greeter
# 删除greeter.pd.go, greeter.pb.micro.go
protoc  --micro_out=. --go_out=. ./gretter.proto 
```