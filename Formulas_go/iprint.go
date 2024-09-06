package main

// Iprint - Provides a print_one function
// Homepage: https://www.samba.org/ftp/unpacked/junkcode/i.c

import (
	"fmt"
	
	"os/exec"
)

func installIprint() {
	// Método 1: Descargar y extraer .tar.gz
	iprint_tar_url := "http://archive.ubuntu.com/ubuntu/pool/universe/i/iprint/iprint_1.3.orig.tar.gz"
	iprint_cmd_tar := exec.Command("curl", "-L", iprint_tar_url, "-o", "package.tar.gz")
	err := iprint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iprint_zip_url := "http://archive.ubuntu.com/ubuntu/pool/universe/i/iprint/iprint_1.3.orig.zip"
	iprint_cmd_zip := exec.Command("curl", "-L", iprint_zip_url, "-o", "package.zip")
	err = iprint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iprint_bin_url := "http://archive.ubuntu.com/ubuntu/pool/universe/i/iprint/iprint_1.3.orig.bin"
	iprint_cmd_bin := exec.Command("curl", "-L", iprint_bin_url, "-o", "binary.bin")
	err = iprint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iprint_src_url := "http://archive.ubuntu.com/ubuntu/pool/universe/i/iprint/iprint_1.3.orig.src.tar.gz"
	iprint_cmd_src := exec.Command("curl", "-L", iprint_src_url, "-o", "source.tar.gz")
	err = iprint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iprint_cmd_direct := exec.Command("./binary")
	err = iprint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
