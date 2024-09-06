package main

// Up - Tool for writing command-line pipes with instant live preview
// Homepage: https://github.com/akavel/up

import (
	"fmt"
	
	"os/exec"
)

func installUp() {
	// Método 1: Descargar y extraer .tar.gz
	up_tar_url := "https://github.com/akavel/up/archive/refs/tags/v0.4.tar.gz"
	up_cmd_tar := exec.Command("curl", "-L", up_tar_url, "-o", "package.tar.gz")
	err := up_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	up_zip_url := "https://github.com/akavel/up/archive/refs/tags/v0.4.zip"
	up_cmd_zip := exec.Command("curl", "-L", up_zip_url, "-o", "package.zip")
	err = up_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	up_bin_url := "https://github.com/akavel/up/archive/refs/tags/v0.4.bin"
	up_cmd_bin := exec.Command("curl", "-L", up_bin_url, "-o", "binary.bin")
	err = up_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	up_src_url := "https://github.com/akavel/up/archive/refs/tags/v0.4.src.tar.gz"
	up_cmd_src := exec.Command("curl", "-L", up_src_url, "-o", "source.tar.gz")
	err = up_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	up_cmd_direct := exec.Command("./binary")
	err = up_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
