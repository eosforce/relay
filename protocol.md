# Relay Protocols

this doc define protocols to use relay services.

relay has two type node: relay node and relay api node, 
In test net, relay node watch main chain and do command emit by transfer transaction on main chain,
we use transfer memo to set command params, it is just a test implement. relay api node give some http api to get relay data.


## 1.Relay Command

All Command in relay is a transfer in chain, which is eosforce chain,
we use memo for params.

### 1.1 Map Account

`Map account` can map a account from chain to relay name space,
the account 's owner and active public key is same with the account in chain,
in relay name space, each account has a tag to show which the chain it mapped from.

now map account is implement by a transfer to "relay.a.map", the token transfer is the fee to map.

For account in main chain, use eosio.transfer to tranfer 0.1 EOSC to "relay.a.map" in main chain.

```
cleos --wallet-url http://127.0.0.1:6666 --url http://127.0.0.1:8001 push action eosio transfer '{"from":"eosforce","to":"r.acc.map","quantity":"0.1000 EOS","memo":""}' -p eosforce
```

For account in side chain, use token.transfer to tranfer 0.1 EOS to "relay.a.map" in side chain.

### 1.2 Token In

`Token in` is map token from a its chain to relay name space.

In a chain use transfer to trans token to the account define in token map, this will make relay give the account in relay same token map in relay. 

```
cleos --wallet-url http://127.0.0.1:6666 --url http://127.0.0.1:8001 push action eosio transfer '{"from":"eosforce","to":"r.token.in","quantity":"10000.1000 EOS","memo":""}' -p eosforce
```

### 1.3 Token Out

`Token Out` is trans token from relay to its chain, user use transfer in **its chain** to emit it.

transfer like this:

```
    {
      "account": "eosio",
      "name": "transfer",
      "authorization": [
        {
          "actor": "useraccount",
          "permission": "active"
        }
      ],
      "data": {
        "from": "useraccount",
        "to": "r.token.out",
        "quantity": "0.0100 EOS",
        "memo": "100.0000 EOS"
      }
    }
```

```
cleos --wallet-url http://127.0.0.1:6666 \
      --url http://127.0.0.1:8001 \
      push action eosio transfer \
      '{"from":"eosforce","to":"r.token.out","quantity":"0.1000 EOS","memo":"1000.0000 EOS"}' \
      -p eosforce
```

which means userAccount take out 100.0000 EOS from relay to main chain.

`r.token.out` is fee account in main chain, after test that may be changed.

memo like `main:100.0000 EOS` format is:

```
{chainName}:{Asset}
```


### 1.4 Token Exchange

`Token Exchange` is exchange token to a other token in relay. user use transfer in **its chain** to emit it.

transfer like this:

```
    {
      "account": "eosio",
      "name": "transfer",
      "authorization": [
        {
          "actor": "useraccount",
          "permission": "active"
        }
      ],
      "data": {
        "from": "useraccount",
        "to": "r.exchange",
        "quantity": "0.0300 EOS",
        "memo": "main:100.0000 EOS:side:SYS"
      }
    }
```

this mean useraccount want to buy SYS in side chain by 100.0000 EOS in main chain at relay bancor exchange map.

memo format is:

```
{Sell Asset chainName}:{Sell Asset}:{Buy Asset chainName}:{Buy AssetSymbol}
```

### 1.5 Define Token Map Pair Between Chain And Relay

User can define a token map between chain and relay, it can support users to map this token from chain to relay.

To create a map pair, need a relay account in the token 's chain, this chain's owner pubkey should set to relay,

Command should emit by transfer in main chain.

transfer like this:
 
 ```
     {
       "account": "eosio",
       "name": "transfer",
       "authorization": [
         {
           "actor": "useraccount",
           "permission": "active"
         }
       ],
       "data": {
         "from": "useraccount",
         "to": "r.def.map",
         "quantity": "10.0000 EOS",
         "memo": "side:SYS:eosio.token"
       }
     }
 ```
 
 this mean useraccount create a map by SYS in side chain to relay, `eosio.token` is AssetSymbol Type. 
 memo format is:
 
 ```
 {Asset chainName}:{AssetSymbol}:{AssetSymbol Type}
 ```

### 1.6 Define Token Bancor Exchange Pair In Relay

User can define a token bancor exchange pair in relay, it can support two type token to exchange use bancor protocol in relay.

TODO By FanYang define need params

## 2.Http API

All of http api is to get info from relay.

### 2.1 Account API

#### 2.1.1 Get account info /v1/get/account

Post :

```shell
curl -X POST \
  http://127.0.0.1:8080/v1/get/account \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{"name":"sss","chain":"main"}'
```

Param :

```json
{
  "name": "sss",
  "chain": "main"
}
```

Get :

```json
{
    "name": "sss",
    "chain": "main" // chain is which chain account map from
    // TODO will add some data to there
}
```

