package main

// Libidn2 - International domain name library (IDNA2008, Punycode and TR46)
// Homepage: https://www.gnu.org/software/libidn/#libidn2

import (
	"fmt"
	
	"os/exec"
)

func installLibidn2() {
	// Método 1: Descargar y extraer .tar.gz
	libidn2_tar_url := "https://ftp.gnu.org/gnu/libidn/libidn2-2.3.7.tar.gz"
	libidn2_cmd_tar := exec.Command("curl", "-L", libidn2_tar_url, "-o", "package.tar.gz")
	err := libidn2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libidn2_zip_url := "https://ftp.gnu.org/gnu/libidn/libidn2-2.3.7.zip"
	libidn2_cmd_zip := exec.Command("curl", "-L", libidn2_zip_url, "-o", "package.zip")
	err = libidn2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libidn2_bin_url := "https://ftp.gnu.org/gnu/libidn/libidn2-2.3.7.bin"
	libidn2_cmd_bin := exec.Command("curl", "-L", libidn2_bin_url, "-o", "binary.bin")
	err = libidn2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libidn2_src_url := "https://ftp.gnu.org/gnu/libidn/libidn2-2.3.7.src.tar.gz"
	libidn2_cmd_src := exec.Command("curl", "-L", libidn2_src_url, "-o", "source.tar.gz")
	err = libidn2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libidn2_cmd_direct := exec.Command("./binary")
	err = libidn2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gengetopt")
exec.Command("latte", "install", "gengetopt")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libunistring")
exec.Command("latte", "install", "libunistring")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
