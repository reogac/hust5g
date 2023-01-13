a simple framework to implement NRF

type `make` to build

run the server as:

./bin/nrf --config ../config/nrf.json

then send requests to the server with `curl`:

register an NF:

`curl -d '{"Id": "testid100", "NfType": "smf", "Load": 100, "Seen": "xxx", "Info": {"Status": "I am ok"}}' -H "Content-Type: application/json" -X POST http://127.0.0.1:9001/mngr/reg`

ping a heartbeat for a NF

`curl -X POST http://127.0.0.1:9001/mngr/beat/testid100`
