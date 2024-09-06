package main

// Karn - Manage multiple Git identities
// Homepage: https://github.com/prydonius/karn

import (
	"fmt"
	
	"os/exec"
)

func installKarn() {
	// Método 1: Descargar y extraer .tar.gz
	karn_tar_url := "https://github.com/prydonius/karn/archive/refs/tags/v0.1.0.tar.gz"
	karn_cmd_tar := exec.Command("curl", "-L", karn_tar_url, "-o", "package.tar.gz")
	err := karn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	karn_zip_url := "https://github.com/prydonius/karn/archive/refs/tags/v0.1.0.zip"
	karn_cmd_zip := exec.Command("curl", "-L", karn_zip_url, "-o", "package.zip")
	err = karn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	karn_bin_url := "https://github.com/prydonius/karn/archive/refs/tags/v0.1.0.bin"
	karn_cmd_bin := exec.Command("curl", "-L", karn_bin_url, "-o", "binary.bin")
	err = karn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	karn_src_url := "https://github.com/prydonius/karn/archive/refs/tags/v0.1.0.src.tar.gz"
	karn_cmd_src := exec.Command("curl", "-L", karn_src_url, "-o", "source.tar.gz")
	err = karn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	karn_cmd_direct := exec.Command("./binary")
	err = karn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
