package main

// ColladaDom - C++ library for loading and saving COLLADA data
// Homepage: https://www.khronos.org/collada/wiki/Portal:COLLADA_DOM

import (
	"fmt"
	
	"os/exec"
)

func installColladaDom() {
	// Método 1: Descargar y extraer .tar.gz
	colladadom_tar_url := "https://github.com/rdiankov/collada-dom/archive/refs/tags/v2.5.0.tar.gz"
	colladadom_cmd_tar := exec.Command("curl", "-L", colladadom_tar_url, "-o", "package.tar.gz")
	err := colladadom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	colladadom_zip_url := "https://github.com/rdiankov/collada-dom/archive/refs/tags/v2.5.0.zip"
	colladadom_cmd_zip := exec.Command("curl", "-L", colladadom_zip_url, "-o", "package.zip")
	err = colladadom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	colladadom_bin_url := "https://github.com/rdiankov/collada-dom/archive/refs/tags/v2.5.0.bin"
	colladadom_cmd_bin := exec.Command("curl", "-L", colladadom_bin_url, "-o", "binary.bin")
	err = colladadom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	colladadom_src_url := "https://github.com/rdiankov/collada-dom/archive/refs/tags/v2.5.0.src.tar.gz"
	colladadom_cmd_src := exec.Command("curl", "-L", colladadom_src_url, "-o", "source.tar.gz")
	err = colladadom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	colladadom_cmd_direct := exec.Command("./binary")
	err = colladadom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: minizip")
exec.Command("latte", "install", "minizip")
	fmt.Println("Instalando dependencia: uriparser")
exec.Command("latte", "install", "uriparser")
}
