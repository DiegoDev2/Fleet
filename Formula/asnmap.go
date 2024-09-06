package main

// Asnmap - Quickly map organization network ranges using ASN information
// Homepage: https://github.com/projectdiscovery/asnmap

import (
	"fmt"
	
	"os/exec"
)

func installAsnmap() {
	// Método 1: Descargar y extraer .tar.gz
	asnmap_tar_url := "https://github.com/projectdiscovery/asnmap/archive/refs/tags/v1.1.1.tar.gz"
	asnmap_cmd_tar := exec.Command("curl", "-L", asnmap_tar_url, "-o", "package.tar.gz")
	err := asnmap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asnmap_zip_url := "https://github.com/projectdiscovery/asnmap/archive/refs/tags/v1.1.1.zip"
	asnmap_cmd_zip := exec.Command("curl", "-L", asnmap_zip_url, "-o", "package.zip")
	err = asnmap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asnmap_bin_url := "https://github.com/projectdiscovery/asnmap/archive/refs/tags/v1.1.1.bin"
	asnmap_cmd_bin := exec.Command("curl", "-L", asnmap_bin_url, "-o", "binary.bin")
	err = asnmap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asnmap_src_url := "https://github.com/projectdiscovery/asnmap/archive/refs/tags/v1.1.1.src.tar.gz"
	asnmap_cmd_src := exec.Command("curl", "-L", asnmap_src_url, "-o", "source.tar.gz")
	err = asnmap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asnmap_cmd_direct := exec.Command("./binary")
	err = asnmap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
