package main

// Chars - Command-line tool to display information about unicode characters
// Homepage: https://github.com/boinkor-net/chars/

import (
	"fmt"
	
	"os/exec"
)

func installChars() {
	// Método 1: Descargar y extraer .tar.gz
	chars_tar_url := "https://github.com/boinkor-net/chars/archive/refs/tags/v0.7.0.tar.gz"
	chars_cmd_tar := exec.Command("curl", "-L", chars_tar_url, "-o", "package.tar.gz")
	err := chars_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chars_zip_url := "https://github.com/boinkor-net/chars/archive/refs/tags/v0.7.0.zip"
	chars_cmd_zip := exec.Command("curl", "-L", chars_zip_url, "-o", "package.zip")
	err = chars_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chars_bin_url := "https://github.com/boinkor-net/chars/archive/refs/tags/v0.7.0.bin"
	chars_cmd_bin := exec.Command("curl", "-L", chars_bin_url, "-o", "binary.bin")
	err = chars_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chars_src_url := "https://github.com/boinkor-net/chars/archive/refs/tags/v0.7.0.src.tar.gz"
	chars_cmd_src := exec.Command("curl", "-L", chars_src_url, "-o", "source.tar.gz")
	err = chars_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chars_cmd_direct := exec.Command("./binary")
	err = chars_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
