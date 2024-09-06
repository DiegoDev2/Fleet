package main

// IconNamingUtils - Script to handle icon names in desktop icon themes
// Homepage: https://specifications.freedesktop.org/icon-naming-spec/icon-naming-spec-latest.html

import (
	"fmt"
	
	"os/exec"
)

func installIconNamingUtils() {
	// Método 1: Descargar y extraer .tar.gz
	iconnamingutils_tar_url := "http://tango.freedesktop.org/releases/icon-naming-utils-0.8.90.tar.gz"
	iconnamingutils_cmd_tar := exec.Command("curl", "-L", iconnamingutils_tar_url, "-o", "package.tar.gz")
	err := iconnamingutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iconnamingutils_zip_url := "http://tango.freedesktop.org/releases/icon-naming-utils-0.8.90.zip"
	iconnamingutils_cmd_zip := exec.Command("curl", "-L", iconnamingutils_zip_url, "-o", "package.zip")
	err = iconnamingutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iconnamingutils_bin_url := "http://tango.freedesktop.org/releases/icon-naming-utils-0.8.90.bin"
	iconnamingutils_cmd_bin := exec.Command("curl", "-L", iconnamingutils_bin_url, "-o", "binary.bin")
	err = iconnamingutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iconnamingutils_src_url := "http://tango.freedesktop.org/releases/icon-naming-utils-0.8.90.src.tar.gz"
	iconnamingutils_cmd_src := exec.Command("curl", "-L", iconnamingutils_src_url, "-o", "source.tar.gz")
	err = iconnamingutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iconnamingutils_cmd_direct := exec.Command("./binary")
	err = iconnamingutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: expat")
exec.Command("latte", "install", "expat")
}
