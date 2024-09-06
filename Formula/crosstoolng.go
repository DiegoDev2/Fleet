package main

// CrosstoolNg - Tool for building toolchains
// Homepage: https://crosstool-ng.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installCrosstoolNg() {
	// Método 1: Descargar y extraer .tar.gz
	crosstoolng_tar_url := "http://crosstool-ng.org/download/crosstool-ng/crosstool-ng-1.26.0.tar.xz"
	crosstoolng_cmd_tar := exec.Command("curl", "-L", crosstoolng_tar_url, "-o", "package.tar.gz")
	err := crosstoolng_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crosstoolng_zip_url := "http://crosstool-ng.org/download/crosstool-ng/crosstool-ng-1.26.0.tar.xz"
	crosstoolng_cmd_zip := exec.Command("curl", "-L", crosstoolng_zip_url, "-o", "package.zip")
	err = crosstoolng_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crosstoolng_bin_url := "http://crosstool-ng.org/download/crosstool-ng/crosstool-ng-1.26.0.tar.xz"
	crosstoolng_cmd_bin := exec.Command("curl", "-L", crosstoolng_bin_url, "-o", "binary.bin")
	err = crosstoolng_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crosstoolng_src_url := "http://crosstool-ng.org/download/crosstool-ng/crosstool-ng-1.26.0.tar.xz"
	crosstoolng_cmd_src := exec.Command("curl", "-L", crosstoolng_src_url, "-o", "source.tar.gz")
	err = crosstoolng_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crosstoolng_cmd_direct := exec.Command("./binary")
	err = crosstoolng_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: binutils")
	exec.Command("latte", "install", "binutils").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: lzip")
	exec.Command("latte", "install", "lzip").Run()
	fmt.Println("Instalando dependencia: m4")
	exec.Command("latte", "install", "m4").Run()
	fmt.Println("Instalando dependencia: ncurses")
	exec.Command("latte", "install", "ncurses").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: bash")
	exec.Command("latte", "install", "bash").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: grep")
	exec.Command("latte", "install", "grep").Run()
	fmt.Println("Instalando dependencia: make")
	exec.Command("latte", "install", "make").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
}
