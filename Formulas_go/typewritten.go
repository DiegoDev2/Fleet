package main

// Typewritten - Minimal zsh prompt
// Homepage: https://typewritten.dev

import (
	"fmt"
	
	"os/exec"
)

func installTypewritten() {
	// Método 1: Descargar y extraer .tar.gz
	typewritten_tar_url := "https://github.com/reobin/typewritten/archive/refs/tags/v1.5.2.tar.gz"
	typewritten_cmd_tar := exec.Command("curl", "-L", typewritten_tar_url, "-o", "package.tar.gz")
	err := typewritten_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typewritten_zip_url := "https://github.com/reobin/typewritten/archive/refs/tags/v1.5.2.zip"
	typewritten_cmd_zip := exec.Command("curl", "-L", typewritten_zip_url, "-o", "package.zip")
	err = typewritten_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typewritten_bin_url := "https://github.com/reobin/typewritten/archive/refs/tags/v1.5.2.bin"
	typewritten_cmd_bin := exec.Command("curl", "-L", typewritten_bin_url, "-o", "binary.bin")
	err = typewritten_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typewritten_src_url := "https://github.com/reobin/typewritten/archive/refs/tags/v1.5.2.src.tar.gz"
	typewritten_cmd_src := exec.Command("curl", "-L", typewritten_src_url, "-o", "source.tar.gz")
	err = typewritten_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typewritten_cmd_direct := exec.Command("./binary")
	err = typewritten_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zsh")
exec.Command("latte", "install", "zsh")
}
