package main

// Libunistring - C string library for manipulating Unicode strings
// Homepage: https://www.gnu.org/software/libunistring/

import (
	"fmt"
	
	"os/exec"
)

func installLibunistring() {
	// Método 1: Descargar y extraer .tar.gz
	libunistring_tar_url := "https://ftp.gnu.org/gnu/libunistring/libunistring-1.2.tar.gz"
	libunistring_cmd_tar := exec.Command("curl", "-L", libunistring_tar_url, "-o", "package.tar.gz")
	err := libunistring_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libunistring_zip_url := "https://ftp.gnu.org/gnu/libunistring/libunistring-1.2.zip"
	libunistring_cmd_zip := exec.Command("curl", "-L", libunistring_zip_url, "-o", "package.zip")
	err = libunistring_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libunistring_bin_url := "https://ftp.gnu.org/gnu/libunistring/libunistring-1.2.bin"
	libunistring_cmd_bin := exec.Command("curl", "-L", libunistring_bin_url, "-o", "binary.bin")
	err = libunistring_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libunistring_src_url := "https://ftp.gnu.org/gnu/libunistring/libunistring-1.2.src.tar.gz"
	libunistring_cmd_src := exec.Command("curl", "-L", libunistring_src_url, "-o", "source.tar.gz")
	err = libunistring_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libunistring_cmd_direct := exec.Command("./binary")
	err = libunistring_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
