package main

// Gwyddion - Scanning Probe Microscopy visualization and analysis tool
// Homepage: http://gwyddion.net/

import (
	"fmt"
	
	"os/exec"
)

func installGwyddion() {
	// Método 1: Descargar y extraer .tar.gz
	gwyddion_tar_url := "https://downloads.sourceforge.net/project/gwyddion/gwyddion/2.66/gwyddion-2.66.tar.xz"
	gwyddion_cmd_tar := exec.Command("curl", "-L", gwyddion_tar_url, "-o", "package.tar.gz")
	err := gwyddion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gwyddion_zip_url := "https://downloads.sourceforge.net/project/gwyddion/gwyddion/2.66/gwyddion-2.66.tar.xz"
	gwyddion_cmd_zip := exec.Command("curl", "-L", gwyddion_zip_url, "-o", "package.zip")
	err = gwyddion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gwyddion_bin_url := "https://downloads.sourceforge.net/project/gwyddion/gwyddion/2.66/gwyddion-2.66.tar.xz"
	gwyddion_cmd_bin := exec.Command("curl", "-L", gwyddion_bin_url, "-o", "binary.bin")
	err = gwyddion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gwyddion_src_url := "https://downloads.sourceforge.net/project/gwyddion/gwyddion/2.66/gwyddion-2.66.tar.xz"
	gwyddion_cmd_src := exec.Command("curl", "-L", gwyddion_src_url, "-o", "source.tar.gz")
	err = gwyddion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gwyddion_cmd_direct := exec.Command("./binary")
	err = gwyddion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+")
exec.Command("latte", "install", "gtk+")
	fmt.Println("Instalando dependencia: gtkglext")
exec.Command("latte", "install", "gtkglext")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: minizip")
exec.Command("latte", "install", "minizip")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gtk-doc")
exec.Command("latte", "install", "gtk-doc")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: gtk-mac-integration")
exec.Command("latte", "install", "gtk-mac-integration")
	fmt.Println("Instalando dependencia: at-spi2-core")
exec.Command("latte", "install", "at-spi2-core")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
