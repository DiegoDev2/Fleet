package main

// X8664ElfGcc - GNU compiler collection for x86_64-elf
// Homepage: https://gcc.gnu.org

import (
	"fmt"
	
	"os/exec"
)

func installX8664ElfGcc() {
	// Método 1: Descargar y extraer .tar.gz
	x8664elfgcc_tar_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	x8664elfgcc_cmd_tar := exec.Command("curl", "-L", x8664elfgcc_tar_url, "-o", "package.tar.gz")
	err := x8664elfgcc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	x8664elfgcc_zip_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	x8664elfgcc_cmd_zip := exec.Command("curl", "-L", x8664elfgcc_zip_url, "-o", "package.zip")
	err = x8664elfgcc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	x8664elfgcc_bin_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	x8664elfgcc_cmd_bin := exec.Command("curl", "-L", x8664elfgcc_bin_url, "-o", "binary.bin")
	err = x8664elfgcc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	x8664elfgcc_src_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	x8664elfgcc_cmd_src := exec.Command("curl", "-L", x8664elfgcc_src_url, "-o", "source.tar.gz")
	err = x8664elfgcc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	x8664elfgcc_cmd_direct := exec.Command("./binary")
	err = x8664elfgcc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: libmpc")
exec.Command("latte", "install", "libmpc")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: x86_64-elf-binutils")
exec.Command("latte", "install", "x86_64-elf-binutils")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
