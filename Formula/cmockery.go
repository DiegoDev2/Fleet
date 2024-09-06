package main

// Cmockery - Unit testing and mocking library for C
// Homepage: https://github.com/google/cmockery

import (
	"fmt"
	
	"os/exec"
)

func installCmockery() {
	// Método 1: Descargar y extraer .tar.gz
	cmockery_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cmockery/cmockery-0.1.2.tar.gz"
	cmockery_cmd_tar := exec.Command("curl", "-L", cmockery_tar_url, "-o", "package.tar.gz")
	err := cmockery_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmockery_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cmockery/cmockery-0.1.2.zip"
	cmockery_cmd_zip := exec.Command("curl", "-L", cmockery_zip_url, "-o", "package.zip")
	err = cmockery_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmockery_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cmockery/cmockery-0.1.2.bin"
	cmockery_cmd_bin := exec.Command("curl", "-L", cmockery_bin_url, "-o", "binary.bin")
	err = cmockery_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmockery_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/cmockery/cmockery-0.1.2.src.tar.gz"
	cmockery_cmd_src := exec.Command("curl", "-L", cmockery_src_url, "-o", "source.tar.gz")
	err = cmockery_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmockery_cmd_direct := exec.Command("./binary")
	err = cmockery_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
