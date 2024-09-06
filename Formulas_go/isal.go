package main

// IsaL - Intelligent Storage Acceleration Library
// Homepage: https://github.com/intel/isa-l

import (
	"fmt"
	
	"os/exec"
)

func installIsaL() {
	// Método 1: Descargar y extraer .tar.gz
	isal_tar_url := "https://github.com/intel/isa-l/archive/refs/tags/v2.31.0.tar.gz"
	isal_cmd_tar := exec.Command("curl", "-L", isal_tar_url, "-o", "package.tar.gz")
	err := isal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	isal_zip_url := "https://github.com/intel/isa-l/archive/refs/tags/v2.31.0.zip"
	isal_cmd_zip := exec.Command("curl", "-L", isal_zip_url, "-o", "package.zip")
	err = isal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	isal_bin_url := "https://github.com/intel/isa-l/archive/refs/tags/v2.31.0.bin"
	isal_cmd_bin := exec.Command("curl", "-L", isal_bin_url, "-o", "binary.bin")
	err = isal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	isal_src_url := "https://github.com/intel/isa-l/archive/refs/tags/v2.31.0.src.tar.gz"
	isal_cmd_src := exec.Command("curl", "-L", isal_src_url, "-o", "source.tar.gz")
	err = isal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	isal_cmd_direct := exec.Command("./binary")
	err = isal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
}
