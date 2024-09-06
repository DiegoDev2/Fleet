package main

// ServiceWeaver - Programming framework for writing and deploying cloud applications
// Homepage: https://serviceweaver.dev/

import (
	"fmt"
	
	"os/exec"
)

func installServiceWeaver() {
	// Método 1: Descargar y extraer .tar.gz
	serviceweaver_tar_url := "https://github.com/ServiceWeaver/weaver/archive/refs/tags/v0.24.4.tar.gz"
	serviceweaver_cmd_tar := exec.Command("curl", "-L", serviceweaver_tar_url, "-o", "package.tar.gz")
	err := serviceweaver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	serviceweaver_zip_url := "https://github.com/ServiceWeaver/weaver/archive/refs/tags/v0.24.4.zip"
	serviceweaver_cmd_zip := exec.Command("curl", "-L", serviceweaver_zip_url, "-o", "package.zip")
	err = serviceweaver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	serviceweaver_bin_url := "https://github.com/ServiceWeaver/weaver/archive/refs/tags/v0.24.4.bin"
	serviceweaver_cmd_bin := exec.Command("curl", "-L", serviceweaver_bin_url, "-o", "binary.bin")
	err = serviceweaver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	serviceweaver_src_url := "https://github.com/ServiceWeaver/weaver/archive/refs/tags/v0.24.4.src.tar.gz"
	serviceweaver_cmd_src := exec.Command("curl", "-L", serviceweaver_src_url, "-o", "source.tar.gz")
	err = serviceweaver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	serviceweaver_cmd_direct := exec.Command("./binary")
	err = serviceweaver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
