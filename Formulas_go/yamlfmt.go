package main

// Yamlfmt - Extensible command-line tool to format YAML files
// Homepage: https://github.com/google/yamlfmt

import (
	"fmt"
	
	"os/exec"
)

func installYamlfmt() {
	// Método 1: Descargar y extraer .tar.gz
	yamlfmt_tar_url := "https://github.com/google/yamlfmt/archive/refs/tags/v0.13.0.tar.gz"
	yamlfmt_cmd_tar := exec.Command("curl", "-L", yamlfmt_tar_url, "-o", "package.tar.gz")
	err := yamlfmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yamlfmt_zip_url := "https://github.com/google/yamlfmt/archive/refs/tags/v0.13.0.zip"
	yamlfmt_cmd_zip := exec.Command("curl", "-L", yamlfmt_zip_url, "-o", "package.zip")
	err = yamlfmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yamlfmt_bin_url := "https://github.com/google/yamlfmt/archive/refs/tags/v0.13.0.bin"
	yamlfmt_cmd_bin := exec.Command("curl", "-L", yamlfmt_bin_url, "-o", "binary.bin")
	err = yamlfmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yamlfmt_src_url := "https://github.com/google/yamlfmt/archive/refs/tags/v0.13.0.src.tar.gz"
	yamlfmt_cmd_src := exec.Command("curl", "-L", yamlfmt_src_url, "-o", "source.tar.gz")
	err = yamlfmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yamlfmt_cmd_direct := exec.Command("./binary")
	err = yamlfmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
