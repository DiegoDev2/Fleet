package main

// Synfig - Command-line renderer
// Homepage: https://synfig.org/

import (
	"fmt"
	
	"os/exec"
)

func installSynfig() {
	// Método 1: Descargar y extraer .tar.gz
	synfig_tar_url := "https://downloads.sourceforge.net/project/synfig/development/1.5.2/source/synfig-1.5.2.tar.gz"
	synfig_cmd_tar := exec.Command("curl", "-L", synfig_tar_url, "-o", "package.tar.gz")
	err := synfig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	synfig_zip_url := "https://downloads.sourceforge.net/project/synfig/development/1.5.2/source/synfig-1.5.2.zip"
	synfig_cmd_zip := exec.Command("curl", "-L", synfig_zip_url, "-o", "package.zip")
	err = synfig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	synfig_bin_url := "https://downloads.sourceforge.net/project/synfig/development/1.5.2/source/synfig-1.5.2.bin"
	synfig_cmd_bin := exec.Command("curl", "-L", synfig_bin_url, "-o", "binary.bin")
	err = synfig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	synfig_src_url := "https://downloads.sourceforge.net/project/synfig/development/1.5.2/source/synfig-1.5.2.src.tar.gz"
	synfig_cmd_src := exec.Command("curl", "-L", synfig_src_url, "-o", "source.tar.gz")
	err = synfig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	synfig_cmd_direct := exec.Command("./binary")
	err = synfig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: etl")
exec.Command("latte", "install", "etl")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: fftw")
exec.Command("latte", "install", "fftw")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: fribidi")
exec.Command("latte", "install", "fribidi")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: glibmm@2.66")
exec.Command("latte", "install", "glibmm@2.66")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: libmng")
exec.Command("latte", "install", "libmng")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsigc++@2")
exec.Command("latte", "install", "libsigc++@2")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libxml++")
exec.Command("latte", "install", "libxml++")
	fmt.Println("Instalando dependencia: mlt")
exec.Command("latte", "install", "mlt")
	fmt.Println("Instalando dependencia: openexr")
exec.Command("latte", "install", "openexr")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: liblqr")
exec.Command("latte", "install", "liblqr")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
