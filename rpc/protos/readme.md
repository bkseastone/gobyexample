安装 protoc-gen-go go get -u github.com/golang/protobuf/protoc-gen-go
安装 protoc  git clone https://github.com/Microsoft/vcpkg.git
    cd vcpkg 
    .\bootstrap-vcpkg.bat
    vcpkg install protobuf protobuf:x64-windows
    或者直接github下载二进制版
编译pb protoc --go_out=plugins=grpc:. *.proto