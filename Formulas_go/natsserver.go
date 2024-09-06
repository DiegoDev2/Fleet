package main

// NatsServer - Lightweight cloud messaging system
// Homepage: https://nats.io

import (
	"fmt"
	
	"os/exec"
)

func installNatsServer() {
	// Método 1: Descargar y extraer .tar.gz
	natsserver_tar_url := "https://github.com/nats-io/nats-server/archive/refs/tags/v2.10.20.tar.gz"
	natsserver_cmd_tar := exec.Command("curl", "-L", natsserver_tar_url, "-o", "package.tar.gz")
	err := natsserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	natsserver_zip_url := "https://github.com/nats-io/nats-server/archive/refs/tags/v2.10.20.zip"
	natsserver_cmd_zip := exec.Command("curl", "-L", natsserver_zip_url, "-o", "package.zip")
	err = natsserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	natsserver_bin_url := "https://github.com/nats-io/nats-server/archive/refs/tags/v2.10.20.bin"
	natsserver_cmd_bin := exec.Command("curl", "-L", natsserver_bin_url, "-o", "binary.bin")
	err = natsserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	natsserver_src_url := "https://github.com/nats-io/nats-server/archive/refs/tags/v2.10.20.src.tar.gz"
	natsserver_cmd_src := exec.Command("curl", "-L", natsserver_src_url, "-o", "source.tar.gz")
	err = natsserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	natsserver_cmd_direct := exec.Command("./binary")
	err = natsserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
