package main

// Tcc - Tiny C compiler
// Homepage: https://bellard.org/tcc/

import (
	"fmt"
	
	"os/exec"
)

func installTcc() {
	// Método 1: Descargar y extraer .tar.gz
	tcc_tar_url := "https://download.savannah.nongnu.org/releases/tinycc/tcc-0.9.27.tar.bz2"
	tcc_cmd_tar := exec.Command("curl", "-L", tcc_tar_url, "-o", "package.tar.gz")
	err := tcc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tcc_zip_url := "https://download.savannah.nongnu.org/releases/tinycc/tcc-0.9.27.tar.bz2"
	tcc_cmd_zip := exec.Command("curl", "-L", tcc_zip_url, "-o", "package.zip")
	err = tcc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tcc_bin_url := "https://download.savannah.nongnu.org/releases/tinycc/tcc-0.9.27.tar.bz2"
	tcc_cmd_bin := exec.Command("curl", "-L", tcc_bin_url, "-o", "binary.bin")
	err = tcc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tcc_src_url := "https://download.savannah.nongnu.org/releases/tinycc/tcc-0.9.27.tar.bz2"
	tcc_cmd_src := exec.Command("curl", "-L", tcc_src_url, "-o", "source.tar.gz")
	err = tcc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tcc_cmd_direct := exec.Command("./binary")
	err = tcc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
