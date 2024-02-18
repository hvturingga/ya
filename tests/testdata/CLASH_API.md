GET PROXIES
```shell
curl --location 'http://127.0.0.1:9090/proxies'
```

SWITCH
```shell
curl --location --request PUT 'http://127.0.0.1:9090/proxies/:Selector' \
--header 'Content-Type: application/json' \
--data '{
    "name": "name"
}'
```

GET CONFIG
```shell
curl --location 'http://127.0.0.1:9090/configs'
```

EDIT CONFIG
```shell
curl --location --request PATCH 'http://127.0.0.1:9090/configs' \
--header 'Content-Type: application/json' \
--data '{
    "mode": "Global"
}'
```