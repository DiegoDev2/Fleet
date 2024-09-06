package main

// Flex - Fast Lexical Analyzer, generates Scanners (tokenizers)
// Homepage: https://github.com/westes/flex

import (
	"fmt"
	
	"os/exec"
)

func installFlex() {
	// Método 1: Descargar y extraer .tar.gz
	flex_tar_url := "https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.tar.gz"
	flex_cmd_tar := exec.Command("curl", "-L", flex_tar_url, "-o", "package.tar.gz")
	err := flex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flex_zip_url := "https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.zip"
	flex_cmd_zip := exec.Command("curl", "-L", flex_zip_url, "-o", "package.zip")
	err = flex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flex_bin_url := "https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.bin"
	flex_cmd_bin := exec.Command("curl", "-L", flex_bin_url, "-o", "binary.bin")
	err = flex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flex_src_url := "https://github.com/westes/flex/releases/download/v2.6.4/flex-2.6.4.src.tar.gz"
	flex_cmd_src := exec.Command("curl", "-L", flex_src_url, "-o", "source.tar.gz")
	err = flex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flex_cmd_direct := exec.Command("./binary")
	err = flex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}
