package main

// Garble - Obfuscate Go builds
// Homepage: https://github.com/burrowers/garble

import (
	"fmt"
	
	"os/exec"
)

func installGarble() {
	// Método 1: Descargar y extraer .tar.gz
	garble_tar_url := "https://github.com/burrowers/garble/archive/refs/tags/v0.13.0.tar.gz"
	garble_cmd_tar := exec.Command("curl", "-L", garble_tar_url, "-o", "package.tar.gz")
	err := garble_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	garble_zip_url := "https://github.com/burrowers/garble/archive/refs/tags/v0.13.0.zip"
	garble_cmd_zip := exec.Command("curl", "-L", garble_zip_url, "-o", "package.zip")
	err = garble_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	garble_bin_url := "https://github.com/burrowers/garble/archive/refs/tags/v0.13.0.bin"
	garble_cmd_bin := exec.Command("curl", "-L", garble_bin_url, "-o", "binary.bin")
	err = garble_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	garble_src_url := "https://github.com/burrowers/garble/archive/refs/tags/v0.13.0.src.tar.gz"
	garble_cmd_src := exec.Command("curl", "-L", garble_src_url, "-o", "source.tar.gz")
	err = garble_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	garble_cmd_direct := exec.Command("./binary")
	err = garble_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: git")
exec.Command("latte", "install", "git")
}
