package main

// Libextractor - Library to extract meta data from files
// Homepage: https://www.gnu.org/software/libextractor/

import (
	"fmt"
	
	"os/exec"
)

func installLibextractor() {
	// Método 1: Descargar y extraer .tar.gz
	libextractor_tar_url := "https://ftp.gnu.org/gnu/libextractor/libextractor-1.13.tar.gz"
	libextractor_cmd_tar := exec.Command("curl", "-L", libextractor_tar_url, "-o", "package.tar.gz")
	err := libextractor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libextractor_zip_url := "https://ftp.gnu.org/gnu/libextractor/libextractor-1.13.zip"
	libextractor_cmd_zip := exec.Command("curl", "-L", libextractor_zip_url, "-o", "package.zip")
	err = libextractor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libextractor_bin_url := "https://ftp.gnu.org/gnu/libextractor/libextractor-1.13.bin"
	libextractor_cmd_bin := exec.Command("curl", "-L", libextractor_bin_url, "-o", "binary.bin")
	err = libextractor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libextractor_src_url := "https://ftp.gnu.org/gnu/libextractor/libextractor-1.13.src.tar.gz"
	libextractor_cmd_src := exec.Command("curl", "-L", libextractor_src_url, "-o", "source.tar.gz")
	err = libextractor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libextractor_cmd_direct := exec.Command("./binary")
	err = libextractor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
