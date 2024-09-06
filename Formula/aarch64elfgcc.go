package main

// Aarch64ElfGcc - GNU compiler collection for aarch64-elf
// Homepage: https://gcc.gnu.org

import (
	"fmt"
	
	"os/exec"
)

func installAarch64ElfGcc() {
	// Método 1: Descargar y extraer .tar.gz
	aarch64elfgcc_tar_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	aarch64elfgcc_cmd_tar := exec.Command("curl", "-L", aarch64elfgcc_tar_url, "-o", "package.tar.gz")
	err := aarch64elfgcc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aarch64elfgcc_zip_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	aarch64elfgcc_cmd_zip := exec.Command("curl", "-L", aarch64elfgcc_zip_url, "-o", "package.zip")
	err = aarch64elfgcc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aarch64elfgcc_bin_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	aarch64elfgcc_cmd_bin := exec.Command("curl", "-L", aarch64elfgcc_bin_url, "-o", "binary.bin")
	err = aarch64elfgcc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aarch64elfgcc_src_url := "https://ftp.gnu.org/gnu/gcc/gcc-14.2.0/gcc-14.2.0.tar.xz"
	aarch64elfgcc_cmd_src := exec.Command("curl", "-L", aarch64elfgcc_src_url, "-o", "source.tar.gz")
	err = aarch64elfgcc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aarch64elfgcc_cmd_direct := exec.Command("./binary")
	err = aarch64elfgcc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: aarch64-elf-binutils")
	exec.Command("latte", "install", "aarch64-elf-binutils").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: isl")
	exec.Command("latte", "install", "isl").Run()
	fmt.Println("Instalando dependencia: libmpc")
	exec.Command("latte", "install", "libmpc").Run()
	fmt.Println("Instalando dependencia: mpfr")
	exec.Command("latte", "install", "mpfr").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
