package main

// Yasm - Modular BSD reimplementation of NASM
// Homepage: https://yasm.tortall.net/

import (
	"fmt"
	
	"os/exec"
)

func installYasm() {
	// Método 1: Descargar y extraer .tar.gz
	yasm_tar_url := "https://www.tortall.net/projects/yasm/releases/yasm-1.3.0.tar.gz"
	yasm_cmd_tar := exec.Command("curl", "-L", yasm_tar_url, "-o", "package.tar.gz")
	err := yasm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yasm_zip_url := "https://www.tortall.net/projects/yasm/releases/yasm-1.3.0.zip"
	yasm_cmd_zip := exec.Command("curl", "-L", yasm_zip_url, "-o", "package.zip")
	err = yasm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yasm_bin_url := "https://www.tortall.net/projects/yasm/releases/yasm-1.3.0.bin"
	yasm_cmd_bin := exec.Command("curl", "-L", yasm_bin_url, "-o", "binary.bin")
	err = yasm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yasm_src_url := "https://www.tortall.net/projects/yasm/releases/yasm-1.3.0.src.tar.gz"
	yasm_cmd_src := exec.Command("curl", "-L", yasm_src_url, "-o", "source.tar.gz")
	err = yasm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yasm_cmd_direct := exec.Command("./binary")
	err = yasm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
