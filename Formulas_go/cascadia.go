package main

// Cascadia - Go cascadia package command-line CSS selector
// Homepage: https://github.com/suntong/cascadia

import (
	"fmt"
	
	"os/exec"
)

func installCascadia() {
	// Método 1: Descargar y extraer .tar.gz
	cascadia_tar_url := "https://github.com/suntong/cascadia/archive/refs/tags/v1.3.0.tar.gz"
	cascadia_cmd_tar := exec.Command("curl", "-L", cascadia_tar_url, "-o", "package.tar.gz")
	err := cascadia_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cascadia_zip_url := "https://github.com/suntong/cascadia/archive/refs/tags/v1.3.0.zip"
	cascadia_cmd_zip := exec.Command("curl", "-L", cascadia_zip_url, "-o", "package.zip")
	err = cascadia_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cascadia_bin_url := "https://github.com/suntong/cascadia/archive/refs/tags/v1.3.0.bin"
	cascadia_cmd_bin := exec.Command("curl", "-L", cascadia_bin_url, "-o", "binary.bin")
	err = cascadia_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cascadia_src_url := "https://github.com/suntong/cascadia/archive/refs/tags/v1.3.0.src.tar.gz"
	cascadia_cmd_src := exec.Command("curl", "-L", cascadia_src_url, "-o", "source.tar.gz")
	err = cascadia_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cascadia_cmd_direct := exec.Command("./binary")
	err = cascadia_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
