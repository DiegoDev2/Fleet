package main

// Pqiv - Powerful image viewer with minimal UI
// Homepage: https://github.com/phillipberndt/pqiv

import (
	"fmt"
	
	"os/exec"
)

func installPqiv() {
	// Método 1: Descargar y extraer .tar.gz
	pqiv_tar_url := "https://github.com/phillipberndt/pqiv/archive/refs/tags/2.13.1.tar.gz"
	pqiv_cmd_tar := exec.Command("curl", "-L", pqiv_tar_url, "-o", "package.tar.gz")
	err := pqiv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pqiv_zip_url := "https://github.com/phillipberndt/pqiv/archive/refs/tags/2.13.1.zip"
	pqiv_cmd_zip := exec.Command("curl", "-L", pqiv_zip_url, "-o", "package.zip")
	err = pqiv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pqiv_bin_url := "https://github.com/phillipberndt/pqiv/archive/refs/tags/2.13.1.bin"
	pqiv_cmd_bin := exec.Command("curl", "-L", pqiv_bin_url, "-o", "binary.bin")
	err = pqiv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pqiv_src_url := "https://github.com/phillipberndt/pqiv/archive/refs/tags/2.13.1.src.tar.gz"
	pqiv_cmd_src := exec.Command("curl", "-L", pqiv_src_url, "-o", "source.tar.gz")
	err = pqiv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pqiv_cmd_direct := exec.Command("./binary")
	err = pqiv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
	fmt.Println("Instalando dependencia: libspectre")
exec.Command("latte", "install", "libspectre")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
}
