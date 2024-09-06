package main

// Pulsarctl - CLI for Apache Pulsar written in Go
// Homepage: https://streamnative.io/

import (
	"fmt"
	
	"os/exec"
)

func installPulsarctl() {
	// Método 1: Descargar y extraer .tar.gz
	pulsarctl_tar_url := "https://github.com/streamnative/pulsarctl/archive/refs/tags/v3.3.1.4.tar.gz"
	pulsarctl_cmd_tar := exec.Command("curl", "-L", pulsarctl_tar_url, "-o", "package.tar.gz")
	err := pulsarctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pulsarctl_zip_url := "https://github.com/streamnative/pulsarctl/archive/refs/tags/v3.3.1.4.zip"
	pulsarctl_cmd_zip := exec.Command("curl", "-L", pulsarctl_zip_url, "-o", "package.zip")
	err = pulsarctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pulsarctl_bin_url := "https://github.com/streamnative/pulsarctl/archive/refs/tags/v3.3.1.4.bin"
	pulsarctl_cmd_bin := exec.Command("curl", "-L", pulsarctl_bin_url, "-o", "binary.bin")
	err = pulsarctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pulsarctl_src_url := "https://github.com/streamnative/pulsarctl/archive/refs/tags/v3.3.1.4.src.tar.gz"
	pulsarctl_cmd_src := exec.Command("curl", "-L", pulsarctl_src_url, "-o", "source.tar.gz")
	err = pulsarctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pulsarctl_cmd_direct := exec.Command("./binary")
	err = pulsarctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
