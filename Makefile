nrf:
	go build -o bin/nrf apps/nrf/cmd/main.go	
clean:
	rm bin/*

.DEFAULT_GOAL := all
all: nrf
