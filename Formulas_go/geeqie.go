package main

// Geeqie - Lightweight Gtk+ based image viewer
// Homepage: https://www.geeqie.org/

import (
	"fmt"
	
	"os/exec"
)

func installGeeqie() {
	// Método 1: Descargar y extraer .tar.gz
	geeqie_tar_url := "https://github.com/BestImageViewer/geeqie/releases/download/v2.4/geeqie-2.4.tar.xz"
	geeqie_cmd_tar := exec.Command("curl", "-L", geeqie_tar_url, "-o", "package.tar.gz")
	err := geeqie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geeqie_zip_url := "https://github.com/BestImageViewer/geeqie/releases/download/v2.4/geeqie-2.4.tar.xz"
	geeqie_cmd_zip := exec.Command("curl", "-L", geeqie_zip_url, "-o", "package.zip")
	err = geeqie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geeqie_bin_url := "https://github.com/BestImageViewer/geeqie/releases/download/v2.4/geeqie-2.4.tar.xz"
	geeqie_cmd_bin := exec.Command("curl", "-L", geeqie_bin_url, "-o", "binary.bin")
	err = geeqie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geeqie_src_url := "https://github.com/BestImageViewer/geeqie/releases/download/v2.4/geeqie-2.4.tar.xz"
	geeqie_cmd_src := exec.Command("curl", "-L", geeqie_src_url, "-o", "source.tar.gz")
	err = geeqie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geeqie_cmd_direct := exec.Command("./binary")
	err = geeqie_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: yelp-tools")
exec.Command("latte", "install", "yelp-tools")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: djvulibre")
exec.Command("latte", "install", "djvulibre")
	fmt.Println("Instalando dependencia: evince")
exec.Command("latte", "install", "evince")
	fmt.Println("Instalando dependencia: exiv2")
exec.Command("latte", "install", "exiv2")
	fmt.Println("Instalando dependencia: ffmpegthumbnailer")
exec.Command("latte", "install", "ffmpegthumbnailer")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gspell")
exec.Command("latte", "install", "gspell")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: jpeg-xl")
exec.Command("latte", "install", "jpeg-xl")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libheif")
exec.Command("latte", "install", "libheif")
	fmt.Println("Instalando dependencia: libraw")
exec.Command("latte", "install", "libraw")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: webp-pixbuf-loader")
exec.Command("latte", "install", "webp-pixbuf-loader")
	fmt.Println("Instalando dependencia: enchant")
exec.Command("latte", "install", "enchant")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
