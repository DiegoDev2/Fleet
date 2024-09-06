package main

// AutoconfAT213 - Automatic configure script builder
// Homepage: https://www.gnu.org/software/autoconf/

import (
	"fmt"
	
	"os/exec"
)

func installAutoconfAT213() {
	// Método 1: Descargar y extraer .tar.gz
	autoconfat213_tar_url := "https://ftp.gnu.org/gnu/autoconf/autoconf-2.13.tar.gz"
	autoconfat213_cmd_tar := exec.Command("curl", "-L", autoconfat213_tar_url, "-o", "package.tar.gz")
	err := autoconfat213_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autoconfat213_zip_url := "https://ftp.gnu.org/gnu/autoconf/autoconf-2.13.zip"
	autoconfat213_cmd_zip := exec.Command("curl", "-L", autoconfat213_zip_url, "-o", "package.zip")
	err = autoconfat213_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autoconfat213_bin_url := "https://ftp.gnu.org/gnu/autoconf/autoconf-2.13.bin"
	autoconfat213_cmd_bin := exec.Command("curl", "-L", autoconfat213_bin_url, "-o", "binary.bin")
	err = autoconfat213_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autoconfat213_src_url := "https://ftp.gnu.org/gnu/autoconf/autoconf-2.13.src.tar.gz"
	autoconfat213_cmd_src := exec.Command("curl", "-L", autoconfat213_src_url, "-o", "source.tar.gz")
	err = autoconfat213_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autoconfat213_cmd_direct := exec.Command("./binary")
	err = autoconfat213_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
