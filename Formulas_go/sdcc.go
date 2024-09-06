package main

// Sdcc - ANSI C compiler for Intel 8051, Maxim 80DS390, and Zilog Z80
// Homepage: https://sdcc.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSdcc() {
	// Método 1: Descargar y extraer .tar.gz
	sdcc_tar_url := "https://downloads.sourceforge.net/project/sdcc/sdcc/4.4.0/sdcc-src-4.4.0.tar.bz2"
	sdcc_cmd_tar := exec.Command("curl", "-L", sdcc_tar_url, "-o", "package.tar.gz")
	err := sdcc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sdcc_zip_url := "https://downloads.sourceforge.net/project/sdcc/sdcc/4.4.0/sdcc-src-4.4.0.tar.bz2"
	sdcc_cmd_zip := exec.Command("curl", "-L", sdcc_zip_url, "-o", "package.zip")
	err = sdcc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sdcc_bin_url := "https://downloads.sourceforge.net/project/sdcc/sdcc/4.4.0/sdcc-src-4.4.0.tar.bz2"
	sdcc_cmd_bin := exec.Command("curl", "-L", sdcc_bin_url, "-o", "binary.bin")
	err = sdcc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sdcc_src_url := "https://downloads.sourceforge.net/project/sdcc/sdcc/4.4.0/sdcc-src-4.4.0.tar.bz2"
	sdcc_cmd_src := exec.Command("curl", "-L", sdcc_src_url, "-o", "source.tar.gz")
	err = sdcc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sdcc_cmd_direct := exec.Command("./binary")
	err = sdcc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: gputils")
exec.Command("latte", "install", "gputils")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
}
