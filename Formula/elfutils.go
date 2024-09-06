package main

// Elfutils - Libraries and utilities for handling ELF objects
// Homepage: https://fedorahosted.org/elfutils/

import (
	"fmt"
	
	"os/exec"
)

func installElfutils() {
	// Método 1: Descargar y extraer .tar.gz
	elfutils_tar_url := "https://sourceware.org/elfutils/ftp/0.191/elfutils-0.191.tar.bz2"
	elfutils_cmd_tar := exec.Command("curl", "-L", elfutils_tar_url, "-o", "package.tar.gz")
	err := elfutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	elfutils_zip_url := "https://sourceware.org/elfutils/ftp/0.191/elfutils-0.191.tar.bz2"
	elfutils_cmd_zip := exec.Command("curl", "-L", elfutils_zip_url, "-o", "package.zip")
	err = elfutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	elfutils_bin_url := "https://sourceware.org/elfutils/ftp/0.191/elfutils-0.191.tar.bz2"
	elfutils_cmd_bin := exec.Command("curl", "-L", elfutils_bin_url, "-o", "binary.bin")
	err = elfutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	elfutils_src_url := "https://sourceware.org/elfutils/ftp/0.191/elfutils-0.191.tar.bz2"
	elfutils_cmd_src := exec.Command("curl", "-L", elfutils_src_url, "-o", "source.tar.gz")
	err = elfutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	elfutils_cmd_direct := exec.Command("./binary")
	err = elfutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: m4")
	exec.Command("latte", "install", "m4").Run()
	fmt.Println("Instalando dependencia: bzip2")
	exec.Command("latte", "install", "bzip2").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
