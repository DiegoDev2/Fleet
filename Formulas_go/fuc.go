package main

// Fuc - Modern, performance focused unix commands
// Homepage: https://github.com/supercilex/fuc

import (
	"fmt"
	
	"os/exec"
)

func installFuc() {
	// Método 1: Descargar y extraer .tar.gz
	fuc_tar_url := "https://github.com/supercilex/fuc/archive/refs/tags/2.2.0.tar.gz"
	fuc_cmd_tar := exec.Command("curl", "-L", fuc_tar_url, "-o", "package.tar.gz")
	err := fuc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fuc_zip_url := "https://github.com/supercilex/fuc/archive/refs/tags/2.2.0.zip"
	fuc_cmd_zip := exec.Command("curl", "-L", fuc_zip_url, "-o", "package.zip")
	err = fuc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fuc_bin_url := "https://github.com/supercilex/fuc/archive/refs/tags/2.2.0.bin"
	fuc_cmd_bin := exec.Command("curl", "-L", fuc_bin_url, "-o", "binary.bin")
	err = fuc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fuc_src_url := "https://github.com/supercilex/fuc/archive/refs/tags/2.2.0.src.tar.gz"
	fuc_cmd_src := exec.Command("curl", "-L", fuc_src_url, "-o", "source.tar.gz")
	err = fuc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fuc_cmd_direct := exec.Command("./binary")
	err = fuc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
