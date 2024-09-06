package main

// Gnucobol - COBOL85-202x compiler supporting lots of dialect specific extensions
// Homepage: https://gnucobol.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installGnucobol() {
	// Método 1: Descargar y extraer .tar.gz
	gnucobol_tar_url := "https://ftp.gnu.org/gnu/gnucobol/gnucobol-3.2.tar.xz"
	gnucobol_cmd_tar := exec.Command("curl", "-L", gnucobol_tar_url, "-o", "package.tar.gz")
	err := gnucobol_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnucobol_zip_url := "https://ftp.gnu.org/gnu/gnucobol/gnucobol-3.2.tar.xz"
	gnucobol_cmd_zip := exec.Command("curl", "-L", gnucobol_zip_url, "-o", "package.zip")
	err = gnucobol_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnucobol_bin_url := "https://ftp.gnu.org/gnu/gnucobol/gnucobol-3.2.tar.xz"
	gnucobol_cmd_bin := exec.Command("curl", "-L", gnucobol_bin_url, "-o", "binary.bin")
	err = gnucobol_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnucobol_src_url := "https://ftp.gnu.org/gnu/gnucobol/gnucobol-3.2.tar.xz"
	gnucobol_cmd_src := exec.Command("curl", "-L", gnucobol_src_url, "-o", "source.tar.gz")
	err = gnucobol_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnucobol_cmd_direct := exec.Command("./binary")
	err = gnucobol_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: lmdb")
	exec.Command("latte", "install", "lmdb").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: berkeley-db")
	exec.Command("latte", "install", "berkeley-db").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
}
