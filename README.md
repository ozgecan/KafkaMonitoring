
This project aims is use Kafka by Go Programing Language. You can build your Kafka docker with the below commands If you have a docker engine and docker-compose on your computer. 

Open KafkaMonitoring/docker folder.

> sudo docker-compose up -d

You need to create "foo" topic on your kafka If your kafka container up on your computer.

You can check docker container with; 

> sudo docker ps

You can create topic with;

> sudo docker exec broker kafka-topics --bootstrap-server broker:9092 --create --topic foo

Now, You can run your main.go;

> go run main.go

