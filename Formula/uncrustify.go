package main

// Uncrustify - Source code beautifier
// Homepage: https://uncrustify.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installUncrustify() {
	// Método 1: Descargar y extraer .tar.gz
	uncrustify_tar_url := "https://github.com/uncrustify/uncrustify/archive/refs/tags/uncrustify-0.79.0.tar.gz"
	uncrustify_cmd_tar := exec.Command("curl", "-L", uncrustify_tar_url, "-o", "package.tar.gz")
	err := uncrustify_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uncrustify_zip_url := "https://github.com/uncrustify/uncrustify/archive/refs/tags/uncrustify-0.79.0.zip"
	uncrustify_cmd_zip := exec.Command("curl", "-L", uncrustify_zip_url, "-o", "package.zip")
	err = uncrustify_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uncrustify_bin_url := "https://github.com/uncrustify/uncrustify/archive/refs/tags/uncrustify-0.79.0.bin"
	uncrustify_cmd_bin := exec.Command("curl", "-L", uncrustify_bin_url, "-o", "binary.bin")
	err = uncrustify_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uncrustify_src_url := "https://github.com/uncrustify/uncrustify/archive/refs/tags/uncrustify-0.79.0.src.tar.gz"
	uncrustify_cmd_src := exec.Command("curl", "-L", uncrustify_src_url, "-o", "source.tar.gz")
	err = uncrustify_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uncrustify_cmd_direct := exec.Command("./binary")
	err = uncrustify_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
