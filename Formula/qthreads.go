package main

// Qthreads - Lightweight locality-aware user-level threading runtime
// Homepage: https://www.sandia.gov/qthreads/

import (
	"fmt"
	
	"os/exec"
)

func installQthreads() {
	// Método 1: Descargar y extraer .tar.gz
	qthreads_tar_url := "https://github.com/sandialabs/qthreads/archive/refs/tags/1.20.tar.gz"
	qthreads_cmd_tar := exec.Command("curl", "-L", qthreads_tar_url, "-o", "package.tar.gz")
	err := qthreads_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qthreads_zip_url := "https://github.com/sandialabs/qthreads/archive/refs/tags/1.20.zip"
	qthreads_cmd_zip := exec.Command("curl", "-L", qthreads_zip_url, "-o", "package.zip")
	err = qthreads_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qthreads_bin_url := "https://github.com/sandialabs/qthreads/archive/refs/tags/1.20.bin"
	qthreads_cmd_bin := exec.Command("curl", "-L", qthreads_bin_url, "-o", "binary.bin")
	err = qthreads_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qthreads_src_url := "https://github.com/sandialabs/qthreads/archive/refs/tags/1.20.src.tar.gz"
	qthreads_cmd_src := exec.Command("curl", "-L", qthreads_src_url, "-o", "source.tar.gz")
	err = qthreads_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qthreads_cmd_direct := exec.Command("./binary")
	err = qthreads_cmd_direct.Run()
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
