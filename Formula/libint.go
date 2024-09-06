package main

// Libint - Library for computing electron repulsion integrals efficiently
// Homepage: https://github.com/evaleev/libint

import (
	"fmt"
	
	"os/exec"
)

func installLibint() {
	// Método 1: Descargar y extraer .tar.gz
	libint_tar_url := "https://github.com/evaleev/libint/archive/refs/tags/v2.9.0.tar.gz"
	libint_cmd_tar := exec.Command("curl", "-L", libint_tar_url, "-o", "package.tar.gz")
	err := libint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libint_zip_url := "https://github.com/evaleev/libint/archive/refs/tags/v2.9.0.zip"
	libint_cmd_zip := exec.Command("curl", "-L", libint_zip_url, "-o", "package.zip")
	err = libint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libint_bin_url := "https://github.com/evaleev/libint/archive/refs/tags/v2.9.0.bin"
	libint_cmd_bin := exec.Command("curl", "-L", libint_bin_url, "-o", "binary.bin")
	err = libint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libint_src_url := "https://github.com/evaleev/libint/archive/refs/tags/v2.9.0.src.tar.gz"
	libint_cmd_src := exec.Command("curl", "-L", libint_src_url, "-o", "source.tar.gz")
	err = libint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libint_cmd_direct := exec.Command("./binary")
	err = libint_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: eigen")
	exec.Command("latte", "install", "eigen").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
}
