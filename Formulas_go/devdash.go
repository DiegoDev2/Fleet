package main

// Devdash - Highly Configurable Terminal Dashboard for Developers
// Homepage: https://thedevdash.com

import (
	"fmt"
	
	"os/exec"
)

func installDevdash() {
	// Método 1: Descargar y extraer .tar.gz
	devdash_tar_url := "https://github.com/Phantas0s/devdash/archive/refs/tags/v0.5.0.tar.gz"
	devdash_cmd_tar := exec.Command("curl", "-L", devdash_tar_url, "-o", "package.tar.gz")
	err := devdash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	devdash_zip_url := "https://github.com/Phantas0s/devdash/archive/refs/tags/v0.5.0.zip"
	devdash_cmd_zip := exec.Command("curl", "-L", devdash_zip_url, "-o", "package.zip")
	err = devdash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	devdash_bin_url := "https://github.com/Phantas0s/devdash/archive/refs/tags/v0.5.0.bin"
	devdash_cmd_bin := exec.Command("curl", "-L", devdash_bin_url, "-o", "binary.bin")
	err = devdash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	devdash_src_url := "https://github.com/Phantas0s/devdash/archive/refs/tags/v0.5.0.src.tar.gz"
	devdash_cmd_src := exec.Command("curl", "-L", devdash_src_url, "-o", "source.tar.gz")
	err = devdash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	devdash_cmd_direct := exec.Command("./binary")
	err = devdash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
