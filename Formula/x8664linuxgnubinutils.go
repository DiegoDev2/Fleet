package main

// X8664LinuxGnuBinutils - GNU Binutils for x86_64-linux-gnu cross development
// Homepage: https://www.gnu.org/software/binutils/binutils.html

import (
	"fmt"
	
	"os/exec"
)

func installX8664LinuxGnuBinutils() {
	// Método 1: Descargar y extraer .tar.gz
	x8664linuxgnubinutils_tar_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	x8664linuxgnubinutils_cmd_tar := exec.Command("curl", "-L", x8664linuxgnubinutils_tar_url, "-o", "package.tar.gz")
	err := x8664linuxgnubinutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	x8664linuxgnubinutils_zip_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	x8664linuxgnubinutils_cmd_zip := exec.Command("curl", "-L", x8664linuxgnubinutils_zip_url, "-o", "package.zip")
	err = x8664linuxgnubinutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	x8664linuxgnubinutils_bin_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	x8664linuxgnubinutils_cmd_bin := exec.Command("curl", "-L", x8664linuxgnubinutils_bin_url, "-o", "binary.bin")
	err = x8664linuxgnubinutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	x8664linuxgnubinutils_src_url := "https://ftp.gnu.org/gnu/binutils/binutils-2.43.1.tar.bz2"
	x8664linuxgnubinutils_cmd_src := exec.Command("curl", "-L", x8664linuxgnubinutils_src_url, "-o", "source.tar.gz")
	err = x8664linuxgnubinutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	x8664linuxgnubinutils_cmd_direct := exec.Command("./binary")
	err = x8664linuxgnubinutils_cmd_direct.Run()
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
