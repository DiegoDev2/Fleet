package main

// Uncover - Tool to discover exposed hosts on the internet using multiple search engines
// Homepage: https://github.com/projectdiscovery/uncover

import (
	"fmt"
	
	"os/exec"
)

func installUncover() {
	// Método 1: Descargar y extraer .tar.gz
	uncover_tar_url := "https://github.com/projectdiscovery/uncover/archive/refs/tags/v1.0.9.tar.gz"
	uncover_cmd_tar := exec.Command("curl", "-L", uncover_tar_url, "-o", "package.tar.gz")
	err := uncover_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uncover_zip_url := "https://github.com/projectdiscovery/uncover/archive/refs/tags/v1.0.9.zip"
	uncover_cmd_zip := exec.Command("curl", "-L", uncover_zip_url, "-o", "package.zip")
	err = uncover_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uncover_bin_url := "https://github.com/projectdiscovery/uncover/archive/refs/tags/v1.0.9.bin"
	uncover_cmd_bin := exec.Command("curl", "-L", uncover_bin_url, "-o", "binary.bin")
	err = uncover_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uncover_src_url := "https://github.com/projectdiscovery/uncover/archive/refs/tags/v1.0.9.src.tar.gz"
	uncover_cmd_src := exec.Command("curl", "-L", uncover_src_url, "-o", "source.tar.gz")
	err = uncover_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uncover_cmd_direct := exec.Command("./binary")
	err = uncover_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
