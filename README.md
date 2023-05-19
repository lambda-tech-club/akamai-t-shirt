# PEACE FOR ALL グラフィックTシャツ（半袖・レギュラーフィット）アカマイ

## Introduction
UNIQLOが発売しているTシャツに印刷されているプログラムを写経・加筆して実行可能にしてみました。
https://www.uniqlo.com/jp/ja/products/E459561-000

### YouTube Shorts

https://youtube.com/shorts/vfnzovi-HG8

### Sample web client

https://codesandbox.io/s/r0hx18

## Getting Started

0. downloading and installing Go: https://go.dev/doc/install
1. clone this repository: `git clone git@github.com:yoidea/akamai-t-shirt.git `
2. change the directory: `cd akamai-t-shirt`
3. start a server: `go run main.go`

## API Usage

### Host Domain
```
http://localhost:3000
```

### Get server status
```http
GET /status
```

#### Description
Get a status that whether the server is processing.

#### Responses

| Value | Description |
| --- | --- |
| ACTIVE | processing |
| INACTIVE | not processing |
| TIMEOUT | could not get status in time |

#### Example

##### Request

```bash
curl localhost:3000/status
```

##### Response

```
INACTIVE
```

### Register target
```http
POST /admin
```

#### Description
Register target properties to the server.

#### Request body parameters

| Name | Type | Required | Description |
| --- | --- | --- | --- |
| target | String | Yes | the name of the target |
| count | Uint32 | Yes | the target of the quantity |

#### Responses

| Value |
| --- |
| Control message issued for Target: `target`, Count: `count` |

#### Example

##### Request

```bash
curl -X POST -d 'target=ちょろすぎて草Tシャツ' -d 'count=200' localhost:3000/admin
```

##### Response

```
Control message issued for Target: ちょろすぎて草Tシャツ, Count: 200
```
