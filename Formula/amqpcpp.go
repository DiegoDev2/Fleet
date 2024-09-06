package main

// AmqpCpp - C++ library for communicating with a RabbitMQ message broker
// Homepage: https://github.com/CopernicaMarketingSoftware/AMQP-CPP

import (
	"fmt"
	
	"os/exec"
)

func installAmqpCpp() {
	// Método 1: Descargar y extraer .tar.gz
	amqpcpp_tar_url := "https://github.com/CopernicaMarketingSoftware/AMQP-CPP/archive/refs/tags/v4.3.26.tar.gz"
	amqpcpp_cmd_tar := exec.Command("curl", "-L", amqpcpp_tar_url, "-o", "package.tar.gz")
	err := amqpcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	amqpcpp_zip_url := "https://github.com/CopernicaMarketingSoftware/AMQP-CPP/archive/refs/tags/v4.3.26.zip"
	amqpcpp_cmd_zip := exec.Command("curl", "-L", amqpcpp_zip_url, "-o", "package.zip")
	err = amqpcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	amqpcpp_bin_url := "https://github.com/CopernicaMarketingSoftware/AMQP-CPP/archive/refs/tags/v4.3.26.bin"
	amqpcpp_cmd_bin := exec.Command("curl", "-L", amqpcpp_bin_url, "-o", "binary.bin")
	err = amqpcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	amqpcpp_src_url := "https://github.com/CopernicaMarketingSoftware/AMQP-CPP/archive/refs/tags/v4.3.26.src.tar.gz"
	amqpcpp_cmd_src := exec.Command("curl", "-L", amqpcpp_src_url, "-o", "source.tar.gz")
	err = amqpcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	amqpcpp_cmd_direct := exec.Command("./binary")
	err = amqpcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
