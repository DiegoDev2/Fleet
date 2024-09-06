package main

// BoreCli - Modern, simple TCP tunnel in Rust that exposes local ports to a remote server
// Homepage: https://github.com/ekzhang/bore

import (
	"fmt"
	
	"os/exec"
)

func installBoreCli() {
	// Método 1: Descargar y extraer .tar.gz
	borecli_tar_url := "https://github.com/ekzhang/bore/archive/refs/tags/v0.5.1.tar.gz"
	borecli_cmd_tar := exec.Command("curl", "-L", borecli_tar_url, "-o", "package.tar.gz")
	err := borecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	borecli_zip_url := "https://github.com/ekzhang/bore/archive/refs/tags/v0.5.1.zip"
	borecli_cmd_zip := exec.Command("curl", "-L", borecli_zip_url, "-o", "package.zip")
	err = borecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	borecli_bin_url := "https://github.com/ekzhang/bore/archive/refs/tags/v0.5.1.bin"
	borecli_cmd_bin := exec.Command("curl", "-L", borecli_bin_url, "-o", "binary.bin")
	err = borecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	borecli_src_url := "https://github.com/ekzhang/bore/archive/refs/tags/v0.5.1.src.tar.gz"
	borecli_cmd_src := exec.Command("curl", "-L", borecli_src_url, "-o", "source.tar.gz")
	err = borecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	borecli_cmd_direct := exec.Command("./binary")
	err = borecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
