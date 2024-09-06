package main

// Dutree - Tool to analyze file system usage written in Rust
// Homepage: https://github.com/nachoparker/dutree

import (
	"fmt"
	
	"os/exec"
)

func installDutree() {
	// Método 1: Descargar y extraer .tar.gz
	dutree_tar_url := "https://github.com/nachoparker/dutree/archive/refs/tags/v0.2.18.tar.gz"
	dutree_cmd_tar := exec.Command("curl", "-L", dutree_tar_url, "-o", "package.tar.gz")
	err := dutree_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dutree_zip_url := "https://github.com/nachoparker/dutree/archive/refs/tags/v0.2.18.zip"
	dutree_cmd_zip := exec.Command("curl", "-L", dutree_zip_url, "-o", "package.zip")
	err = dutree_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dutree_bin_url := "https://github.com/nachoparker/dutree/archive/refs/tags/v0.2.18.bin"
	dutree_cmd_bin := exec.Command("curl", "-L", dutree_bin_url, "-o", "binary.bin")
	err = dutree_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dutree_src_url := "https://github.com/nachoparker/dutree/archive/refs/tags/v0.2.18.src.tar.gz"
	dutree_cmd_src := exec.Command("curl", "-L", dutree_src_url, "-o", "source.tar.gz")
	err = dutree_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dutree_cmd_direct := exec.Command("./binary")
	err = dutree_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
