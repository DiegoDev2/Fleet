package main

// Aarch64ElfBinutils - GNU Binutils for aarch64-elf cross development
// Homepage: https://www.gnu.org/software/binutils/

import (
	"fmt"
	
	"os/exec"
)

func installAarch64ElfBinutils() {
	// Método 1: Descargar y extraer .tar.gz
	aarch64elfbinutils_tar_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	aarch64elfbinutils_cmd_tar := exec.Command("curl", "-L", aarch64elfbinutils_tar_url, "-o", "package.tar.gz")
	err := aarch64elfbinutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aarch64elfbinutils_zip_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	aarch64elfbinutils_cmd_zip := exec.Command("curl", "-L", aarch64elfbinutils_zip_url, "-o", "package.zip")
	err = aarch64elfbinutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aarch64elfbinutils_bin_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	aarch64elfbinutils_cmd_bin := exec.Command("curl", "-L", aarch64elfbinutils_bin_url, "-o", "binary.bin")
	err = aarch64elfbinutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aarch64elfbinutils_src_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	aarch64elfbinutils_cmd_src := exec.Command("curl", "-L", aarch64elfbinutils_src_url, "-o", "source.tar.gz")
	err = aarch64elfbinutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aarch64elfbinutils_cmd_direct := exec.Command("./binary")
	err = aarch64elfbinutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
