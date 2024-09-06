package main

// Microplane - CLI tool to make git changes across many repos
// Homepage: https://github.com/Clever/microplane

import (
	"fmt"
	
	"os/exec"
)

func installMicroplane() {
	// Método 1: Descargar y extraer .tar.gz
	microplane_tar_url := "https://github.com/Clever/microplane/archive/refs/tags/v0.0.34.tar.gz"
	microplane_cmd_tar := exec.Command("curl", "-L", microplane_tar_url, "-o", "package.tar.gz")
	err := microplane_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	microplane_zip_url := "https://github.com/Clever/microplane/archive/refs/tags/v0.0.34.zip"
	microplane_cmd_zip := exec.Command("curl", "-L", microplane_zip_url, "-o", "package.zip")
	err = microplane_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	microplane_bin_url := "https://github.com/Clever/microplane/archive/refs/tags/v0.0.34.bin"
	microplane_cmd_bin := exec.Command("curl", "-L", microplane_bin_url, "-o", "binary.bin")
	err = microplane_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	microplane_src_url := "https://github.com/Clever/microplane/archive/refs/tags/v0.0.34.src.tar.gz"
	microplane_cmd_src := exec.Command("curl", "-L", microplane_src_url, "-o", "source.tar.gz")
	err = microplane_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	microplane_cmd_direct := exec.Command("./binary")
	err = microplane_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
