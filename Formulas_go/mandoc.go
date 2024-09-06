package main

// Mandoc - UNIX manpage compiler toolset
// Homepage: https://mandoc.bsd.lv/

import (
	"fmt"
	
	"os/exec"
)

func installMandoc() {
	// Método 1: Descargar y extraer .tar.gz
	mandoc_tar_url := "https://mandoc.bsd.lv/snapshots/mandoc-1.14.6.tar.gz"
	mandoc_cmd_tar := exec.Command("curl", "-L", mandoc_tar_url, "-o", "package.tar.gz")
	err := mandoc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mandoc_zip_url := "https://mandoc.bsd.lv/snapshots/mandoc-1.14.6.zip"
	mandoc_cmd_zip := exec.Command("curl", "-L", mandoc_zip_url, "-o", "package.zip")
	err = mandoc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mandoc_bin_url := "https://mandoc.bsd.lv/snapshots/mandoc-1.14.6.bin"
	mandoc_cmd_bin := exec.Command("curl", "-L", mandoc_bin_url, "-o", "binary.bin")
	err = mandoc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mandoc_src_url := "https://mandoc.bsd.lv/snapshots/mandoc-1.14.6.src.tar.gz"
	mandoc_cmd_src := exec.Command("curl", "-L", mandoc_src_url, "-o", "source.tar.gz")
	err = mandoc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mandoc_cmd_direct := exec.Command("./binary")
	err = mandoc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
