package main

// Siril - Astronomical image processing tool
// Homepage: https://www.siril.org

import (
	"fmt"
	
	"os/exec"
)

func installSiril() {
	// Método 1: Descargar y extraer .tar.gz
	siril_tar_url := "https://free-astro.org/download/siril-1.2.3.tar.bz2"
	siril_cmd_tar := exec.Command("curl", "-L", siril_tar_url, "-o", "package.tar.gz")
	err := siril_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	siril_zip_url := "https://free-astro.org/download/siril-1.2.3.tar.bz2"
	siril_cmd_zip := exec.Command("curl", "-L", siril_zip_url, "-o", "package.zip")
	err = siril_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	siril_bin_url := "https://free-astro.org/download/siril-1.2.3.tar.bz2"
	siril_cmd_bin := exec.Command("curl", "-L", siril_bin_url, "-o", "binary.bin")
	err = siril_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	siril_src_url := "https://free-astro.org/download/siril-1.2.3.tar.bz2"
	siril_cmd_src := exec.Command("curl", "-L", siril_src_url, "-o", "source.tar.gz")
	err = siril_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	siril_cmd_direct := exec.Command("./binary")
	err = siril_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
exec.Command("latte", "install", "adwaita-icon-theme")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: cfitsio")
exec.Command("latte", "install", "cfitsio")
	fmt.Println("Instalando dependencia: exiv2")
exec.Command("latte", "install", "exiv2")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: ffms2")
exec.Command("latte", "install", "ffms2")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gnuplot")
exec.Command("latte", "install", "gnuplot")
	fmt.Println("Instalando dependencia: gsl")
exec.Command("latte", "install", "gsl")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: json-glib")
exec.Command("latte", "install", "json-glib")
	fmt.Println("Instalando dependencia: libconfig")
exec.Command("latte", "install", "libconfig")
	fmt.Println("Instalando dependencia: libheif")
exec.Command("latte", "install", "libheif")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libraw")
exec.Command("latte", "install", "libraw")
	fmt.Println("Instalando dependencia: librsvg")
exec.Command("latte", "install", "librsvg")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: netpbm")
exec.Command("latte", "install", "netpbm")
	fmt.Println("Instalando dependencia: opencv")
exec.Command("latte", "install", "opencv")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: wcslib")
exec.Command("latte", "install", "wcslib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gtk-mac-integration")
exec.Command("latte", "install", "gtk-mac-integration")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
}
