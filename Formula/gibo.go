package main

// Gibo - Access GitHub's .gitignore boilerplates
// Homepage: https://github.com/simonwhitaker/gibo

import (
	"fmt"
	
	"os/exec"
)

func installGibo() {
	// Método 1: Descargar y extraer .tar.gz
	gibo_tar_url := "https://github.com/simonwhitaker/gibo/archive/refs/tags/v3.0.12.tar.gz"
	gibo_cmd_tar := exec.Command("curl", "-L", gibo_tar_url, "-o", "package.tar.gz")
	err := gibo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gibo_zip_url := "https://github.com/simonwhitaker/gibo/archive/refs/tags/v3.0.12.zip"
	gibo_cmd_zip := exec.Command("curl", "-L", gibo_zip_url, "-o", "package.zip")
	err = gibo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gibo_bin_url := "https://github.com/simonwhitaker/gibo/archive/refs/tags/v3.0.12.bin"
	gibo_cmd_bin := exec.Command("curl", "-L", gibo_bin_url, "-o", "binary.bin")
	err = gibo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gibo_src_url := "https://github.com/simonwhitaker/gibo/archive/refs/tags/v3.0.12.src.tar.gz"
	gibo_cmd_src := exec.Command("curl", "-L", gibo_src_url, "-o", "source.tar.gz")
	err = gibo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gibo_cmd_direct := exec.Command("./binary")
	err = gibo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
