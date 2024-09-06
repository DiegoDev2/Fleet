package main

// NatsStreamingServer - Lightweight cloud messaging system
// Homepage: https://nats.io

import (
	"fmt"
	
	"os/exec"
)

func installNatsStreamingServer() {
	// Método 1: Descargar y extraer .tar.gz
	natsstreamingserver_tar_url := "https://github.com/nats-io/nats-streaming-server/archive/refs/tags/v0.25.6.tar.gz"
	natsstreamingserver_cmd_tar := exec.Command("curl", "-L", natsstreamingserver_tar_url, "-o", "package.tar.gz")
	err := natsstreamingserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	natsstreamingserver_zip_url := "https://github.com/nats-io/nats-streaming-server/archive/refs/tags/v0.25.6.zip"
	natsstreamingserver_cmd_zip := exec.Command("curl", "-L", natsstreamingserver_zip_url, "-o", "package.zip")
	err = natsstreamingserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	natsstreamingserver_bin_url := "https://github.com/nats-io/nats-streaming-server/archive/refs/tags/v0.25.6.bin"
	natsstreamingserver_cmd_bin := exec.Command("curl", "-L", natsstreamingserver_bin_url, "-o", "binary.bin")
	err = natsstreamingserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	natsstreamingserver_src_url := "https://github.com/nats-io/nats-streaming-server/archive/refs/tags/v0.25.6.src.tar.gz"
	natsstreamingserver_cmd_src := exec.Command("curl", "-L", natsstreamingserver_src_url, "-o", "source.tar.gz")
	err = natsstreamingserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	natsstreamingserver_cmd_direct := exec.Command("./binary")
	err = natsstreamingserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
