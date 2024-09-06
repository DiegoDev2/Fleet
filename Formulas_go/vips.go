package main

// Vips - Image processing library
// Homepage: https://github.com/libvips/libvips

import (
	"fmt"
	
	"os/exec"
)

func installVips() {
	// Método 1: Descargar y extraer .tar.gz
	vips_tar_url := "https://github.com/libvips/libvips/releases/download/v8.15.3/vips-8.15.3.tar.xz"
	vips_cmd_tar := exec.Command("curl", "-L", vips_tar_url, "-o", "package.tar.gz")
	err := vips_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vips_zip_url := "https://github.com/libvips/libvips/releases/download/v8.15.3/vips-8.15.3.tar.xz"
	vips_cmd_zip := exec.Command("curl", "-L", vips_zip_url, "-o", "package.zip")
	err = vips_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vips_bin_url := "https://github.com/libvips/libvips/releases/download/v8.15.3/vips-8.15.3.tar.xz"
	vips_cmd_bin := exec.Command("curl", "-L", vips_bin_url, "-o", "binary.bin")
	err = vips_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vips_src_url := "https://github.com/libvips/libvips/releases/download/v8.15.3/vips-8.15.3.tar.xz"
	vips_cmd_src := exec.Command("curl", "-L", vips_src_url, "-o", "source.tar.gz")
	err = vips_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vips_cmd_direct := exec.Command("./binary")
	err = vips_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: cfitsio")
exec.Command("latte", "install", "cfitsio")
	fmt.Println("Instalando dependencia: cgif")
exec.Command("latte", "install", "cgif")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: highway")
exec.Command("latte", "install", "highway")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: jpeg-xl")
exec.Command("latte", "install", "jpeg-xl")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libexif")
exec.Command("latte", "install", "libexif")
	fmt.Println("Instalando dependencia: libheif")
exec.Command("latte", "install", "libheif")
	fmt.Println("Instalando dependencia: libimagequant")
exec.Command("latte", "install", "libimagequant")
	fmt.Println("Instalando dependencia: libmatio")
exec.Command("latte", "install", "libmatio")
	fmt.Println("Instalando dependencia: librsvg")
exec.Command("latte", "install", "librsvg")
	fmt.Println("Instalando dependencia: libspng")
exec.Command("latte", "install", "libspng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: mozjpeg")
exec.Command("latte", "install", "mozjpeg")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: openslide")
exec.Command("latte", "install", "openslide")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
}
