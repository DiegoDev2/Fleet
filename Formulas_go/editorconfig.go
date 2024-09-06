package main

// Editorconfig - Maintain consistent coding style between multiple editors
// Homepage: https://editorconfig.org/

import (
	"fmt"
	
	"os/exec"
)

func installEditorconfig() {
	// Método 1: Descargar y extraer .tar.gz
	editorconfig_tar_url := "https://github.com/editorconfig/editorconfig-core-c/archive/refs/tags/v0.12.9.tar.gz"
	editorconfig_cmd_tar := exec.Command("curl", "-L", editorconfig_tar_url, "-o", "package.tar.gz")
	err := editorconfig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	editorconfig_zip_url := "https://github.com/editorconfig/editorconfig-core-c/archive/refs/tags/v0.12.9.zip"
	editorconfig_cmd_zip := exec.Command("curl", "-L", editorconfig_zip_url, "-o", "package.zip")
	err = editorconfig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	editorconfig_bin_url := "https://github.com/editorconfig/editorconfig-core-c/archive/refs/tags/v0.12.9.bin"
	editorconfig_cmd_bin := exec.Command("curl", "-L", editorconfig_bin_url, "-o", "binary.bin")
	err = editorconfig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	editorconfig_src_url := "https://github.com/editorconfig/editorconfig-core-c/archive/refs/tags/v0.12.9.src.tar.gz"
	editorconfig_cmd_src := exec.Command("curl", "-L", editorconfig_src_url, "-o", "source.tar.gz")
	err = editorconfig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	editorconfig_cmd_direct := exec.Command("./binary")
	err = editorconfig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
}
