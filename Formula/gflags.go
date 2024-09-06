package main

// Gflags - Library for processing command-line flags
// Homepage: https://gflags.github.io/gflags/

import (
	"fmt"
	
	"os/exec"
)

func installGflags() {
	// Método 1: Descargar y extraer .tar.gz
	gflags_tar_url := "https://github.com/gflags/gflags/archive/refs/tags/v2.2.2.tar.gz"
	gflags_cmd_tar := exec.Command("curl", "-L", gflags_tar_url, "-o", "package.tar.gz")
	err := gflags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gflags_zip_url := "https://github.com/gflags/gflags/archive/refs/tags/v2.2.2.zip"
	gflags_cmd_zip := exec.Command("curl", "-L", gflags_zip_url, "-o", "package.zip")
	err = gflags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gflags_bin_url := "https://github.com/gflags/gflags/archive/refs/tags/v2.2.2.bin"
	gflags_cmd_bin := exec.Command("curl", "-L", gflags_bin_url, "-o", "binary.bin")
	err = gflags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gflags_src_url := "https://github.com/gflags/gflags/archive/refs/tags/v2.2.2.src.tar.gz"
	gflags_cmd_src := exec.Command("curl", "-L", gflags_src_url, "-o", "source.tar.gz")
	err = gflags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gflags_cmd_direct := exec.Command("./binary")
	err = gflags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
