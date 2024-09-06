package main

// Vlang - V programming language
// Homepage: https://vlang.io

import (
	"fmt"
	
	"os/exec"
)

func installVlang() {
	// Método 1: Descargar y extraer .tar.gz
	vlang_tar_url := "https://github.com/vlang/v/archive/refs/tags/0.4.7.tar.gz"
	vlang_cmd_tar := exec.Command("curl", "-L", vlang_tar_url, "-o", "package.tar.gz")
	err := vlang_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vlang_zip_url := "https://github.com/vlang/v/archive/refs/tags/0.4.7.zip"
	vlang_cmd_zip := exec.Command("curl", "-L", vlang_zip_url, "-o", "package.zip")
	err = vlang_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vlang_bin_url := "https://github.com/vlang/v/archive/refs/tags/0.4.7.bin"
	vlang_cmd_bin := exec.Command("curl", "-L", vlang_bin_url, "-o", "binary.bin")
	err = vlang_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vlang_src_url := "https://github.com/vlang/v/archive/refs/tags/0.4.7.src.tar.gz"
	vlang_cmd_src := exec.Command("curl", "-L", vlang_src_url, "-o", "source.tar.gz")
	err = vlang_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vlang_cmd_direct := exec.Command("./binary")
	err = vlang_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bdw-gc")
	exec.Command("latte", "install", "bdw-gc").Run()
}
