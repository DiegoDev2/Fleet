package main

// Fontforge - Command-line outline and bitmap font editor/converter
// Homepage: https://fontforge.github.io

import (
	"fmt"
	
	"os/exec"
)

func installFontforge() {
	// Método 1: Descargar y extraer .tar.gz
	fontforge_tar_url := "https://github.com/fontforge/fontforge/releases/download/20230101/fontforge-20230101.tar.xz"
	fontforge_cmd_tar := exec.Command("curl", "-L", fontforge_tar_url, "-o", "package.tar.gz")
	err := fontforge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fontforge_zip_url := "https://github.com/fontforge/fontforge/releases/download/20230101/fontforge-20230101.tar.xz"
	fontforge_cmd_zip := exec.Command("curl", "-L", fontforge_zip_url, "-o", "package.zip")
	err = fontforge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fontforge_bin_url := "https://github.com/fontforge/fontforge/releases/download/20230101/fontforge-20230101.tar.xz"
	fontforge_cmd_bin := exec.Command("curl", "-L", fontforge_bin_url, "-o", "binary.bin")
	err = fontforge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fontforge_src_url := "https://github.com/fontforge/fontforge/releases/download/20230101/fontforge-20230101.tar.xz"
	fontforge_cmd_src := exec.Command("curl", "-L", fontforge_src_url, "-o", "source.tar.gz")
	err = fontforge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fontforge_cmd_direct := exec.Command("./binary")
	err = fontforge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libspiro")
	exec.Command("latte", "install", "libspiro").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libuninameslist")
	exec.Command("latte", "install", "libuninameslist").Run()
	fmt.Println("Instalando dependencia: pango")
	exec.Command("latte", "install", "pango").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
	fmt.Println("Instalando dependencia: woff2")
	exec.Command("latte", "install", "woff2").Run()
	fmt.Println("Instalando dependencia: brotli")
	exec.Command("latte", "install", "brotli").Run()
}
