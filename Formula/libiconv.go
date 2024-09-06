package main

// Libiconv - Conversion library
// Homepage: https://www.gnu.org/software/libiconv/

import (
	"fmt"
	
	"os/exec"
)

func installLibiconv() {
	// Método 1: Descargar y extraer .tar.gz
	libiconv_tar_url := "https://ftp.gnu.org/gnu/libiconv/libiconv-1.17.tar.gz"
	libiconv_cmd_tar := exec.Command("curl", "-L", libiconv_tar_url, "-o", "package.tar.gz")
	err := libiconv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libiconv_zip_url := "https://ftp.gnu.org/gnu/libiconv/libiconv-1.17.zip"
	libiconv_cmd_zip := exec.Command("curl", "-L", libiconv_zip_url, "-o", "package.zip")
	err = libiconv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libiconv_bin_url := "https://ftp.gnu.org/gnu/libiconv/libiconv-1.17.bin"
	libiconv_cmd_bin := exec.Command("curl", "-L", libiconv_bin_url, "-o", "binary.bin")
	err = libiconv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libiconv_src_url := "https://ftp.gnu.org/gnu/libiconv/libiconv-1.17.src.tar.gz"
	libiconv_cmd_src := exec.Command("curl", "-L", libiconv_src_url, "-o", "source.tar.gz")
	err = libiconv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libiconv_cmd_direct := exec.Command("./binary")
	err = libiconv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
