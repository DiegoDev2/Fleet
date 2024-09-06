package main

// Poppler - PDF rendering library (based on the xpdf-3.0 code base)
// Homepage: https://poppler.freedesktop.org/

import (
	"fmt"
	
	"os/exec"
)

func installPoppler() {
	// Método 1: Descargar y extraer .tar.gz
	poppler_tar_url := "https://poppler.freedesktop.org/poppler-24.04.0.tar.xz"
	poppler_cmd_tar := exec.Command("curl", "-L", poppler_tar_url, "-o", "package.tar.gz")
	err := poppler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	poppler_zip_url := "https://poppler.freedesktop.org/poppler-24.04.0.tar.xz"
	poppler_cmd_zip := exec.Command("curl", "-L", poppler_zip_url, "-o", "package.zip")
	err = poppler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	poppler_bin_url := "https://poppler.freedesktop.org/poppler-24.04.0.tar.xz"
	poppler_cmd_bin := exec.Command("curl", "-L", poppler_bin_url, "-o", "binary.bin")
	err = poppler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	poppler_src_url := "https://poppler.freedesktop.org/poppler-24.04.0.tar.xz"
	poppler_cmd_src := exec.Command("curl", "-L", poppler_src_url, "-o", "source.tar.gz")
	err = poppler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	poppler_cmd_direct := exec.Command("./binary")
	err = poppler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gpgme")
	exec.Command("latte", "install", "gpgme").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: little-cms2")
	exec.Command("latte", "install", "little-cms2").Run()
	fmt.Println("Instalando dependencia: nspr")
	exec.Command("latte", "install", "nspr").Run()
	fmt.Println("Instalando dependencia: nss")
	exec.Command("latte", "install", "nss").Run()
	fmt.Println("Instalando dependencia: openjpeg")
	exec.Command("latte", "install", "openjpeg").Run()
	fmt.Println("Instalando dependencia: libassuan")
	exec.Command("latte", "install", "libassuan").Run()
}
