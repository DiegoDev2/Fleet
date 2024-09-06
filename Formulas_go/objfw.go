package main

// Objfw - Portable, lightweight framework for the Objective-C language
// Homepage: https://objfw.nil.im/

import (
	"fmt"
	
	"os/exec"
)

func installObjfw() {
	// Método 1: Descargar y extraer .tar.gz
	objfw_tar_url := "https://objfw.nil.im/downloads/objfw-1.1.7.tar.gz"
	objfw_cmd_tar := exec.Command("curl", "-L", objfw_tar_url, "-o", "package.tar.gz")
	err := objfw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	objfw_zip_url := "https://objfw.nil.im/downloads/objfw-1.1.7.zip"
	objfw_cmd_zip := exec.Command("curl", "-L", objfw_zip_url, "-o", "package.zip")
	err = objfw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	objfw_bin_url := "https://objfw.nil.im/downloads/objfw-1.1.7.bin"
	objfw_cmd_bin := exec.Command("curl", "-L", objfw_bin_url, "-o", "binary.bin")
	err = objfw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	objfw_src_url := "https://objfw.nil.im/downloads/objfw-1.1.7.src.tar.gz"
	objfw_cmd_src := exec.Command("curl", "-L", objfw_src_url, "-o", "source.tar.gz")
	err = objfw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	objfw_cmd_direct := exec.Command("./binary")
	err = objfw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
