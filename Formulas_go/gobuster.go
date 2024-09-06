package main

// Gobuster - Directory/file & DNS busting tool written in Go
// Homepage: https://github.com/OJ/gobuster

import (
	"fmt"
	
	"os/exec"
)

func installGobuster() {
	// Método 1: Descargar y extraer .tar.gz
	gobuster_tar_url := "https://github.com/OJ/gobuster/archive/refs/tags/v3.6.0.tar.gz"
	gobuster_cmd_tar := exec.Command("curl", "-L", gobuster_tar_url, "-o", "package.tar.gz")
	err := gobuster_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gobuster_zip_url := "https://github.com/OJ/gobuster/archive/refs/tags/v3.6.0.zip"
	gobuster_cmd_zip := exec.Command("curl", "-L", gobuster_zip_url, "-o", "package.zip")
	err = gobuster_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gobuster_bin_url := "https://github.com/OJ/gobuster/archive/refs/tags/v3.6.0.bin"
	gobuster_cmd_bin := exec.Command("curl", "-L", gobuster_bin_url, "-o", "binary.bin")
	err = gobuster_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gobuster_src_url := "https://github.com/OJ/gobuster/archive/refs/tags/v3.6.0.src.tar.gz"
	gobuster_cmd_src := exec.Command("curl", "-L", gobuster_src_url, "-o", "source.tar.gz")
	err = gobuster_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gobuster_cmd_direct := exec.Command("./binary")
	err = gobuster_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
