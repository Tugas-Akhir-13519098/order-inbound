# kafka-publisher

How to run:
1. Run `docker-compose up`
2. Create topic by executing `docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic test`
3. Run the service using command `go run main.go`