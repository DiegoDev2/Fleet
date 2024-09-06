package main

// Mapcidr - Subnet/CIDR operation utility
// Homepage: https://projectdiscovery.io

import (
	"fmt"
	
	"os/exec"
)

func installMapcidr() {
	// Método 1: Descargar y extraer .tar.gz
	mapcidr_tar_url := "https://github.com/projectdiscovery/mapcidr/archive/refs/tags/v1.1.34.tar.gz"
	mapcidr_cmd_tar := exec.Command("curl", "-L", mapcidr_tar_url, "-o", "package.tar.gz")
	err := mapcidr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mapcidr_zip_url := "https://github.com/projectdiscovery/mapcidr/archive/refs/tags/v1.1.34.zip"
	mapcidr_cmd_zip := exec.Command("curl", "-L", mapcidr_zip_url, "-o", "package.zip")
	err = mapcidr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mapcidr_bin_url := "https://github.com/projectdiscovery/mapcidr/archive/refs/tags/v1.1.34.bin"
	mapcidr_cmd_bin := exec.Command("curl", "-L", mapcidr_bin_url, "-o", "binary.bin")
	err = mapcidr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mapcidr_src_url := "https://github.com/projectdiscovery/mapcidr/archive/refs/tags/v1.1.34.src.tar.gz"
	mapcidr_cmd_src := exec.Command("curl", "-L", mapcidr_src_url, "-o", "source.tar.gz")
	err = mapcidr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mapcidr_cmd_direct := exec.Command("./binary")
	err = mapcidr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
