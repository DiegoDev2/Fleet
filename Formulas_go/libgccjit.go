package main

// Libgccjit - JIT library for the GNU compiler collection
// Homepage: https://gcc.gnu.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibgccjit() {
	// Método 1: Descargar y extraer .tar.gz
	libgccjit_tar_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	libgccjit_cmd_tar := exec.Command("curl", "-L", libgccjit_tar_url, "-o", "package.tar.gz")
	err := libgccjit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgccjit_zip_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	libgccjit_cmd_zip := exec.Command("curl", "-L", libgccjit_zip_url, "-o", "package.zip")
	err = libgccjit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgccjit_bin_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	libgccjit_cmd_bin := exec.Command("curl", "-L", libgccjit_bin_url, "-o", "binary.bin")
	err = libgccjit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgccjit_src_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	libgccjit_cmd_src := exec.Command("curl", "-L", libgccjit_src_url, "-o", "source.tar.gz")
	err = libgccjit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgccjit_cmd_direct := exec.Command("./binary")
	err = libgccjit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: isl")
exec.Command("latte", "install", "isl")
	fmt.Println("Instalando dependencia: libmpc")
exec.Command("latte", "install", "libmpc")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
}
