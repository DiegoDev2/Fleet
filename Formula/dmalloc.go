package main

// Dmalloc - Debug versions of system memory management routines
// Homepage: https://dmalloc.com/

import (
	"fmt"
	
	"os/exec"
)

func installDmalloc() {
	// Método 1: Descargar y extraer .tar.gz
	dmalloc_tar_url := "https://dmalloc.com/releases/dmalloc-5.6.5.tgz"
	dmalloc_cmd_tar := exec.Command("curl", "-L", dmalloc_tar_url, "-o", "package.tar.gz")
	err := dmalloc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dmalloc_zip_url := "https://dmalloc.com/releases/dmalloc-5.6.5.tgz"
	dmalloc_cmd_zip := exec.Command("curl", "-L", dmalloc_zip_url, "-o", "package.zip")
	err = dmalloc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dmalloc_bin_url := "https://dmalloc.com/releases/dmalloc-5.6.5.tgz"
	dmalloc_cmd_bin := exec.Command("curl", "-L", dmalloc_bin_url, "-o", "binary.bin")
	err = dmalloc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dmalloc_src_url := "https://dmalloc.com/releases/dmalloc-5.6.5.tgz"
	dmalloc_cmd_src := exec.Command("curl", "-L", dmalloc_src_url, "-o", "source.tar.gz")
	err = dmalloc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dmalloc_cmd_direct := exec.Command("./binary")
	err = dmalloc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
