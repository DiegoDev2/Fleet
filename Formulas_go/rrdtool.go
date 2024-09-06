package main

// Rrdtool - Round Robin Database
// Homepage: https://oss.oetiker.ch/rrdtool/index.en.html

import (
	"fmt"
	
	"os/exec"
)

func installRrdtool() {
	// Método 1: Descargar y extraer .tar.gz
	rrdtool_tar_url := "https://github.com/oetiker/rrdtool-1.x/releases/download/v1.9.0/rrdtool-1.9.0.tar.gz"
	rrdtool_cmd_tar := exec.Command("curl", "-L", rrdtool_tar_url, "-o", "package.tar.gz")
	err := rrdtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rrdtool_zip_url := "https://github.com/oetiker/rrdtool-1.x/releases/download/v1.9.0/rrdtool-1.9.0.zip"
	rrdtool_cmd_zip := exec.Command("curl", "-L", rrdtool_zip_url, "-o", "package.zip")
	err = rrdtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rrdtool_bin_url := "https://github.com/oetiker/rrdtool-1.x/releases/download/v1.9.0/rrdtool-1.9.0.bin"
	rrdtool_cmd_bin := exec.Command("curl", "-L", rrdtool_bin_url, "-o", "binary.bin")
	err = rrdtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rrdtool_src_url := "https://github.com/oetiker/rrdtool-1.x/releases/download/v1.9.0/rrdtool-1.9.0.src.tar.gz"
	rrdtool_cmd_src := exec.Command("curl", "-L", rrdtool_src_url, "-o", "source.tar.gz")
	err = rrdtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rrdtool_cmd_direct := exec.Command("./binary")
	err = rrdtool_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: groff")
exec.Command("latte", "install", "groff")
}
