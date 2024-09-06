package main

// Kor - CLI tool to discover unused Kubernetes resources
// Homepage: https://github.com/yonahd/kor

import (
	"fmt"
	
	"os/exec"
)

func installKor() {
	// Método 1: Descargar y extraer .tar.gz
	kor_tar_url := "https://github.com/yonahd/kor/archive/refs/tags/v0.5.5.tar.gz"
	kor_cmd_tar := exec.Command("curl", "-L", kor_tar_url, "-o", "package.tar.gz")
	err := kor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kor_zip_url := "https://github.com/yonahd/kor/archive/refs/tags/v0.5.5.zip"
	kor_cmd_zip := exec.Command("curl", "-L", kor_zip_url, "-o", "package.zip")
	err = kor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kor_bin_url := "https://github.com/yonahd/kor/archive/refs/tags/v0.5.5.bin"
	kor_cmd_bin := exec.Command("curl", "-L", kor_bin_url, "-o", "binary.bin")
	err = kor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kor_src_url := "https://github.com/yonahd/kor/archive/refs/tags/v0.5.5.src.tar.gz"
	kor_cmd_src := exec.Command("curl", "-L", kor_src_url, "-o", "source.tar.gz")
	err = kor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kor_cmd_direct := exec.Command("./binary")
	err = kor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
