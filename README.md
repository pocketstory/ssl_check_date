# ssl_check_date
检查https域名证书的过期时间

---

#### 构建

##### windows下构建

打开cmd命令行

git clone 

cd ssl_check_date

set GOOS=linux

set GOARCH=amd64

go build 

#### 命令使用方法
ssl_check_date -url=https://你的域名
