# Relay Cli Tool

a tool to relay for eosforce

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