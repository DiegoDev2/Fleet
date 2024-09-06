package main

// Kafka - Open-source distributed event streaming platform
// Homepage: https://kafka.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installKafka() {
	// Método 1: Descargar y extraer .tar.gz
	kafka_tar_url := "https://www.apache.org/dyn/closer.lua?path=kafka/3.8.0/kafka_2.13-3.8.0.tgz"
	kafka_cmd_tar := exec.Command("curl", "-L", kafka_tar_url, "-o", "package.tar.gz")
	err := kafka_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kafka_zip_url := "https://www.apache.org/dyn/closer.lua?path=kafka/3.8.0/kafka_2.13-3.8.0.tgz"
	kafka_cmd_zip := exec.Command("curl", "-L", kafka_zip_url, "-o", "package.zip")
	err = kafka_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kafka_bin_url := "https://www.apache.org/dyn/closer.lua?path=kafka/3.8.0/kafka_2.13-3.8.0.tgz"
	kafka_cmd_bin := exec.Command("curl", "-L", kafka_bin_url, "-o", "binary.bin")
	err = kafka_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kafka_src_url := "https://www.apache.org/dyn/closer.lua?path=kafka/3.8.0/kafka_2.13-3.8.0.tgz"
	kafka_cmd_src := exec.Command("curl", "-L", kafka_src_url, "-o", "source.tar.gz")
	err = kafka_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kafka_cmd_direct := exec.Command("./binary")
	err = kafka_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: zookeeper")
	exec.Command("latte", "install", "zookeeper").Run()
}
