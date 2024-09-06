package main

// Mprocs - Run multiple commands in parallel
// Homepage: https://github.com/pvolok/mprocs

import (
	"fmt"
	
	"os/exec"
)

func installMprocs() {
	// Método 1: Descargar y extraer .tar.gz
	mprocs_tar_url := "https://github.com/pvolok/mprocs/archive/refs/tags/v0.7.1.tar.gz"
	mprocs_cmd_tar := exec.Command("curl", "-L", mprocs_tar_url, "-o", "package.tar.gz")
	err := mprocs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mprocs_zip_url := "https://github.com/pvolok/mprocs/archive/refs/tags/v0.7.1.zip"
	mprocs_cmd_zip := exec.Command("curl", "-L", mprocs_zip_url, "-o", "package.zip")
	err = mprocs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mprocs_bin_url := "https://github.com/pvolok/mprocs/archive/refs/tags/v0.7.1.bin"
	mprocs_cmd_bin := exec.Command("curl", "-L", mprocs_bin_url, "-o", "binary.bin")
	err = mprocs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mprocs_src_url := "https://github.com/pvolok/mprocs/archive/refs/tags/v0.7.1.src.tar.gz"
	mprocs_cmd_src := exec.Command("curl", "-L", mprocs_src_url, "-o", "source.tar.gz")
	err = mprocs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mprocs_cmd_direct := exec.Command("./binary")
	err = mprocs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
}
