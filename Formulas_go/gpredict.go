package main

// Gpredict - Real-time satellite tracking/prediction application
// Homepage: https://gpredict.oz9aec.net/

import (
	"fmt"
	
	"os/exec"
)

func installGpredict() {
	// Método 1: Descargar y extraer .tar.gz
	gpredict_tar_url := "https://github.com/csete/gpredict/releases/download/v2.2.1/gpredict-2.2.1.tar.bz2"
	gpredict_cmd_tar := exec.Command("curl", "-L", gpredict_tar_url, "-o", "package.tar.gz")
	err := gpredict_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpredict_zip_url := "https://github.com/csete/gpredict/releases/download/v2.2.1/gpredict-2.2.1.tar.bz2"
	gpredict_cmd_zip := exec.Command("curl", "-L", gpredict_zip_url, "-o", "package.zip")
	err = gpredict_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpredict_bin_url := "https://github.com/csete/gpredict/releases/download/v2.2.1/gpredict-2.2.1.tar.bz2"
	gpredict_cmd_bin := exec.Command("curl", "-L", gpredict_bin_url, "-o", "binary.bin")
	err = gpredict_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpredict_src_url := "https://github.com/csete/gpredict/releases/download/v2.2.1/gpredict-2.2.1.tar.bz2"
	gpredict_cmd_src := exec.Command("curl", "-L", gpredict_src_url, "-o", "source.tar.gz")
	err = gpredict_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpredict_cmd_direct := exec.Command("./binary")
	err = gpredict_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: goocanvas")
exec.Command("latte", "install", "goocanvas")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: hamlib")
exec.Command("latte", "install", "hamlib")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
