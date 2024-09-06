package main

// CodeCli - Command-line interface built-in Visual Studio Code
// Homepage: https://github.com/microsoft/vscode

import (
	"fmt"
	
	"os/exec"
)

func installCodeCli() {
	// Método 1: Descargar y extraer .tar.gz
	codecli_tar_url := "https://github.com/microsoft/vscode/archive/refs/tags/1.93.0.tar.gz"
	codecli_cmd_tar := exec.Command("curl", "-L", codecli_tar_url, "-o", "package.tar.gz")
	err := codecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	codecli_zip_url := "https://github.com/microsoft/vscode/archive/refs/tags/1.93.0.zip"
	codecli_cmd_zip := exec.Command("curl", "-L", codecli_zip_url, "-o", "package.zip")
	err = codecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	codecli_bin_url := "https://github.com/microsoft/vscode/archive/refs/tags/1.93.0.bin"
	codecli_cmd_bin := exec.Command("curl", "-L", codecli_bin_url, "-o", "binary.bin")
	err = codecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	codecli_src_url := "https://github.com/microsoft/vscode/archive/refs/tags/1.93.0.src.tar.gz"
	codecli_cmd_src := exec.Command("curl", "-L", codecli_src_url, "-o", "source.tar.gz")
	err = codecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	codecli_cmd_direct := exec.Command("./binary")
	err = codecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
