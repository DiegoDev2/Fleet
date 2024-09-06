package main

// Webkitgtk - GTK interface to WebKit
// Homepage: https://webkitgtk.org

import (
	"fmt"
	
	"os/exec"
)

func installWebkitgtk() {
	// Método 1: Descargar y extraer .tar.gz
	webkitgtk_tar_url := "https://webkitgtk.org/releases/webkitgtk-2.44.3.tar.xz"
	webkitgtk_cmd_tar := exec.Command("curl", "-L", webkitgtk_tar_url, "-o", "package.tar.gz")
	err := webkitgtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	webkitgtk_zip_url := "https://webkitgtk.org/releases/webkitgtk-2.44.3.tar.xz"
	webkitgtk_cmd_zip := exec.Command("curl", "-L", webkitgtk_zip_url, "-o", "package.zip")
	err = webkitgtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	webkitgtk_bin_url := "https://webkitgtk.org/releases/webkitgtk-2.44.3.tar.xz"
	webkitgtk_cmd_bin := exec.Command("curl", "-L", webkitgtk_bin_url, "-o", "binary.bin")
	err = webkitgtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	webkitgtk_src_url := "https://webkitgtk.org/releases/webkitgtk-2.44.3.tar.xz"
	webkitgtk_cmd_src := exec.Command("curl", "-L", webkitgtk_src_url, "-o", "source.tar.gz")
	err = webkitgtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	webkitgtk_cmd_direct := exec.Command("./binary")
	err = webkitgtk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: gperf")
exec.Command("latte", "install", "gperf")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
	fmt.Println("Instalando dependencia: unifdef")
exec.Command("latte", "install", "unifdef")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: enchant")
exec.Command("latte", "install", "enchant")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gstreamer")
exec.Command("latte", "install", "gstreamer")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: jpeg-xl")
exec.Command("latte", "install", "jpeg-xl")
	fmt.Println("Instalando dependencia: libavif")
exec.Command("latte", "install", "libavif")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libnotify")
exec.Command("latte", "install", "libnotify")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsecret")
exec.Command("latte", "install", "libsecret")
	fmt.Println("Instalando dependencia: libsoup")
exec.Command("latte", "install", "libsoup")
	fmt.Println("Instalando dependencia: libwpe")
exec.Command("latte", "install", "libwpe")
	fmt.Println("Instalando dependencia: libxcomposite")
exec.Command("latte", "install", "libxcomposite")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: libxslt")
exec.Command("latte", "install", "libxslt")
	fmt.Println("Instalando dependencia: libxt")
exec.Command("latte", "install", "libxt")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: woff2")
exec.Command("latte", "install", "woff2")
	fmt.Println("Instalando dependencia: wpebackend-fdo")
exec.Command("latte", "install", "wpebackend-fdo")
	fmt.Println("Instalando dependencia: zlib")
exec.Command("latte", "install", "zlib")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: libdrm")
exec.Command("latte", "install", "libdrm")
	fmt.Println("Instalando dependencia: libepoxy")
exec.Command("latte", "install", "libepoxy")
	fmt.Println("Instalando dependencia: libtasn1")
exec.Command("latte", "install", "libtasn1")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: wayland")
exec.Command("latte", "install", "wayland")
}
