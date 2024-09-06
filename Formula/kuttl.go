package main

// Kuttl - KUbernetes Test TooL
// Homepage: https://kuttl.dev

import (
	"fmt"
	
	"os/exec"
)

func installKuttl() {
	// Método 1: Descargar y extraer .tar.gz
	kuttl_tar_url := "https://github.com/kudobuilder/kuttl/archive/refs/tags/v0.19.0.tar.gz"
	kuttl_cmd_tar := exec.Command("curl", "-L", kuttl_tar_url, "-o", "package.tar.gz")
	err := kuttl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kuttl_zip_url := "https://github.com/kudobuilder/kuttl/archive/refs/tags/v0.19.0.zip"
	kuttl_cmd_zip := exec.Command("curl", "-L", kuttl_zip_url, "-o", "package.zip")
	err = kuttl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kuttl_bin_url := "https://github.com/kudobuilder/kuttl/archive/refs/tags/v0.19.0.bin"
	kuttl_cmd_bin := exec.Command("curl", "-L", kuttl_bin_url, "-o", "binary.bin")
	err = kuttl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kuttl_src_url := "https://github.com/kudobuilder/kuttl/archive/refs/tags/v0.19.0.src.tar.gz"
	kuttl_cmd_src := exec.Command("curl", "-L", kuttl_src_url, "-o", "source.tar.gz")
	err = kuttl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kuttl_cmd_direct := exec.Command("./binary")
	err = kuttl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.22")
	exec.Command("latte", "install", "go@1.22").Run()
	fmt.Println("Instalando dependencia: kubernetes-cli")
	exec.Command("latte", "install", "kubernetes-cli").Run()
}
