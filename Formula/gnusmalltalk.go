package main

// GnuSmalltalk - Implementation of the Smalltalk language
// Homepage: https://www.gnu.org/software/smalltalk/

import (
	"fmt"
	
	"os/exec"
)

func installGnuSmalltalk() {
	// Método 1: Descargar y extraer .tar.gz
	gnusmalltalk_tar_url := "https://ftp.gnu.org/gnu/smalltalk/smalltalk-3.2.5.tar.xz"
	gnusmalltalk_cmd_tar := exec.Command("curl", "-L", gnusmalltalk_tar_url, "-o", "package.tar.gz")
	err := gnusmalltalk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnusmalltalk_zip_url := "https://ftp.gnu.org/gnu/smalltalk/smalltalk-3.2.5.tar.xz"
	gnusmalltalk_cmd_zip := exec.Command("curl", "-L", gnusmalltalk_zip_url, "-o", "package.zip")
	err = gnusmalltalk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnusmalltalk_bin_url := "https://ftp.gnu.org/gnu/smalltalk/smalltalk-3.2.5.tar.xz"
	gnusmalltalk_cmd_bin := exec.Command("curl", "-L", gnusmalltalk_bin_url, "-o", "binary.bin")
	err = gnusmalltalk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnusmalltalk_src_url := "https://ftp.gnu.org/gnu/smalltalk/smalltalk-3.2.5.tar.xz"
	gnusmalltalk_cmd_src := exec.Command("curl", "-L", gnusmalltalk_src_url, "-o", "source.tar.gz")
	err = gnusmalltalk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnusmalltalk_cmd_direct := exec.Command("./binary")
	err = gnusmalltalk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gdbm")
	exec.Command("latte", "install", "gdbm").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: libsigsegv")
	exec.Command("latte", "install", "libsigsegv").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
