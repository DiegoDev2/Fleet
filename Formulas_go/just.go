package main

// Just - Handy way to save and run project-specific commands
// Homepage: https://github.com/casey/just

import (
	"fmt"
	
	"os/exec"
)

func installJust() {
	// Método 1: Descargar y extraer .tar.gz
	just_tar_url := "https://github.com/casey/just/archive/refs/tags/1.35.0.tar.gz"
	just_cmd_tar := exec.Command("curl", "-L", just_tar_url, "-o", "package.tar.gz")
	err := just_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	just_zip_url := "https://github.com/casey/just/archive/refs/tags/1.35.0.zip"
	just_cmd_zip := exec.Command("curl", "-L", just_zip_url, "-o", "package.zip")
	err = just_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	just_bin_url := "https://github.com/casey/just/archive/refs/tags/1.35.0.bin"
	just_cmd_bin := exec.Command("curl", "-L", just_bin_url, "-o", "binary.bin")
	err = just_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	just_src_url := "https://github.com/casey/just/archive/refs/tags/1.35.0.src.tar.gz"
	just_cmd_src := exec.Command("curl", "-L", just_src_url, "-o", "source.tar.gz")
	err = just_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	just_cmd_direct := exec.Command("./binary")
	err = just_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
