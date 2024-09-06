package main

// Mpfrcx - Arbitrary precision library for arithmetic of univariate polynomials
// Homepage: https://www.multiprecision.org/mpfrcx/home.html

import (
	"fmt"
	
	"os/exec"
)

func installMpfrcx() {
	// Método 1: Descargar y extraer .tar.gz
	mpfrcx_tar_url := "https://www.multiprecision.org/downloads/mpfrcx-0.6.3.tar.gz"
	mpfrcx_cmd_tar := exec.Command("curl", "-L", mpfrcx_tar_url, "-o", "package.tar.gz")
	err := mpfrcx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpfrcx_zip_url := "https://www.multiprecision.org/downloads/mpfrcx-0.6.3.zip"
	mpfrcx_cmd_zip := exec.Command("curl", "-L", mpfrcx_zip_url, "-o", "package.zip")
	err = mpfrcx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpfrcx_bin_url := "https://www.multiprecision.org/downloads/mpfrcx-0.6.3.bin"
	mpfrcx_cmd_bin := exec.Command("curl", "-L", mpfrcx_bin_url, "-o", "binary.bin")
	err = mpfrcx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpfrcx_src_url := "https://www.multiprecision.org/downloads/mpfrcx-0.6.3.src.tar.gz"
	mpfrcx_cmd_src := exec.Command("curl", "-L", mpfrcx_src_url, "-o", "source.tar.gz")
	err = mpfrcx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpfrcx_cmd_direct := exec.Command("./binary")
	err = mpfrcx_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: libmpc")
exec.Command("latte", "install", "libmpc")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
}
