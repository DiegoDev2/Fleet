package main

// Flint - C library for number theory
// Homepage: https://flintlib.org/

import (
	"fmt"
	
	"os/exec"
)

func installFlint() {
	// Método 1: Descargar y extraer .tar.gz
	flint_tar_url := "https://github.com/flintlib/flint/releases/download/v3.1.3-p1/flint-3.1.3-p1.tar.gz"
	flint_cmd_tar := exec.Command("curl", "-L", flint_tar_url, "-o", "package.tar.gz")
	err := flint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flint_zip_url := "https://github.com/flintlib/flint/releases/download/v3.1.3-p1/flint-3.1.3-p1.zip"
	flint_cmd_zip := exec.Command("curl", "-L", flint_zip_url, "-o", "package.zip")
	err = flint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flint_bin_url := "https://github.com/flintlib/flint/releases/download/v3.1.3-p1/flint-3.1.3-p1.bin"
	flint_cmd_bin := exec.Command("curl", "-L", flint_bin_url, "-o", "binary.bin")
	err = flint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flint_src_url := "https://github.com/flintlib/flint/releases/download/v3.1.3-p1/flint-3.1.3-p1.src.tar.gz"
	flint_cmd_src := exec.Command("curl", "-L", flint_src_url, "-o", "source.tar.gz")
	err = flint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flint_cmd_direct := exec.Command("./binary")
	err = flint_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
}
