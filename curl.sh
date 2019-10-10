curl -X POST \
  http://localhost:1323/data \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Length: 105' \
  -H 'Content-Type: application/json' \
  -H 'Host: localhost:1323' \
  -H 'Postman-Token: 8e34363b-c5ee-4d7c-b75e-9f548e723f17,5c300d9b-0811-4cbb-83bf-c58a3f174f47' \
  -H 'User-Agent: PostmanRuntime/7.17.1' \
  -H 'cache-control: no-cache' \
  -d '{
	"name": "golf",
	"type": "error",
	"error": {
		"code": "MKT-0001",
		"timestamp": 15668011491111
	}
}'


curl -X POST \
  http://localhost:1323/data \
  -H 'Accept: */*' \
  -H 'Accept-Encoding: gzip, deflate' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Length: 76' \
  -H 'Content-Type: application/json' \
  -H 'Host: localhost:1323' \
  -H 'Postman-Token: 4139384b-9934-4a01-8bbf-6a0f8368c1bb,ff486556-d107-4ad8-b7d0-6b244f956871' \
  -H 'User-Agent: PostmanRuntime/7.17.1' \
  -H 'cache-control: no-cache' \
  -d '{
	"name": "golf",
	"type": "campaign_stat",
	"timestamp": 15668011491111
}'