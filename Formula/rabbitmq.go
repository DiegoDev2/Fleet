package main

// Rabbitmq - Messaging and streaming broker
// Homepage: https://www.rabbitmq.com

import (
	"fmt"
	
	"os/exec"
)

func installRabbitmq() {
	// Método 1: Descargar y extraer .tar.gz
	rabbitmq_tar_url := "https://github.com/rabbitmq/rabbitmq-server/releases/download/v3.13.7/rabbitmq-server-generic-unix-3.13.7.tar.xz"
	rabbitmq_cmd_tar := exec.Command("curl", "-L", rabbitmq_tar_url, "-o", "package.tar.gz")
	err := rabbitmq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rabbitmq_zip_url := "https://github.com/rabbitmq/rabbitmq-server/releases/download/v3.13.7/rabbitmq-server-generic-unix-3.13.7.tar.xz"
	rabbitmq_cmd_zip := exec.Command("curl", "-L", rabbitmq_zip_url, "-o", "package.zip")
	err = rabbitmq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rabbitmq_bin_url := "https://github.com/rabbitmq/rabbitmq-server/releases/download/v3.13.7/rabbitmq-server-generic-unix-3.13.7.tar.xz"
	rabbitmq_cmd_bin := exec.Command("curl", "-L", rabbitmq_bin_url, "-o", "binary.bin")
	err = rabbitmq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rabbitmq_src_url := "https://github.com/rabbitmq/rabbitmq-server/releases/download/v3.13.7/rabbitmq-server-generic-unix-3.13.7.tar.xz"
	rabbitmq_cmd_src := exec.Command("curl", "-L", rabbitmq_src_url, "-o", "source.tar.gz")
	err = rabbitmq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rabbitmq_cmd_direct := exec.Command("./binary")
	err = rabbitmq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: erlang")
	exec.Command("latte", "install", "erlang").Run()
}
