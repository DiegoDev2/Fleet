package main

// Mdt - Command-line markdown todo list manager
// Homepage: https://github.com/basilioss/mdt

import (
	"fmt"
	
	"os/exec"
)

func installMdt() {
	// Método 1: Descargar y extraer .tar.gz
	mdt_tar_url := "https://github.com/basilioss/mdt/archive/refs/tags/1.4.0.tar.gz"
	mdt_cmd_tar := exec.Command("curl", "-L", mdt_tar_url, "-o", "package.tar.gz")
	err := mdt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdt_zip_url := "https://github.com/basilioss/mdt/archive/refs/tags/1.4.0.zip"
	mdt_cmd_zip := exec.Command("curl", "-L", mdt_zip_url, "-o", "package.zip")
	err = mdt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdt_bin_url := "https://github.com/basilioss/mdt/archive/refs/tags/1.4.0.bin"
	mdt_cmd_bin := exec.Command("curl", "-L", mdt_bin_url, "-o", "binary.bin")
	err = mdt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdt_src_url := "https://github.com/basilioss/mdt/archive/refs/tags/1.4.0.src.tar.gz"
	mdt_cmd_src := exec.Command("curl", "-L", mdt_src_url, "-o", "source.tar.gz")
	err = mdt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdt_cmd_direct := exec.Command("./binary")
	err = mdt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gum")
exec.Command("latte", "install", "gum")
}
