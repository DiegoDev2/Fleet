package main

// Lsof - Utility to list open files
// Homepage: https://github.com/lsof-org/lsof

import (
	"fmt"
	
	"os/exec"
)

func installLsof() {
	// Método 1: Descargar y extraer .tar.gz
	lsof_tar_url := "https://github.com/lsof-org/lsof/archive/refs/tags/4.99.3.tar.gz"
	lsof_cmd_tar := exec.Command("curl", "-L", lsof_tar_url, "-o", "package.tar.gz")
	err := lsof_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lsof_zip_url := "https://github.com/lsof-org/lsof/archive/refs/tags/4.99.3.zip"
	lsof_cmd_zip := exec.Command("curl", "-L", lsof_zip_url, "-o", "package.zip")
	err = lsof_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lsof_bin_url := "https://github.com/lsof-org/lsof/archive/refs/tags/4.99.3.bin"
	lsof_cmd_bin := exec.Command("curl", "-L", lsof_bin_url, "-o", "binary.bin")
	err = lsof_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lsof_src_url := "https://github.com/lsof-org/lsof/archive/refs/tags/4.99.3.src.tar.gz"
	lsof_cmd_src := exec.Command("curl", "-L", lsof_src_url, "-o", "source.tar.gz")
	err = lsof_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lsof_cmd_direct := exec.Command("./binary")
	err = lsof_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtirpc")
	exec.Command("latte", "install", "libtirpc").Run()
}
