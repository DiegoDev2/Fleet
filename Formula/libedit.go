package main

// Libedit - BSD-style licensed readline alternative
// Homepage: https://thrysoee.dk/editline/

import (
	"fmt"
	
	"os/exec"
)

func installLibedit() {
	// Método 1: Descargar y extraer .tar.gz
	libedit_tar_url := "https://thrysoee.dk/editline/libedit-20240808-3.1.tar.gz"
	libedit_cmd_tar := exec.Command("curl", "-L", libedit_tar_url, "-o", "package.tar.gz")
	err := libedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libedit_zip_url := "https://thrysoee.dk/editline/libedit-20240808-3.1.zip"
	libedit_cmd_zip := exec.Command("curl", "-L", libedit_zip_url, "-o", "package.zip")
	err = libedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libedit_bin_url := "https://thrysoee.dk/editline/libedit-20240808-3.1.bin"
	libedit_cmd_bin := exec.Command("curl", "-L", libedit_bin_url, "-o", "binary.bin")
	err = libedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libedit_src_url := "https://thrysoee.dk/editline/libedit-20240808-3.1.src.tar.gz"
	libedit_cmd_src := exec.Command("curl", "-L", libedit_src_url, "-o", "source.tar.gz")
	err = libedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libedit_cmd_direct := exec.Command("./binary")
	err = libedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
