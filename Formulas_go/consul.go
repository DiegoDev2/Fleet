package main

// Consul - Tool for service discovery, monitoring and configuration
// Homepage: https://www.consul.io

import (
	"fmt"
	
	"os/exec"
)

func installConsul() {
	// Método 1: Descargar y extraer .tar.gz
	consul_tar_url := "https://github.com/hashicorp/consul/archive/refs/tags/v1.16.2.tar.gz"
	consul_cmd_tar := exec.Command("curl", "-L", consul_tar_url, "-o", "package.tar.gz")
	err := consul_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	consul_zip_url := "https://github.com/hashicorp/consul/archive/refs/tags/v1.16.2.zip"
	consul_cmd_zip := exec.Command("curl", "-L", consul_zip_url, "-o", "package.zip")
	err = consul_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	consul_bin_url := "https://github.com/hashicorp/consul/archive/refs/tags/v1.16.2.bin"
	consul_cmd_bin := exec.Command("curl", "-L", consul_bin_url, "-o", "binary.bin")
	err = consul_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	consul_src_url := "https://github.com/hashicorp/consul/archive/refs/tags/v1.16.2.src.tar.gz"
	consul_cmd_src := exec.Command("curl", "-L", consul_src_url, "-o", "source.tar.gz")
	err = consul_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	consul_cmd_direct := exec.Command("./binary")
	err = consul_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
