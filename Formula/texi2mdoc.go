package main

// Texi2mdoc - Convert Texinfo data to mdoc input
// Homepage: https://mandoc.bsd.lv/texi2mdoc/

import (
	"fmt"
	
	"os/exec"
)

func installTexi2mdoc() {
	// Método 1: Descargar y extraer .tar.gz
	texi2mdoc_tar_url := "https://mandoc.bsd.lv/texi2mdoc/snapshots/texi2mdoc-0.1.2.tgz"
	texi2mdoc_cmd_tar := exec.Command("curl", "-L", texi2mdoc_tar_url, "-o", "package.tar.gz")
	err := texi2mdoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	texi2mdoc_zip_url := "https://mandoc.bsd.lv/texi2mdoc/snapshots/texi2mdoc-0.1.2.tgz"
	texi2mdoc_cmd_zip := exec.Command("curl", "-L", texi2mdoc_zip_url, "-o", "package.zip")
	err = texi2mdoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	texi2mdoc_bin_url := "https://mandoc.bsd.lv/texi2mdoc/snapshots/texi2mdoc-0.1.2.tgz"
	texi2mdoc_cmd_bin := exec.Command("curl", "-L", texi2mdoc_bin_url, "-o", "binary.bin")
	err = texi2mdoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	texi2mdoc_src_url := "https://mandoc.bsd.lv/texi2mdoc/snapshots/texi2mdoc-0.1.2.tgz"
	texi2mdoc_cmd_src := exec.Command("curl", "-L", texi2mdoc_src_url, "-o", "source.tar.gz")
	err = texi2mdoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	texi2mdoc_cmd_direct := exec.Command("./binary")
	err = texi2mdoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
