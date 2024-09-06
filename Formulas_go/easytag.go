package main

// EasyTag - Application for viewing and editing audio file tags
// Homepage: https://wiki.gnome.org/Apps/EasyTAG

import (
	"fmt"
	
	"os/exec"
)

func installEasyTag() {
	// Método 1: Descargar y extraer .tar.gz
	easytag_tar_url := "https://download.gnome.org/sources/easytag/2.4/easytag-2.4.3.tar.xz"
	easytag_cmd_tar := exec.Command("curl", "-L", easytag_tar_url, "-o", "package.tar.gz")
	err := easytag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	easytag_zip_url := "https://download.gnome.org/sources/easytag/2.4/easytag-2.4.3.tar.xz"
	easytag_cmd_zip := exec.Command("curl", "-L", easytag_zip_url, "-o", "package.zip")
	err = easytag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	easytag_bin_url := "https://download.gnome.org/sources/easytag/2.4/easytag-2.4.3.tar.xz"
	easytag_cmd_bin := exec.Command("curl", "-L", easytag_bin_url, "-o", "binary.bin")
	err = easytag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	easytag_src_url := "https://download.gnome.org/sources/easytag/2.4/easytag-2.4.3.tar.xz"
	easytag_cmd_src := exec.Command("curl", "-L", easytag_src_url, "-o", "source.tar.gz")
	err = easytag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	easytag_cmd_direct := exec.Command("./binary")
	err = easytag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: itstool")
exec.Command("latte", "install", "itstool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: hicolor-icon-theme")
exec.Command("latte", "install", "hicolor-icon-theme")
	fmt.Println("Instalando dependencia: id3lib")
exec.Command("latte", "install", "id3lib")
	fmt.Println("Instalando dependencia: libid3tag")
exec.Command("latte", "install", "libid3tag")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: speex")
exec.Command("latte", "install", "speex")
	fmt.Println("Instalando dependencia: taglib")
exec.Command("latte", "install", "taglib")
	fmt.Println("Instalando dependencia: wavpack")
exec.Command("latte", "install", "wavpack")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
