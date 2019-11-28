# warung-pintar-test
Pre Test Warung Pintar


[Prerequisites](#prerequisites) |
[Installation](#installation) |
[Configuration](#environment) |
[License](#license)

## Prerequisites
### Download
- Golang 1.12 - Download and Install [Golang](https://golang.org/dl/)
- Kafka - Download and Install [Kafka](https://kafka.apache.org/quickstart)
- ZooKeeper - Download and Install [Zookeeper](https://zookeeper.apache.org/releases.html/)

### Using Docker
- Docker - Download and Install Docker [Docker](https://docs.docker.com/install/linux/docker-ce/ubuntu/) - Make sure docker running well
- Creating network name it kafka

  ```
    $ docker network create kafka
  ```
- For running kafka we need zookeeper, for running zookeeper using docker do this

  ```
    $ docker run --net=kafka -d --name=zookeeper -e ZOOKEEPER_CLIENT_PORT=2181 confluentinc/cp-zookeeper:4.1.0
  ```
- Next step is Running Kafka, you can use following script

  ```
    docker run --net=kafka -d -p 9092:9092 --name=kafka -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:9092 -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 confluentinc/cp-kafka:4.1.0
  ```
- For connect to kafka we need to setting DNS in server or in our PC. And DNS must be same with KAFKA_ADVERTISED_LISTENERS, because the DNS we will use is "kafka"

  For linux user can edit following file
  
  ```
    $ vim /etc/hosts
  ```
  
  Then add DNS of kafka to that file
  
  ```
    127.0.0.1 kafka
    127.0.0.1 kafka
  ```
- Running zookeeper container

  ```
    $ docker start zookeeper
  ```
- Running kafka container
  ```
    $ docker start kafka
  ```
## Installation
- You can install by clone this repo
  ```
    $ git clone https://github.com/fitraaditama7/warung-pintar-test.git
  ```
- Change directory to directory warung-pintar-test
  ```
    $ cd warung-pintar-test
  ```
- Download dependency using go module
  ```
    $ go mod download
  ```
- Open 3 terminal to running producer, consumer and socket
  - Running producer
    ```
      $ make producer
    ```
  - Running consumer
    ```
      $ make consumer
    ```
  - Running socker
    ```
      $ make socket
    ```
  - Then results like this
    
