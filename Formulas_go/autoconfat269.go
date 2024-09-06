package main

// AutoconfAT269 - Automatic configure script builder
// Homepage: https://www.gnu.org/software/autoconf/

import (
	"fmt"
	
	"os/exec"
)

func installAutoconfAT269() {
	// Método 1: Descargar y extraer .tar.gz
	autoconfat269_tar_url := "https://ftp.gnu.org/gnu/autoconf/autoconf-2.69.tar.gz"
	autoconfat269_cmd_tar := exec.Command("curl", "-L", autoconfat269_tar_url, "-o", "package.tar.gz")
	err := autoconfat269_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autoconfat269_zip_url := "https://ftp.gnu.org/gnu/autoconf/autoconf-2.69.zip"
	autoconfat269_cmd_zip := exec.Command("curl", "-L", autoconfat269_zip_url, "-o", "package.zip")
	err = autoconfat269_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autoconfat269_bin_url := "https://ftp.gnu.org/gnu/autoconf/autoconf-2.69.bin"
	autoconfat269_cmd_bin := exec.Command("curl", "-L", autoconfat269_bin_url, "-o", "binary.bin")
	err = autoconfat269_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autoconfat269_src_url := "https://ftp.gnu.org/gnu/autoconf/autoconf-2.69.src.tar.gz"
	autoconfat269_cmd_src := exec.Command("curl", "-L", autoconfat269_src_url, "-o", "source.tar.gz")
	err = autoconfat269_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autoconfat269_cmd_direct := exec.Command("./binary")
	err = autoconfat269_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: m4")
exec.Command("latte", "install", "m4")
}
