package main

// Clarinet - Command-line tool and runtime for the Clarity smart contract language
// Homepage: https://www.hiro.so/clarinet

import (
	"fmt"
	
	"os/exec"
)

func installClarinet() {
	// Método 1: Descargar y extraer .tar.gz
	clarinet_tar_url := "https://github.com/hirosystems/clarinet/archive/refs/tags/v2.8.0.tar.gz"
	clarinet_cmd_tar := exec.Command("curl", "-L", clarinet_tar_url, "-o", "package.tar.gz")
	err := clarinet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clarinet_zip_url := "https://github.com/hirosystems/clarinet/archive/refs/tags/v2.8.0.zip"
	clarinet_cmd_zip := exec.Command("curl", "-L", clarinet_zip_url, "-o", "package.zip")
	err = clarinet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clarinet_bin_url := "https://github.com/hirosystems/clarinet/archive/refs/tags/v2.8.0.bin"
	clarinet_cmd_bin := exec.Command("curl", "-L", clarinet_bin_url, "-o", "binary.bin")
	err = clarinet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clarinet_src_url := "https://github.com/hirosystems/clarinet/archive/refs/tags/v2.8.0.src.tar.gz"
	clarinet_cmd_src := exec.Command("curl", "-L", clarinet_src_url, "-o", "source.tar.gz")
	err = clarinet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clarinet_cmd_direct := exec.Command("./binary")
	err = clarinet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
