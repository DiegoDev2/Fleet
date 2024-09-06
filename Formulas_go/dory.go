package main

// Dory - Development proxy for docker
// Homepage: https://github.com/freedomben/dory

import (
	"fmt"
	
	"os/exec"
)

func installDory() {
	// Método 1: Descargar y extraer .tar.gz
	dory_tar_url := "https://github.com/FreedomBen/dory/archive/refs/tags/v1.2.0.tar.gz"
	dory_cmd_tar := exec.Command("curl", "-L", dory_tar_url, "-o", "package.tar.gz")
	err := dory_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dory_zip_url := "https://github.com/FreedomBen/dory/archive/refs/tags/v1.2.0.zip"
	dory_cmd_zip := exec.Command("curl", "-L", dory_zip_url, "-o", "package.zip")
	err = dory_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dory_bin_url := "https://github.com/FreedomBen/dory/archive/refs/tags/v1.2.0.bin"
	dory_cmd_bin := exec.Command("curl", "-L", dory_bin_url, "-o", "binary.bin")
	err = dory_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dory_src_url := "https://github.com/FreedomBen/dory/archive/refs/tags/v1.2.0.src.tar.gz"
	dory_cmd_src := exec.Command("curl", "-L", dory_src_url, "-o", "source.tar.gz")
	err = dory_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dory_cmd_direct := exec.Command("./binary")
	err = dory_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
