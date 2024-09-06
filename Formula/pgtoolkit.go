package main

// Pgtoolkit - Tools for PostgreSQL maintenance
// Homepage: https://github.com/grayhemp/pgtoolkit

import (
	"fmt"
	
	"os/exec"
)

func installPgtoolkit() {
	// Método 1: Descargar y extraer .tar.gz
	pgtoolkit_tar_url := "https://github.com/grayhemp/pgtoolkit/archive/refs/tags/v1.0.2.tar.gz"
	pgtoolkit_cmd_tar := exec.Command("curl", "-L", pgtoolkit_tar_url, "-o", "package.tar.gz")
	err := pgtoolkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgtoolkit_zip_url := "https://github.com/grayhemp/pgtoolkit/archive/refs/tags/v1.0.2.zip"
	pgtoolkit_cmd_zip := exec.Command("curl", "-L", pgtoolkit_zip_url, "-o", "package.zip")
	err = pgtoolkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgtoolkit_bin_url := "https://github.com/grayhemp/pgtoolkit/archive/refs/tags/v1.0.2.bin"
	pgtoolkit_cmd_bin := exec.Command("curl", "-L", pgtoolkit_bin_url, "-o", "binary.bin")
	err = pgtoolkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgtoolkit_src_url := "https://github.com/grayhemp/pgtoolkit/archive/refs/tags/v1.0.2.src.tar.gz"
	pgtoolkit_cmd_src := exec.Command("curl", "-L", pgtoolkit_src_url, "-o", "source.tar.gz")
	err = pgtoolkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgtoolkit_cmd_direct := exec.Command("./binary")
	err = pgtoolkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
