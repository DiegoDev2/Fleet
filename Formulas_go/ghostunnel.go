package main

// Ghostunnel - Simple SSL/TLS proxy with mutual authentication
// Homepage: https://github.com/ghostunnel/ghostunnel

import (
	"fmt"
	
	"os/exec"
)

func installGhostunnel() {
	// Método 1: Descargar y extraer .tar.gz
	ghostunnel_tar_url := "https://github.com/ghostunnel/ghostunnel/archive/refs/tags/v1.8.1.tar.gz"
	ghostunnel_cmd_tar := exec.Command("curl", "-L", ghostunnel_tar_url, "-o", "package.tar.gz")
	err := ghostunnel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ghostunnel_zip_url := "https://github.com/ghostunnel/ghostunnel/archive/refs/tags/v1.8.1.zip"
	ghostunnel_cmd_zip := exec.Command("curl", "-L", ghostunnel_zip_url, "-o", "package.zip")
	err = ghostunnel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ghostunnel_bin_url := "https://github.com/ghostunnel/ghostunnel/archive/refs/tags/v1.8.1.bin"
	ghostunnel_cmd_bin := exec.Command("curl", "-L", ghostunnel_bin_url, "-o", "binary.bin")
	err = ghostunnel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ghostunnel_src_url := "https://github.com/ghostunnel/ghostunnel/archive/refs/tags/v1.8.1.src.tar.gz"
	ghostunnel_cmd_src := exec.Command("curl", "-L", ghostunnel_src_url, "-o", "source.tar.gz")
	err = ghostunnel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ghostunnel_cmd_direct := exec.Command("./binary")
	err = ghostunnel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
