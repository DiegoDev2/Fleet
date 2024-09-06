package main

// Mpfr - C library for multiple-precision floating-point computations
// Homepage: https://www.mpfr.org/

import (
	"fmt"
	
	"os/exec"
)

func installMpfr() {
	// Método 1: Descargar y extraer .tar.gz
	mpfr_tar_url := "https://ftp.gnu.org/gnu/mpfr/mpfr-4.2.1.tar.xz"
	mpfr_cmd_tar := exec.Command("curl", "-L", mpfr_tar_url, "-o", "package.tar.gz")
	err := mpfr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mpfr_zip_url := "https://ftp.gnu.org/gnu/mpfr/mpfr-4.2.1.tar.xz"
	mpfr_cmd_zip := exec.Command("curl", "-L", mpfr_zip_url, "-o", "package.zip")
	err = mpfr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mpfr_bin_url := "https://ftp.gnu.org/gnu/mpfr/mpfr-4.2.1.tar.xz"
	mpfr_cmd_bin := exec.Command("curl", "-L", mpfr_bin_url, "-o", "binary.bin")
	err = mpfr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mpfr_src_url := "https://ftp.gnu.org/gnu/mpfr/mpfr-4.2.1.tar.xz"
	mpfr_cmd_src := exec.Command("curl", "-L", mpfr_src_url, "-o", "source.tar.gz")
	err = mpfr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mpfr_cmd_direct := exec.Command("./binary")
	err = mpfr_cmd_direct.Run()
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
}
