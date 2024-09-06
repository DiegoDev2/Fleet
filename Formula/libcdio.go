package main

// Libcdio - Compact Disc Input and Control Library
// Homepage: https://www.gnu.org/software/libcdio/

import (
	"fmt"
	
	"os/exec"
)

func installLibcdio() {
	// Método 1: Descargar y extraer .tar.gz
	libcdio_tar_url := "https://ftp.gnu.org/gnu/libcdio/libcdio-2.1.0.tar.bz2"
	libcdio_cmd_tar := exec.Command("curl", "-L", libcdio_tar_url, "-o", "package.tar.gz")
	err := libcdio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcdio_zip_url := "https://ftp.gnu.org/gnu/libcdio/libcdio-2.1.0.tar.bz2"
	libcdio_cmd_zip := exec.Command("curl", "-L", libcdio_zip_url, "-o", "package.zip")
	err = libcdio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcdio_bin_url := "https://ftp.gnu.org/gnu/libcdio/libcdio-2.1.0.tar.bz2"
	libcdio_cmd_bin := exec.Command("curl", "-L", libcdio_bin_url, "-o", "binary.bin")
	err = libcdio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcdio_src_url := "https://ftp.gnu.org/gnu/libcdio/libcdio-2.1.0.tar.bz2"
	libcdio_cmd_src := exec.Command("curl", "-L", libcdio_src_url, "-o", "source.tar.gz")
	err = libcdio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcdio_cmd_direct := exec.Command("./binary")
	err = libcdio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
