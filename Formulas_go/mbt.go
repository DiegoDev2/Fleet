package main

// Mbt - Multi-Target Application (MTA) build tool for Cloud Applications
// Homepage: https://sap.github.io/cloud-mta-build-tool

import (
	"fmt"
	
	"os/exec"
)

func installMbt() {
	// Método 1: Descargar y extraer .tar.gz
	mbt_tar_url := "https://github.com/SAP/cloud-mta-build-tool/archive/refs/tags/v1.2.31.tar.gz"
	mbt_cmd_tar := exec.Command("curl", "-L", mbt_tar_url, "-o", "package.tar.gz")
	err := mbt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mbt_zip_url := "https://github.com/SAP/cloud-mta-build-tool/archive/refs/tags/v1.2.31.zip"
	mbt_cmd_zip := exec.Command("curl", "-L", mbt_zip_url, "-o", "package.zip")
	err = mbt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mbt_bin_url := "https://github.com/SAP/cloud-mta-build-tool/archive/refs/tags/v1.2.31.bin"
	mbt_cmd_bin := exec.Command("curl", "-L", mbt_bin_url, "-o", "binary.bin")
	err = mbt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mbt_src_url := "https://github.com/SAP/cloud-mta-build-tool/archive/refs/tags/v1.2.31.src.tar.gz"
	mbt_cmd_src := exec.Command("curl", "-L", mbt_src_url, "-o", "source.tar.gz")
	err = mbt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mbt_cmd_direct := exec.Command("./binary")
	err = mbt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
