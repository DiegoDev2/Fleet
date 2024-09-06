package main

// Descope - Command-line utility for performing common tasks on Descope projects
// Homepage: https://www.descope.com

import (
	"fmt"
	
	"os/exec"
)

func installDescope() {
	// Método 1: Descargar y extraer .tar.gz
	descope_tar_url := "https://github.com/descope/descopecli/archive/refs/tags/v0.8.7.tar.gz"
	descope_cmd_tar := exec.Command("curl", "-L", descope_tar_url, "-o", "package.tar.gz")
	err := descope_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	descope_zip_url := "https://github.com/descope/descopecli/archive/refs/tags/v0.8.7.zip"
	descope_cmd_zip := exec.Command("curl", "-L", descope_zip_url, "-o", "package.zip")
	err = descope_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	descope_bin_url := "https://github.com/descope/descopecli/archive/refs/tags/v0.8.7.bin"
	descope_cmd_bin := exec.Command("curl", "-L", descope_bin_url, "-o", "binary.bin")
	err = descope_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	descope_src_url := "https://github.com/descope/descopecli/archive/refs/tags/v0.8.7.src.tar.gz"
	descope_cmd_src := exec.Command("curl", "-L", descope_src_url, "-o", "source.tar.gz")
	err = descope_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	descope_cmd_direct := exec.Command("./binary")
	err = descope_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
