package main

// Mvnvm - Maven version manager
// Homepage: https://bitbucket.org/mjensen/mvnvm/

import (
	"fmt"
	
	"os/exec"
)

func installMvnvm() {
	// Método 1: Descargar y extraer .tar.gz
	mvnvm_tar_url := "https://bitbucket.org/mjensen/mvnvm/get/mvnvm-1.0.29.tar.gz"
	mvnvm_cmd_tar := exec.Command("curl", "-L", mvnvm_tar_url, "-o", "package.tar.gz")
	err := mvnvm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mvnvm_zip_url := "https://bitbucket.org/mjensen/mvnvm/get/mvnvm-1.0.29.zip"
	mvnvm_cmd_zip := exec.Command("curl", "-L", mvnvm_zip_url, "-o", "package.zip")
	err = mvnvm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mvnvm_bin_url := "https://bitbucket.org/mjensen/mvnvm/get/mvnvm-1.0.29.bin"
	mvnvm_cmd_bin := exec.Command("curl", "-L", mvnvm_bin_url, "-o", "binary.bin")
	err = mvnvm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mvnvm_src_url := "https://bitbucket.org/mjensen/mvnvm/get/mvnvm-1.0.29.src.tar.gz"
	mvnvm_cmd_src := exec.Command("curl", "-L", mvnvm_src_url, "-o", "source.tar.gz")
	err = mvnvm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mvnvm_cmd_direct := exec.Command("./binary")
	err = mvnvm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
