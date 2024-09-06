package main

// Ko - Build and deploy Go applications on Kubernetes
// Homepage: https://ko.build

import (
	"fmt"
	
	"os/exec"
)

func installKo() {
	// Método 1: Descargar y extraer .tar.gz
	ko_tar_url := "https://github.com/ko-build/ko/archive/refs/tags/v0.16.0.tar.gz"
	ko_cmd_tar := exec.Command("curl", "-L", ko_tar_url, "-o", "package.tar.gz")
	err := ko_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ko_zip_url := "https://github.com/ko-build/ko/archive/refs/tags/v0.16.0.zip"
	ko_cmd_zip := exec.Command("curl", "-L", ko_zip_url, "-o", "package.zip")
	err = ko_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ko_bin_url := "https://github.com/ko-build/ko/archive/refs/tags/v0.16.0.bin"
	ko_cmd_bin := exec.Command("curl", "-L", ko_bin_url, "-o", "binary.bin")
	err = ko_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ko_src_url := "https://github.com/ko-build/ko/archive/refs/tags/v0.16.0.src.tar.gz"
	ko_cmd_src := exec.Command("curl", "-L", ko_src_url, "-o", "source.tar.gz")
	err = ko_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ko_cmd_direct := exec.Command("./binary")
	err = ko_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
