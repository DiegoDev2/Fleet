package main

// Apophenia - C library for statistical and scientific computing
// Homepage: https://github.com/b-k/apophenia

import (
	"fmt"
	
	"os/exec"
)

func installApophenia() {
	// Método 1: Descargar y extraer .tar.gz
	apophenia_tar_url := "https://github.com/b-k/apophenia/archive/refs/tags/v1.0.tar.gz"
	apophenia_cmd_tar := exec.Command("curl", "-L", apophenia_tar_url, "-o", "package.tar.gz")
	err := apophenia_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apophenia_zip_url := "https://github.com/b-k/apophenia/archive/refs/tags/v1.0.zip"
	apophenia_cmd_zip := exec.Command("curl", "-L", apophenia_zip_url, "-o", "package.zip")
	err = apophenia_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apophenia_bin_url := "https://github.com/b-k/apophenia/archive/refs/tags/v1.0.bin"
	apophenia_cmd_bin := exec.Command("curl", "-L", apophenia_bin_url, "-o", "binary.bin")
	err = apophenia_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apophenia_src_url := "https://github.com/b-k/apophenia/archive/refs/tags/v1.0.src.tar.gz"
	apophenia_cmd_src := exec.Command("curl", "-L", apophenia_src_url, "-o", "source.tar.gz")
	err = apophenia_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apophenia_cmd_direct := exec.Command("./binary")
	err = apophenia_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
