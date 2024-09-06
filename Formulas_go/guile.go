package main

// Guile - GNU Ubiquitous Intelligent Language for Extensions
// Homepage: https://www.gnu.org/software/guile/

import (
	"fmt"
	
	"os/exec"
)

func installGuile() {
	// Método 1: Descargar y extraer .tar.gz
	guile_tar_url := "https://ftp.gnu.org/gnu/guile/guile-3.0.10.tar.xz"
	guile_cmd_tar := exec.Command("curl", "-L", guile_tar_url, "-o", "package.tar.gz")
	err := guile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	guile_zip_url := "https://ftp.gnu.org/gnu/guile/guile-3.0.10.tar.xz"
	guile_cmd_zip := exec.Command("curl", "-L", guile_zip_url, "-o", "package.zip")
	err = guile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	guile_bin_url := "https://ftp.gnu.org/gnu/guile/guile-3.0.10.tar.xz"
	guile_cmd_bin := exec.Command("curl", "-L", guile_bin_url, "-o", "binary.bin")
	err = guile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	guile_src_url := "https://ftp.gnu.org/gnu/guile/guile-3.0.10.tar.xz"
	guile_cmd_src := exec.Command("curl", "-L", guile_src_url, "-o", "source.tar.gz")
	err = guile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	guile_cmd_direct := exec.Command("./binary")
	err = guile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libunistring")
exec.Command("latte", "install", "libunistring")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
