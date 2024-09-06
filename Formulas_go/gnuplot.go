package main

// Gnuplot - Command-driven, interactive function plotting
// Homepage: http://www.gnuplot.info/

import (
	"fmt"
	
	"os/exec"
)

func installGnuplot() {
	// Método 1: Descargar y extraer .tar.gz
	gnuplot_tar_url := "https://downloads.sourceforge.net/project/gnuplot/gnuplot/6.0.1/gnuplot-6.0.1.tar.gz"
	gnuplot_cmd_tar := exec.Command("curl", "-L", gnuplot_tar_url, "-o", "package.tar.gz")
	err := gnuplot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnuplot_zip_url := "https://downloads.sourceforge.net/project/gnuplot/gnuplot/6.0.1/gnuplot-6.0.1.zip"
	gnuplot_cmd_zip := exec.Command("curl", "-L", gnuplot_zip_url, "-o", "package.zip")
	err = gnuplot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnuplot_bin_url := "https://downloads.sourceforge.net/project/gnuplot/gnuplot/6.0.1/gnuplot-6.0.1.bin"
	gnuplot_cmd_bin := exec.Command("curl", "-L", gnuplot_bin_url, "-o", "binary.bin")
	err = gnuplot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnuplot_src_url := "https://downloads.sourceforge.net/project/gnuplot/gnuplot/6.0.1/gnuplot-6.0.1.src.tar.gz"
	gnuplot_cmd_src := exec.Command("curl", "-L", gnuplot_src_url, "-o", "source.tar.gz")
	err = gnuplot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnuplot_cmd_direct := exec.Command("./binary")
	err = gnuplot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: gnu-sed")
exec.Command("latte", "install", "gnu-sed")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gd")
exec.Command("latte", "install", "gd")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libcerf")
exec.Command("latte", "install", "libcerf")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
