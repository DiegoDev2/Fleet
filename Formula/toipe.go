package main

// Toipe - Yet another typing test, but crab flavoured
// Homepage: https://github.com/Samyak2/toipe

import (
	"fmt"
	
	"os/exec"
)

func installToipe() {
	// Método 1: Descargar y extraer .tar.gz
	toipe_tar_url := "https://github.com/Samyak2/toipe/archive/refs/tags/v0.5.0.tar.gz"
	toipe_cmd_tar := exec.Command("curl", "-L", toipe_tar_url, "-o", "package.tar.gz")
	err := toipe_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	toipe_zip_url := "https://github.com/Samyak2/toipe/archive/refs/tags/v0.5.0.zip"
	toipe_cmd_zip := exec.Command("curl", "-L", toipe_zip_url, "-o", "package.zip")
	err = toipe_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	toipe_bin_url := "https://github.com/Samyak2/toipe/archive/refs/tags/v0.5.0.bin"
	toipe_cmd_bin := exec.Command("curl", "-L", toipe_bin_url, "-o", "binary.bin")
	err = toipe_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	toipe_src_url := "https://github.com/Samyak2/toipe/archive/refs/tags/v0.5.0.src.tar.gz"
	toipe_cmd_src := exec.Command("curl", "-L", toipe_src_url, "-o", "source.tar.gz")
	err = toipe_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	toipe_cmd_direct := exec.Command("./binary")
	err = toipe_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
