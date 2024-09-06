package main

// KtConnect - Toolkit for integrating with kubernetes dev environment more efficiently
// Homepage: https://alibaba.github.io/kt-connect

import (
	"fmt"
	
	"os/exec"
)

func installKtConnect() {
	// Método 1: Descargar y extraer .tar.gz
	ktconnect_tar_url := "https://github.com/alibaba/kt-connect/archive/refs/tags/v0.3.7.tar.gz"
	ktconnect_cmd_tar := exec.Command("curl", "-L", ktconnect_tar_url, "-o", "package.tar.gz")
	err := ktconnect_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ktconnect_zip_url := "https://github.com/alibaba/kt-connect/archive/refs/tags/v0.3.7.zip"
	ktconnect_cmd_zip := exec.Command("curl", "-L", ktconnect_zip_url, "-o", "package.zip")
	err = ktconnect_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ktconnect_bin_url := "https://github.com/alibaba/kt-connect/archive/refs/tags/v0.3.7.bin"
	ktconnect_cmd_bin := exec.Command("curl", "-L", ktconnect_bin_url, "-o", "binary.bin")
	err = ktconnect_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ktconnect_src_url := "https://github.com/alibaba/kt-connect/archive/refs/tags/v0.3.7.src.tar.gz"
	ktconnect_cmd_src := exec.Command("curl", "-L", ktconnect_src_url, "-o", "source.tar.gz")
	err = ktconnect_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ktconnect_cmd_direct := exec.Command("./binary")
	err = ktconnect_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.19")
	exec.Command("latte", "install", "go@1.19").Run()
}
