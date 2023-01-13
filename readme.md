A simple framework to implement NRF

From the repo root dir: 

1. type `make` to build

2. run the server with below command:

>> `./bin/nrf --config ../config/nrf.json`

3. Test sending requests to the server with `curl`:

>> register an NF:

>>>> `curl -d '{"Id": "testid100", "NfType": "smf", "Load": 100, "Seen": "xxx", "Info": {"Status": "I am ok"}}' -H "Content-Type: application/json" -X POST http://127.0.0.1:9001/mngr/reg`

>> ping a heartbeat for an NF

>>>> `curl -X POST http://127.0.0.1:9001/mngr/beat/testid100`
