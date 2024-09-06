package main

// GnuApl - GNU implementation of the programming language APL
// Homepage: https://www.gnu.org/software/apl/

import (
	"fmt"
	
	"os/exec"
)

func installGnuApl() {
	// Método 1: Descargar y extraer .tar.gz
	gnuapl_tar_url := "https://ftp.gnu.org/gnu/apl/apl-1.9.tar.gz"
	gnuapl_cmd_tar := exec.Command("curl", "-L", gnuapl_tar_url, "-o", "package.tar.gz")
	err := gnuapl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnuapl_zip_url := "https://ftp.gnu.org/gnu/apl/apl-1.9.zip"
	gnuapl_cmd_zip := exec.Command("curl", "-L", gnuapl_zip_url, "-o", "package.zip")
	err = gnuapl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnuapl_bin_url := "https://ftp.gnu.org/gnu/apl/apl-1.9.bin"
	gnuapl_cmd_bin := exec.Command("curl", "-L", gnuapl_bin_url, "-o", "binary.bin")
	err = gnuapl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnuapl_src_url := "https://ftp.gnu.org/gnu/apl/apl-1.9.src.tar.gz"
	gnuapl_cmd_src := exec.Command("curl", "-L", gnuapl_src_url, "-o", "source.tar.gz")
	err = gnuapl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnuapl_cmd_direct := exec.Command("./binary")
	err = gnuapl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
}
