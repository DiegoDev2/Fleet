package main

// Scilla - DNS, subdomain, port, directory enumeration tool
// Homepage: https://github.com/edoardottt/scilla

import (
	"fmt"
	
	"os/exec"
)

func installScilla() {
	// Método 1: Descargar y extraer .tar.gz
	scilla_tar_url := "https://github.com/edoardottt/scilla/archive/refs/tags/v1.3.0.tar.gz"
	scilla_cmd_tar := exec.Command("curl", "-L", scilla_tar_url, "-o", "package.tar.gz")
	err := scilla_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scilla_zip_url := "https://github.com/edoardottt/scilla/archive/refs/tags/v1.3.0.zip"
	scilla_cmd_zip := exec.Command("curl", "-L", scilla_zip_url, "-o", "package.zip")
	err = scilla_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scilla_bin_url := "https://github.com/edoardottt/scilla/archive/refs/tags/v1.3.0.bin"
	scilla_cmd_bin := exec.Command("curl", "-L", scilla_bin_url, "-o", "binary.bin")
	err = scilla_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scilla_src_url := "https://github.com/edoardottt/scilla/archive/refs/tags/v1.3.0.src.tar.gz"
	scilla_cmd_src := exec.Command("curl", "-L", scilla_src_url, "-o", "source.tar.gz")
	err = scilla_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scilla_cmd_direct := exec.Command("./binary")
	err = scilla_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
