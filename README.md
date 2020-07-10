```
## go服务的模板程序

* server: http server。  
* lib: 该业务用到的公共代码。
* 

## 环境搭建

1. 安装 golang
2. 设定 gopath 

mkdir -p ~/workspace/
mkdir -p ~/workspace/xbj/src

在 ~/.bashrc 设置 GOPATH
    export GOPATH=~/workspace/vendor:~/workspace/xbj
    export PATH="$GOPATH/bin:$PATH"

source ~/.bashrc

3. clone依赖的代码

4. 编写代码
5. 编译运行
cd ~/workspace/xbj/src/go-demo-project/server
make
make start

