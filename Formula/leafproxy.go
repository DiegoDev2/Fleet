package main

// LeafProxy - Lightweight and fast proxy utility
// Homepage: https://github.com/eycorsican/leaf

import (
	"fmt"
	
	"os/exec"
)

func installLeafProxy() {
	// Método 1: Descargar y extraer .tar.gz
	leafproxy_tar_url := "https://github.com/eycorsican/leaf/archive/refs/tags/v0.11.0.tar.gz"
	leafproxy_cmd_tar := exec.Command("curl", "-L", leafproxy_tar_url, "-o", "package.tar.gz")
	err := leafproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leafproxy_zip_url := "https://github.com/eycorsican/leaf/archive/refs/tags/v0.11.0.zip"
	leafproxy_cmd_zip := exec.Command("curl", "-L", leafproxy_zip_url, "-o", "package.zip")
	err = leafproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leafproxy_bin_url := "https://github.com/eycorsican/leaf/archive/refs/tags/v0.11.0.bin"
	leafproxy_cmd_bin := exec.Command("curl", "-L", leafproxy_bin_url, "-o", "binary.bin")
	err = leafproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leafproxy_src_url := "https://github.com/eycorsican/leaf/archive/refs/tags/v0.11.0.src.tar.gz"
	leafproxy_cmd_src := exec.Command("curl", "-L", leafproxy_src_url, "-o", "source.tar.gz")
	err = leafproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leafproxy_cmd_direct := exec.Command("./binary")
	err = leafproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
