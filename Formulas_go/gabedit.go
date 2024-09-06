package main

// Gabedit - GUI to computational chemistry packages like Gamess-US, Gaussian, etc.
// Homepage: https://gabedit.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGabedit() {
	// Método 1: Descargar y extraer .tar.gz
	gabedit_tar_url := "https://downloads.sourceforge.net/project/gabedit/gabedit/Gabedit251/GabeditSrc251.tar.gz"
	gabedit_cmd_tar := exec.Command("curl", "-L", gabedit_tar_url, "-o", "package.tar.gz")
	err := gabedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gabedit_zip_url := "https://downloads.sourceforge.net/project/gabedit/gabedit/Gabedit251/GabeditSrc251.zip"
	gabedit_cmd_zip := exec.Command("curl", "-L", gabedit_zip_url, "-o", "package.zip")
	err = gabedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gabedit_bin_url := "https://downloads.sourceforge.net/project/gabedit/gabedit/Gabedit251/GabeditSrc251.bin"
	gabedit_cmd_bin := exec.Command("curl", "-L", gabedit_bin_url, "-o", "binary.bin")
	err = gabedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gabedit_src_url := "https://downloads.sourceforge.net/project/gabedit/gabedit/Gabedit251/GabeditSrc251.src.tar.gz"
	gabedit_cmd_src := exec.Command("curl", "-L", gabedit_src_url, "-o", "source.tar.gz")
	err = gabedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gabedit_cmd_direct := exec.Command("./binary")
	err = gabedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
	fmt.Println("Instalando dependencia: gtkglext")
exec.Command("latte", "install", "gtkglext")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
