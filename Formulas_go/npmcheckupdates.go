package main

// NpmCheckUpdates - Find newer versions of dependencies than what your package.json allows
// Homepage: https://github.com/raineorshine/npm-check-updates

import (
	"fmt"
	
	"os/exec"
)

func installNpmCheckUpdates() {
	// Método 1: Descargar y extraer .tar.gz
	npmcheckupdates_tar_url := "https://registry.npmjs.org/npm-check-updates/-/npm-check-updates-17.1.1.tgz"
	npmcheckupdates_cmd_tar := exec.Command("curl", "-L", npmcheckupdates_tar_url, "-o", "package.tar.gz")
	err := npmcheckupdates_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	npmcheckupdates_zip_url := "https://registry.npmjs.org/npm-check-updates/-/npm-check-updates-17.1.1.tgz"
	npmcheckupdates_cmd_zip := exec.Command("curl", "-L", npmcheckupdates_zip_url, "-o", "package.zip")
	err = npmcheckupdates_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	npmcheckupdates_bin_url := "https://registry.npmjs.org/npm-check-updates/-/npm-check-updates-17.1.1.tgz"
	npmcheckupdates_cmd_bin := exec.Command("curl", "-L", npmcheckupdates_bin_url, "-o", "binary.bin")
	err = npmcheckupdates_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	npmcheckupdates_src_url := "https://registry.npmjs.org/npm-check-updates/-/npm-check-updates-17.1.1.tgz"
	npmcheckupdates_cmd_src := exec.Command("curl", "-L", npmcheckupdates_src_url, "-o", "source.tar.gz")
	err = npmcheckupdates_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	npmcheckupdates_cmd_direct := exec.Command("./binary")
	err = npmcheckupdates_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
