package main

// Fibjs - JavaScript on Fiber
// Homepage: https://fibjs.org/

import (
	"fmt"
	
	"os/exec"
)

func installFibjs() {
	// Método 1: Descargar y extraer .tar.gz
	fibjs_tar_url := "https://github.com/fibjs/fibjs/releases/download/v0.37.0/fullsrc.zip"
	fibjs_cmd_tar := exec.Command("curl", "-L", fibjs_tar_url, "-o", "package.tar.gz")
	err := fibjs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fibjs_zip_url := "https://github.com/fibjs/fibjs/releases/download/v0.37.0/fullsrc.zip"
	fibjs_cmd_zip := exec.Command("curl", "-L", fibjs_zip_url, "-o", "package.zip")
	err = fibjs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fibjs_bin_url := "https://github.com/fibjs/fibjs/releases/download/v0.37.0/fullsrc.zip"
	fibjs_cmd_bin := exec.Command("curl", "-L", fibjs_bin_url, "-o", "binary.bin")
	err = fibjs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fibjs_src_url := "https://github.com/fibjs/fibjs/releases/download/v0.37.0/fullsrc.zip"
	fibjs_cmd_src := exec.Command("curl", "-L", fibjs_src_url, "-o", "source.tar.gz")
	err = fibjs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fibjs_cmd_direct := exec.Command("./binary")
	err = fibjs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
}
