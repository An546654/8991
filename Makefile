.PHONY: build test clean

build:
	go build -o server server.go
	go build -o client client.go

test: build
	@echo "Starting server in background..."
	./server &
	echo $$! > server.pid
	sleep 1
	@echo "Running client..."
	./client
	RESULT=$$?
	./stop_server.sh
	exit $$RESULT

clean:
	rm -f server client server.pid

