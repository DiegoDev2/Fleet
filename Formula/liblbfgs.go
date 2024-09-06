package main

// Liblbfgs - C library for limited-memory BFGS optimization algorithm
// Homepage: https://www.chokkan.org/software/liblbfgs

import (
	"fmt"
	
	"os/exec"
)

func installLiblbfgs() {
	// Método 1: Descargar y extraer .tar.gz
	liblbfgs_tar_url := "https://github.com/chokkan/liblbfgs/archive/refs/tags/v1.10.tar.gz"
	liblbfgs_cmd_tar := exec.Command("curl", "-L", liblbfgs_tar_url, "-o", "package.tar.gz")
	err := liblbfgs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liblbfgs_zip_url := "https://github.com/chokkan/liblbfgs/archive/refs/tags/v1.10.zip"
	liblbfgs_cmd_zip := exec.Command("curl", "-L", liblbfgs_zip_url, "-o", "package.zip")
	err = liblbfgs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liblbfgs_bin_url := "https://github.com/chokkan/liblbfgs/archive/refs/tags/v1.10.bin"
	liblbfgs_cmd_bin := exec.Command("curl", "-L", liblbfgs_bin_url, "-o", "binary.bin")
	err = liblbfgs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liblbfgs_src_url := "https://github.com/chokkan/liblbfgs/archive/refs/tags/v1.10.src.tar.gz"
	liblbfgs_cmd_src := exec.Command("curl", "-L", liblbfgs_src_url, "-o", "source.tar.gz")
	err = liblbfgs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liblbfgs_cmd_direct := exec.Command("./binary")
	err = liblbfgs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
