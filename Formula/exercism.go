package main

// Exercism - Command-line tool to interact with exercism.io
// Homepage: https://exercism.io/cli/

import (
	"fmt"
	
	"os/exec"
)

func installExercism() {
	// Método 1: Descargar y extraer .tar.gz
	exercism_tar_url := "https://github.com/exercism/cli/archive/refs/tags/v3.5.1.tar.gz"
	exercism_cmd_tar := exec.Command("curl", "-L", exercism_tar_url, "-o", "package.tar.gz")
	err := exercism_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exercism_zip_url := "https://github.com/exercism/cli/archive/refs/tags/v3.5.1.zip"
	exercism_cmd_zip := exec.Command("curl", "-L", exercism_zip_url, "-o", "package.zip")
	err = exercism_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exercism_bin_url := "https://github.com/exercism/cli/archive/refs/tags/v3.5.1.bin"
	exercism_cmd_bin := exec.Command("curl", "-L", exercism_bin_url, "-o", "binary.bin")
	err = exercism_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exercism_src_url := "https://github.com/exercism/cli/archive/refs/tags/v3.5.1.src.tar.gz"
	exercism_cmd_src := exec.Command("curl", "-L", exercism_src_url, "-o", "source.tar.gz")
	err = exercism_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exercism_cmd_direct := exec.Command("./binary")
	err = exercism_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
