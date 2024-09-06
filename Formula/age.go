package main

// Age - Simple, modern, secure file encryption
// Homepage: https://github.com/FiloSottile/age

import (
	"fmt"
	
	"os/exec"
)

func installAge() {
	// Método 1: Descargar y extraer .tar.gz
	age_tar_url := "https://github.com/FiloSottile/age/archive/refs/tags/v1.2.0.tar.gz"
	age_cmd_tar := exec.Command("curl", "-L", age_tar_url, "-o", "package.tar.gz")
	err := age_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	age_zip_url := "https://github.com/FiloSottile/age/archive/refs/tags/v1.2.0.zip"
	age_cmd_zip := exec.Command("curl", "-L", age_zip_url, "-o", "package.zip")
	err = age_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	age_bin_url := "https://github.com/FiloSottile/age/archive/refs/tags/v1.2.0.bin"
	age_cmd_bin := exec.Command("curl", "-L", age_bin_url, "-o", "binary.bin")
	err = age_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	age_src_url := "https://github.com/FiloSottile/age/archive/refs/tags/v1.2.0.src.tar.gz"
	age_cmd_src := exec.Command("curl", "-L", age_src_url, "-o", "source.tar.gz")
	err = age_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	age_cmd_direct := exec.Command("./binary")
	err = age_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
