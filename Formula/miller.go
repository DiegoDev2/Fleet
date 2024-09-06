package main

// Miller - Like sed, awk, cut, join & sort for name-indexed data such as CSV
// Homepage: https://github.com/johnkerl/miller

import (
	"fmt"
	
	"os/exec"
)

func installMiller() {
	// Método 1: Descargar y extraer .tar.gz
	miller_tar_url := "https://github.com/johnkerl/miller/archive/refs/tags/v6.12.0.tar.gz"
	miller_cmd_tar := exec.Command("curl", "-L", miller_tar_url, "-o", "package.tar.gz")
	err := miller_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	miller_zip_url := "https://github.com/johnkerl/miller/archive/refs/tags/v6.12.0.zip"
	miller_cmd_zip := exec.Command("curl", "-L", miller_zip_url, "-o", "package.zip")
	err = miller_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	miller_bin_url := "https://github.com/johnkerl/miller/archive/refs/tags/v6.12.0.bin"
	miller_cmd_bin := exec.Command("curl", "-L", miller_bin_url, "-o", "binary.bin")
	err = miller_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	miller_src_url := "https://github.com/johnkerl/miller/archive/refs/tags/v6.12.0.src.tar.gz"
	miller_cmd_src := exec.Command("curl", "-L", miller_src_url, "-o", "source.tar.gz")
	err = miller_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	miller_cmd_direct := exec.Command("./binary")
	err = miller_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
