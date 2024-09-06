package main

// ReleaseIt - Generic CLI tool to automate versioning and package publishing related tasks
// Homepage: https://github.com/release-it/release-it

import (
	"fmt"
	
	"os/exec"
)

func installReleaseIt() {
	// Método 1: Descargar y extraer .tar.gz
	releaseit_tar_url := "https://registry.npmjs.org/release-it/-/release-it-17.6.0.tgz"
	releaseit_cmd_tar := exec.Command("curl", "-L", releaseit_tar_url, "-o", "package.tar.gz")
	err := releaseit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	releaseit_zip_url := "https://registry.npmjs.org/release-it/-/release-it-17.6.0.tgz"
	releaseit_cmd_zip := exec.Command("curl", "-L", releaseit_zip_url, "-o", "package.zip")
	err = releaseit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	releaseit_bin_url := "https://registry.npmjs.org/release-it/-/release-it-17.6.0.tgz"
	releaseit_cmd_bin := exec.Command("curl", "-L", releaseit_bin_url, "-o", "binary.bin")
	err = releaseit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	releaseit_src_url := "https://registry.npmjs.org/release-it/-/release-it-17.6.0.tgz"
	releaseit_cmd_src := exec.Command("curl", "-L", releaseit_src_url, "-o", "source.tar.gz")
	err = releaseit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	releaseit_cmd_direct := exec.Command("./binary")
	err = releaseit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
