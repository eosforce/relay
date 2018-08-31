# Relay Cli Tool

a tool to relay for eosforce

## 0. Start

the tool imp by golang, you can install golang by https://golang.org/

build relay-cli tool :

```shell
go get -u -v github.com/eosforce/relay/cmd/relay-cli
cd $GOPATH/src/github.com/eosforce/relay/cmd/relay-cli
go get -u -v ./...
go build
```

then start keosd for tool to send cmd, in default cfg, 
keosd for main eosforce chain is in http://127.0.0.1:6666,
keosd for side eos chain is in http://127.0.0.1:16666.

note the keosd for eos and eosforce has some diff, you should start two process for relay-cli tool.

then import the keys for account.

## 1. Cmds

### 1.1 map account

```shell
./relay-cli account map main eosforce
```

map eosforce in main chain to relay

### 1.2 token in

```shell
./relay-cli token in main eosforce 1000 EOS 
```


### 1.3 token out

```shell
./relay-cli token out main eosforce main eosforce 10000 EOS
```

```shell
./relay-cli token out main eosforce side useraaaaaaaa 10000 SYS
```


### 1.4 exchange test

```shell
./relay-cli exchange test main eosforce 10000 EOS side SYS 
```