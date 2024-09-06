package main

// Uni - Unicode database query tool for the command-line
// Homepage: https://github.com/arp242/uni

import (
	"fmt"
	
	"os/exec"
)

func installUni() {
	// Método 1: Descargar y extraer .tar.gz
	uni_tar_url := "https://github.com/arp242/uni/archive/refs/tags/v2.7.0.tar.gz"
	uni_cmd_tar := exec.Command("curl", "-L", uni_tar_url, "-o", "package.tar.gz")
	err := uni_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uni_zip_url := "https://github.com/arp242/uni/archive/refs/tags/v2.7.0.zip"
	uni_cmd_zip := exec.Command("curl", "-L", uni_zip_url, "-o", "package.zip")
	err = uni_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uni_bin_url := "https://github.com/arp242/uni/archive/refs/tags/v2.7.0.bin"
	uni_cmd_bin := exec.Command("curl", "-L", uni_bin_url, "-o", "binary.bin")
	err = uni_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uni_src_url := "https://github.com/arp242/uni/archive/refs/tags/v2.7.0.src.tar.gz"
	uni_cmd_src := exec.Command("curl", "-L", uni_src_url, "-o", "source.tar.gz")
	err = uni_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uni_cmd_direct := exec.Command("./binary")
	err = uni_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
