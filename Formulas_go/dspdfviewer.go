package main

// Dspdfviewer - Dual-Screen PDF Viewer for latex-beamer
// Homepage: https://dspdfviewer.danny-edel.de/

import (
	"fmt"
	
	"os/exec"
)

func installDspdfviewer() {
	// Método 1: Descargar y extraer .tar.gz
	dspdfviewer_tar_url := "https://github.com/dannyedel/dspdfviewer/archive/refs/tags/v1.15.1.tar.gz"
	dspdfviewer_cmd_tar := exec.Command("curl", "-L", dspdfviewer_tar_url, "-o", "package.tar.gz")
	err := dspdfviewer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dspdfviewer_zip_url := "https://github.com/dannyedel/dspdfviewer/archive/refs/tags/v1.15.1.zip"
	dspdfviewer_cmd_zip := exec.Command("curl", "-L", dspdfviewer_zip_url, "-o", "package.zip")
	err = dspdfviewer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dspdfviewer_bin_url := "https://github.com/dannyedel/dspdfviewer/archive/refs/tags/v1.15.1.bin"
	dspdfviewer_cmd_bin := exec.Command("curl", "-L", dspdfviewer_bin_url, "-o", "binary.bin")
	err = dspdfviewer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dspdfviewer_src_url := "https://github.com/dannyedel/dspdfviewer/archive/refs/tags/v1.15.1.src.tar.gz"
	dspdfviewer_cmd_src := exec.Command("curl", "-L", dspdfviewer_src_url, "-o", "source.tar.gz")
	err = dspdfviewer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dspdfviewer_cmd_direct := exec.Command("./binary")
	err = dspdfviewer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: openjpeg")
exec.Command("latte", "install", "openjpeg")
	fmt.Println("Instalando dependencia: poppler-qt5")
exec.Command("latte", "install", "poppler-qt5")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
