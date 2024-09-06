package main

// Gotags - Tag generator for Go, compatible with ctags
// Homepage: https://github.com/jstemmer/gotags

import (
	"fmt"
	
	"os/exec"
)

func installGotags() {
	// Método 1: Descargar y extraer .tar.gz
	gotags_tar_url := "https://github.com/jstemmer/gotags/archive/refs/tags/v1.4.1.tar.gz"
	gotags_cmd_tar := exec.Command("curl", "-L", gotags_tar_url, "-o", "package.tar.gz")
	err := gotags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gotags_zip_url := "https://github.com/jstemmer/gotags/archive/refs/tags/v1.4.1.zip"
	gotags_cmd_zip := exec.Command("curl", "-L", gotags_zip_url, "-o", "package.zip")
	err = gotags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gotags_bin_url := "https://github.com/jstemmer/gotags/archive/refs/tags/v1.4.1.bin"
	gotags_cmd_bin := exec.Command("curl", "-L", gotags_bin_url, "-o", "binary.bin")
	err = gotags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gotags_src_url := "https://github.com/jstemmer/gotags/archive/refs/tags/v1.4.1.src.tar.gz"
	gotags_cmd_src := exec.Command("curl", "-L", gotags_src_url, "-o", "source.tar.gz")
	err = gotags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gotags_cmd_direct := exec.Command("./binary")
	err = gotags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
