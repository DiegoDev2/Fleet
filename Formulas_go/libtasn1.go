package main

// Libtasn1 - ASN.1 structure parser library
// Homepage: https://www.gnu.org/software/libtasn1/

import (
	"fmt"
	
	"os/exec"
)

func installLibtasn1() {
	// Método 1: Descargar y extraer .tar.gz
	libtasn1_tar_url := "https://ftp.gnu.org/gnu/libtasn1/libtasn1-4.19.0.tar.gz"
	libtasn1_cmd_tar := exec.Command("curl", "-L", libtasn1_tar_url, "-o", "package.tar.gz")
	err := libtasn1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtasn1_zip_url := "https://ftp.gnu.org/gnu/libtasn1/libtasn1-4.19.0.zip"
	libtasn1_cmd_zip := exec.Command("curl", "-L", libtasn1_zip_url, "-o", "package.zip")
	err = libtasn1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtasn1_bin_url := "https://ftp.gnu.org/gnu/libtasn1/libtasn1-4.19.0.bin"
	libtasn1_cmd_bin := exec.Command("curl", "-L", libtasn1_bin_url, "-o", "binary.bin")
	err = libtasn1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtasn1_src_url := "https://ftp.gnu.org/gnu/libtasn1/libtasn1-4.19.0.src.tar.gz"
	libtasn1_cmd_src := exec.Command("curl", "-L", libtasn1_src_url, "-o", "source.tar.gz")
	err = libtasn1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtasn1_cmd_direct := exec.Command("./binary")
	err = libtasn1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