#### 2.1.2 Get all accounts with pub key /v1/query/account

Post :

```shell
curl -X POST \
  http://127.0.0.1:8080/v1/query/account \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{"pubkey":"EOSXXXXXXXXXXXXXXXXXXXX"}'
```

Param :

```json
{
  "pubkey": "EOSXXXXXXXXXXXXXXXXXXXX"
}
```

Get :

```json
{
    "account": [
        {
            "chain": "main",
            "name": "eosio.acc1"
        },
        {
            "chain": "main",
            "name": "eosio.acc2"
        },
        {
            "chain": "main",
            "name": "eosio.acc3"
        },
        {
            "chain": "main",
            "name": "eosio.acc4"
        },
        {
            "chain": "side",
            "name": "eosio.acc5"
        }
    ]
}
```

#### 2.1.3 Get account maps info /v1/get/account/maps

No use now.

#### 2.1.4 Get account exchange historys info /v1/get/account/exchanges

get account exchange historys.

Post :

```shell
curl -X POST \
  http://127.0.0.1:8080/v1/get/account/exchanges \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{"name":"sss","chain":"main"}'
```

Param :

```json
{
  "name": "sss",
  "chain": "main"
}
```

Get :

```json
{
    "account": "sss",
    "chain": "main",
    "exchanges": [
        {
            "name": "test",
            "from_account": {
                "chain": "main",
                "name": "sss"
            },
            "to_account": {
                "chain": "main",
                "name": "others"
            },
            "from_token": {
                "Amount": 11111,
                "Precision": 4,
                "Symbol": "EOS",
                "chain": "main"
            },
            "to_token": {
                "Amount": 22222,
                "Precision": 4,
                "Symbol": "SYS",
                "chain": "main"
            }
        },
        {
            "name": "test",
            "from_account": {
                "chain": "main",
                "name": "sss"
            },
            "to_account": {
                "chain": "main",
                "name": "others"
            },
            "from_token": {
                "Amount": 11111,
                "Precision": 4,
                "Symbol": "EOS",
                "chain": "main"
            },
            "to_token": {
                "Amount": 22222,
                "Precision": 4,
                "Symbol": "SYS",
                "chain": "main"
            }
        }
    ]
}
```

#### 2.1.5 Get acoount's all assets  info /v1/get/account/assets

Post :

```shell
curl -X POST \
  http://127.0.0.1:8080/v1/get/account/assets \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{"name":"sss","chain":"main"}'
```

Param :

```json
{
  "name": "sss",
  "chain": "main"
}
```

Get :

```json
{
    "account": "sss",
    "chain": "main",
    "asset": [
        {
            "Amount": 11111,
            "Precision": 4,
            "Symbol": "EOS",
            "chain": "main"
        },
        {
            "Amount": 22222,
            "Precision": 4,
            "Symbol": "EEE",
            "chain": "main"
        },
        {
            "Amount": 33333,
            "Precision": 4,
            "Symbol": "AAA",
            "chain": "main"
        },
        {
            "Amount": 44444,
            "Precision": 4,
            "Symbol": "SYS",
            "chain": "side"
        }
    ]
}
```

### 2.2 Chain State API

#### 2.2.1 Get all chain info /v1/get/chain

Post :

```shell
curl -X POST \
  http://127.0.0.1:8080/v1/get/chain \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json'
```

Get :

```json
{
    "chains": [
        {
            "chain_name": "main",
            "note": "main chain"
        },
        {
            "chain_name": "side",
            "note": "side chain"
        }
    ]
}
```

#### 2.2.2 Get get chain info /v1/get/chain/info

Post :

```shell
curl -X POST \
  http://127.0.0.1:8080/v1/get/chain/info \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -d '{"name":"main"}'
```

Params:

```json
{
  "name" : "main" // name is chain name to get info
}
```

Get :

```json
{
    "account": "main",
    "note": "ddddddddd"
}
```

### 2.3 Relay Token State API

#### 2.3.1 Get all token map info /v1/get/token/maps

Post :

```shell
curl -X POST \
  http://127.0.0.1:8080/v1/get/token/maps
```

Get :

```json
{
    "maps": [
        {
            "Precision": 4,
            "Symbol": "EOS",
            "chain": "main"
        },
        {
            "Precision": 4,
            "Symbol": "SYS",
            "chain": "side"
        }
    ]
}
```

#### 2.3.2 Get all exchanges info /v1/get/token/exchanges

Post :

```shell
curl -X POST \
  http://127.0.0.1:8080/v1/get/token/exchanges
```

Get :

```json
{
    "exchanges": [
        {
            "name": "eos2sys",
            "typ": "bancor",
            "a": {
                "Precision": 4,
                "Symbol": "SYS",
                "chain": "main"
            },
            "b": {
                "Precision": 4,
                "Symbol": "SYS",
                "chain": "side"
            }
        }
    ]
}
```