package main

// Gickup - Backup all your repositories with Ease
// Homepage: https://cooperspencer.github.io/gickup-documentation/

import (
	"fmt"
	
	"os/exec"
)

func installGickup() {
	// Método 1: Descargar y extraer .tar.gz
	gickup_tar_url := "https://github.com/cooperspencer/gickup/archive/refs/tags/v0.10.36.tar.gz"
	gickup_cmd_tar := exec.Command("curl", "-L", gickup_tar_url, "-o", "package.tar.gz")
	err := gickup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gickup_zip_url := "https://github.com/cooperspencer/gickup/archive/refs/tags/v0.10.36.zip"
	gickup_cmd_zip := exec.Command("curl", "-L", gickup_zip_url, "-o", "package.zip")
	err = gickup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gickup_bin_url := "https://github.com/cooperspencer/gickup/archive/refs/tags/v0.10.36.bin"
	gickup_cmd_bin := exec.Command("curl", "-L", gickup_bin_url, "-o", "binary.bin")
	err = gickup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gickup_src_url := "https://github.com/cooperspencer/gickup/archive/refs/tags/v0.10.36.src.tar.gz"
	gickup_cmd_src := exec.Command("curl", "-L", gickup_src_url, "-o", "source.tar.gz")
	err = gickup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gickup_cmd_direct := exec.Command("./binary")
	err = gickup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
