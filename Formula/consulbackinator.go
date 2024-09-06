package main

// ConsulBackinator - Consul backup and restoration application
// Homepage: https://github.com/myENA/consul-backinator

import (
	"fmt"
	
	"os/exec"
)

func installConsulBackinator() {
	// Método 1: Descargar y extraer .tar.gz
	consulbackinator_tar_url := "https://github.com/myENA/consul-backinator/archive/refs/tags/v1.6.6.tar.gz"
	consulbackinator_cmd_tar := exec.Command("curl", "-L", consulbackinator_tar_url, "-o", "package.tar.gz")
	err := consulbackinator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	consulbackinator_zip_url := "https://github.com/myENA/consul-backinator/archive/refs/tags/v1.6.6.zip"
	consulbackinator_cmd_zip := exec.Command("curl", "-L", consulbackinator_zip_url, "-o", "package.zip")
	err = consulbackinator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	consulbackinator_bin_url := "https://github.com/myENA/consul-backinator/archive/refs/tags/v1.6.6.bin"
	consulbackinator_cmd_bin := exec.Command("curl", "-L", consulbackinator_bin_url, "-o", "binary.bin")
	err = consulbackinator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	consulbackinator_src_url := "https://github.com/myENA/consul-backinator/archive/refs/tags/v1.6.6.src.tar.gz"
	consulbackinator_cmd_src := exec.Command("curl", "-L", consulbackinator_src_url, "-o", "source.tar.gz")
	err = consulbackinator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	consulbackinator_cmd_direct := exec.Command("./binary")
	err = consulbackinator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
