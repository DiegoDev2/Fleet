package main

// Dashing - Generate Dash documentation from HTML files
// Homepage: https://github.com/technosophos/dashing

import (
	"fmt"
	
	"os/exec"
)

func installDashing() {
	// Método 1: Descargar y extraer .tar.gz
	dashing_tar_url := "https://github.com/technosophos/dashing/archive/refs/tags/0.4.0.tar.gz"
	dashing_cmd_tar := exec.Command("curl", "-L", dashing_tar_url, "-o", "package.tar.gz")
	err := dashing_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dashing_zip_url := "https://github.com/technosophos/dashing/archive/refs/tags/0.4.0.zip"
	dashing_cmd_zip := exec.Command("curl", "-L", dashing_zip_url, "-o", "package.zip")
	err = dashing_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dashing_bin_url := "https://github.com/technosophos/dashing/archive/refs/tags/0.4.0.bin"
	dashing_cmd_bin := exec.Command("curl", "-L", dashing_bin_url, "-o", "binary.bin")
	err = dashing_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dashing_src_url := "https://github.com/technosophos/dashing/archive/refs/tags/0.4.0.src.tar.gz"
	dashing_cmd_src := exec.Command("curl", "-L", dashing_src_url, "-o", "source.tar.gz")
	err = dashing_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dashing_cmd_direct := exec.Command("./binary")
	err = dashing_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
