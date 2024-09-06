package main

// Autotrace - Convert bitmap to vector graphics
// Homepage: https://autotrace.sourceforge.io

import (
	"fmt"
	
	"os/exec"
)

func installAutotrace() {
	// Método 1: Descargar y extraer .tar.gz
	autotrace_tar_url := "https://github.com/autotrace/autotrace/archive/refs/tags/0.31.10.tar.gz"
	autotrace_cmd_tar := exec.Command("curl", "-L", autotrace_tar_url, "-o", "package.tar.gz")
	err := autotrace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autotrace_zip_url := "https://github.com/autotrace/autotrace/archive/refs/tags/0.31.10.zip"
	autotrace_cmd_zip := exec.Command("curl", "-L", autotrace_zip_url, "-o", "package.zip")
	err = autotrace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autotrace_bin_url := "https://github.com/autotrace/autotrace/archive/refs/tags/0.31.10.bin"
	autotrace_cmd_bin := exec.Command("curl", "-L", autotrace_bin_url, "-o", "binary.bin")
	err = autotrace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autotrace_src_url := "https://github.com/autotrace/autotrace/archive/refs/tags/0.31.10.src.tar.gz"
	autotrace_cmd_src := exec.Command("curl", "-L", autotrace_src_url, "-o", "source.tar.gz")
	err = autotrace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autotrace_cmd_direct := exec.Command("./binary")
	err = autotrace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: pstoedit")
exec.Command("latte", "install", "pstoedit")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: liblqr")
exec.Command("latte", "install", "liblqr")
	fmt.Println("Instalando dependencia: libomp")
exec.Command("latte", "install", "libomp")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: little-cms2")
exec.Command("latte", "install", "little-cms2")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
}
