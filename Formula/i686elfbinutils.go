package main

// I686ElfBinutils - GNU Binutils for i686-elf cross development
// Homepage: https://www.gnu.org/software/binutils/

import (
	"fmt"
	
	"os/exec"
)

func installI686ElfBinutils() {
	// Método 1: Descargar y extraer .tar.gz
	i686elfbinutils_tar_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	i686elfbinutils_cmd_tar := exec.Command("curl", "-L", i686elfbinutils_tar_url, "-o", "package.tar.gz")
	err := i686elfbinutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	i686elfbinutils_zip_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	i686elfbinutils_cmd_zip := exec.Command("curl", "-L", i686elfbinutils_zip_url, "-o", "package.zip")
	err = i686elfbinutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	i686elfbinutils_bin_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	i686elfbinutils_cmd_bin := exec.Command("curl", "-L", i686elfbinutils_bin_url, "-o", "binary.bin")
	err = i686elfbinutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	i686elfbinutils_src_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	i686elfbinutils_cmd_src := exec.Command("curl", "-L", i686elfbinutils_src_url, "-o", "source.tar.gz")
	err = i686elfbinutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	i686elfbinutils_cmd_direct := exec.Command("./binary")
	err = i686elfbinutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
}
