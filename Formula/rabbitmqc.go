package main

// RabbitmqC - C AMQP client library for RabbitMQ
// Homepage: https://github.com/alanxz/rabbitmq-c

import (
	"fmt"
	
	"os/exec"
)

func installRabbitmqC() {
	// Método 1: Descargar y extraer .tar.gz
	rabbitmqc_tar_url := "https://github.com/alanxz/rabbitmq-c/archive/refs/tags/v0.14.0.tar.gz"
	rabbitmqc_cmd_tar := exec.Command("curl", "-L", rabbitmqc_tar_url, "-o", "package.tar.gz")
	err := rabbitmqc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rabbitmqc_zip_url := "https://github.com/alanxz/rabbitmq-c/archive/refs/tags/v0.14.0.zip"
	rabbitmqc_cmd_zip := exec.Command("curl", "-L", rabbitmqc_zip_url, "-o", "package.zip")
	err = rabbitmqc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rabbitmqc_bin_url := "https://github.com/alanxz/rabbitmq-c/archive/refs/tags/v0.14.0.bin"
	rabbitmqc_cmd_bin := exec.Command("curl", "-L", rabbitmqc_bin_url, "-o", "binary.bin")
	err = rabbitmqc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rabbitmqc_src_url := "https://github.com/alanxz/rabbitmq-c/archive/refs/tags/v0.14.0.src.tar.gz"
	rabbitmqc_cmd_src := exec.Command("curl", "-L", rabbitmqc_src_url, "-o", "source.tar.gz")
	err = rabbitmqc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rabbitmqc_cmd_direct := exec.Command("./binary")
	err = rabbitmqc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xmlto")
	exec.Command("latte", "install", "xmlto").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: popt")
	exec.Command("latte", "install", "popt").Run()
}
