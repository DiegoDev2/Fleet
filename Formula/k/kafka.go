package main

import (
	"fmt"
	"os/exec"
)

func installKafka() {
	url := "https://downloads.apache.org/kafka/3.3.0/kafka_2.13-3.3.0.tgz"
	fileName := "kafka_2.13-3.3.0.tgz"

	fmt.Println("Downloading Kafka...")
	exec.Command("curl", "-LO", url).Run()

	fmt.Println("Extracting Kafka...")
	exec.Command("tar", "-xzf", fileName).Run()

	fmt.Println("Kafka is ready to use.")
}
