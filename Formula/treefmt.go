package main

// Treefmt - One CLI to format the code tree
// Homepage: https://github.com/numtide/treefmt

import (
	"fmt"
	
	"os/exec"
)

func installTreefmt() {
	// Método 1: Descargar y extraer .tar.gz
	treefmt_tar_url := "https://github.com/numtide/treefmt/archive/refs/tags/v2.0.5.tar.gz"
	treefmt_cmd_tar := exec.Command("curl", "-L", treefmt_tar_url, "-o", "package.tar.gz")
	err := treefmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	treefmt_zip_url := "https://github.com/numtide/treefmt/archive/refs/tags/v2.0.5.zip"
	treefmt_cmd_zip := exec.Command("curl", "-L", treefmt_zip_url, "-o", "package.zip")
	err = treefmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	treefmt_bin_url := "https://github.com/numtide/treefmt/archive/refs/tags/v2.0.5.bin"
	treefmt_cmd_bin := exec.Command("curl", "-L", treefmt_bin_url, "-o", "binary.bin")
	err = treefmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	treefmt_src_url := "https://github.com/numtide/treefmt/archive/refs/tags/v2.0.5.src.tar.gz"
	treefmt_cmd_src := exec.Command("curl", "-L", treefmt_src_url, "-o", "source.tar.gz")
	err = treefmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	treefmt_cmd_direct := exec.Command("./binary")
	err = treefmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
