package main

// Dra - Tool to download release assets from GitHub
// Homepage: https://github.com/devmatteini/dra

import (
	"fmt"
	
	"os/exec"
)

func installDra() {
	// Método 1: Descargar y extraer .tar.gz
	dra_tar_url := "https://github.com/devmatteini/dra/archive/refs/tags/0.6.1.tar.gz"
	dra_cmd_tar := exec.Command("curl", "-L", dra_tar_url, "-o", "package.tar.gz")
	err := dra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dra_zip_url := "https://github.com/devmatteini/dra/archive/refs/tags/0.6.1.zip"
	dra_cmd_zip := exec.Command("curl", "-L", dra_zip_url, "-o", "package.zip")
	err = dra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dra_bin_url := "https://github.com/devmatteini/dra/archive/refs/tags/0.6.1.bin"
	dra_cmd_bin := exec.Command("curl", "-L", dra_bin_url, "-o", "binary.bin")
	err = dra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dra_src_url := "https://github.com/devmatteini/dra/archive/refs/tags/0.6.1.src.tar.gz"
	dra_cmd_src := exec.Command("curl", "-L", dra_src_url, "-o", "source.tar.gz")
	err = dra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dra_cmd_direct := exec.Command("./binary")
	err = dra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
