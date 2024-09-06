package main

// Taglib - Audio metadata library
// Homepage: https://taglib.org/

import (
	"fmt"
	
	"os/exec"
)

func installTaglib() {
	// Método 1: Descargar y extraer .tar.gz
	taglib_tar_url := "https://taglib.github.io/releases/taglib-1.13.1.tar.gz"
	taglib_cmd_tar := exec.Command("curl", "-L", taglib_tar_url, "-o", "package.tar.gz")
	err := taglib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	taglib_zip_url := "https://taglib.github.io/releases/taglib-1.13.1.zip"
	taglib_cmd_zip := exec.Command("curl", "-L", taglib_zip_url, "-o", "package.zip")
	err = taglib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	taglib_bin_url := "https://taglib.github.io/releases/taglib-1.13.1.bin"
	taglib_cmd_bin := exec.Command("curl", "-L", taglib_bin_url, "-o", "binary.bin")
	err = taglib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	taglib_src_url := "https://taglib.github.io/releases/taglib-1.13.1.src.tar.gz"
	taglib_cmd_src := exec.Command("curl", "-L", taglib_src_url, "-o", "source.tar.gz")
	err = taglib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	taglib_cmd_direct := exec.Command("./binary")
	err = taglib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
