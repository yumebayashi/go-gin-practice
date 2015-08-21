golang-gin-example

goのwebサーバーの実験

```
#go setup
wget https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz
sudo tar zxvf go1.2.1.linux-amd64.tar.gz -C /usr/local

mkdir .go
export GOROOT=/usr/local/go
export GOPATH=~/.go
export PATH=$PATH:$GOROOT/bin

git clone https://github.com/yumebayashi/go-gin-practice.git
cd go-gin-practice
go get github.com/nitrous-io/goop
goop install
export PATH=$PATH:$GOPATH/bin
/srcで
go run *.go
```

