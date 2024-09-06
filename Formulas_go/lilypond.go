package main

// Lilypond - Music engraving system
// Homepage: https://lilypond.org

import (
	"fmt"
	
	"os/exec"
)

func installLilypond() {
	// Método 1: Descargar y extraer .tar.gz
	lilypond_tar_url := "https://lilypond.org/download/sources/v2.24/lilypond-2.24.4.tar.gz"
	lilypond_cmd_tar := exec.Command("curl", "-L", lilypond_tar_url, "-o", "package.tar.gz")
	err := lilypond_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lilypond_zip_url := "https://lilypond.org/download/sources/v2.24/lilypond-2.24.4.zip"
	lilypond_cmd_zip := exec.Command("curl", "-L", lilypond_zip_url, "-o", "package.zip")
	err = lilypond_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lilypond_bin_url := "https://lilypond.org/download/sources/v2.24/lilypond-2.24.4.bin"
	lilypond_cmd_bin := exec.Command("curl", "-L", lilypond_bin_url, "-o", "binary.bin")
	err = lilypond_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lilypond_src_url := "https://lilypond.org/download/sources/v2.24/lilypond-2.24.4.src.tar.gz"
	lilypond_cmd_src := exec.Command("curl", "-L", lilypond_src_url, "-o", "source.tar.gz")
	err = lilypond_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lilypond_cmd_direct := exec.Command("./binary")
	err = lilypond_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: fontforge")
exec.Command("latte", "install", "fontforge")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: t1utils")
exec.Command("latte", "install", "t1utils")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: texlive")
exec.Command("latte", "install", "texlive")
	fmt.Println("Instalando dependencia: bdw-gc")
exec.Command("latte", "install", "bdw-gc")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: ghostscript")
exec.Command("latte", "install", "ghostscript")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: guile")
exec.Command("latte", "install", "guile")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
