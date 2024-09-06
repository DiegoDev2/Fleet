package main

// Mockery - Mock code autogenerator for Golang
// Homepage: https://github.com/vektra/mockery

import (
	"fmt"
	
	"os/exec"
)

func installMockery() {
	// Método 1: Descargar y extraer .tar.gz
	mockery_tar_url := "https://github.com/vektra/mockery/archive/refs/tags/v2.45.0.tar.gz"
	mockery_cmd_tar := exec.Command("curl", "-L", mockery_tar_url, "-o", "package.tar.gz")
	err := mockery_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mockery_zip_url := "https://github.com/vektra/mockery/archive/refs/tags/v2.45.0.zip"
	mockery_cmd_zip := exec.Command("curl", "-L", mockery_zip_url, "-o", "package.zip")
	err = mockery_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mockery_bin_url := "https://github.com/vektra/mockery/archive/refs/tags/v2.45.0.bin"
	mockery_cmd_bin := exec.Command("curl", "-L", mockery_bin_url, "-o", "binary.bin")
	err = mockery_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mockery_src_url := "https://github.com/vektra/mockery/archive/refs/tags/v2.45.0.src.tar.gz"
	mockery_cmd_src := exec.Command("curl", "-L", mockery_src_url, "-o", "source.tar.gz")
	err = mockery_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mockery_cmd_direct := exec.Command("./binary")
	err = mockery_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
