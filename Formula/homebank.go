package main

// Homebank - Manage your personal accounts at home
// Homepage: http://homebank.free.fr

import (
	"fmt"
	
	"os/exec"
)

func installHomebank() {
	// Método 1: Descargar y extraer .tar.gz
	homebank_tar_url := "https://deb.debian.org/debian/pool/main/h/homebank/homebank_5.8.2.orig.tar.gz"
	homebank_cmd_tar := exec.Command("curl", "-L", homebank_tar_url, "-o", "package.tar.gz")
	err := homebank_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	homebank_zip_url := "https://deb.debian.org/debian/pool/main/h/homebank/homebank_5.8.2.orig.zip"
	homebank_cmd_zip := exec.Command("curl", "-L", homebank_zip_url, "-o", "package.zip")
	err = homebank_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	homebank_bin_url := "https://deb.debian.org/debian/pool/main/h/homebank/homebank_5.8.2.orig.bin"
	homebank_cmd_bin := exec.Command("curl", "-L", homebank_bin_url, "-o", "binary.bin")
	err = homebank_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	homebank_src_url := "https://deb.debian.org/debian/pool/main/h/homebank/homebank_5.8.2.orig.src.tar.gz"
	homebank_cmd_src := exec.Command("curl", "-L", homebank_src_url, "-o", "source.tar.gz")
	err = homebank_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	homebank_cmd_direct := exec.Command("./binary")
	err = homebank_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: intltool")
	exec.Command("latte", "install", "intltool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
	exec.Command("latte", "install", "adwaita-icon-theme").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gdk-pixbuf")
	exec.Command("latte", "install", "gdk-pixbuf").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: hicolor-icon-theme")
	exec.Command("latte", "install", "hicolor-icon-theme").Run()
	fmt.Println("Instalando dependencia: libofx")
	exec.Command("latte", "install", "libofx").Run()
	fmt.Println("Instalando dependencia: libsoup")
	exec.Command("latte", "install", "libsoup").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: at-spi2-core")
	exec.Command("latte", "install", "at-spi2-core").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: perl-xml-parser")
	exec.Command("latte", "install", "perl-xml-parser").Run()
}
