package main

// Ttfautohint - Auto-hinter for TrueType fonts
// Homepage: https://www.freetype.org/ttfautohint/

import (
	"fmt"
	
	"os/exec"
)

func installTtfautohint() {
	// Método 1: Descargar y extraer .tar.gz
	ttfautohint_tar_url := "https://downloads.sourceforge.net/project/freetype/ttfautohint/1.8.4/ttfautohint-1.8.4.tar.gz"
	ttfautohint_cmd_tar := exec.Command("curl", "-L", ttfautohint_tar_url, "-o", "package.tar.gz")
	err := ttfautohint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ttfautohint_zip_url := "https://downloads.sourceforge.net/project/freetype/ttfautohint/1.8.4/ttfautohint-1.8.4.zip"
	ttfautohint_cmd_zip := exec.Command("curl", "-L", ttfautohint_zip_url, "-o", "package.zip")
	err = ttfautohint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ttfautohint_bin_url := "https://downloads.sourceforge.net/project/freetype/ttfautohint/1.8.4/ttfautohint-1.8.4.bin"
	ttfautohint_cmd_bin := exec.Command("curl", "-L", ttfautohint_bin_url, "-o", "binary.bin")
	err = ttfautohint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ttfautohint_src_url := "https://downloads.sourceforge.net/project/freetype/ttfautohint/1.8.4/ttfautohint-1.8.4.src.tar.gz"
	ttfautohint_cmd_src := exec.Command("curl", "-L", ttfautohint_src_url, "-o", "source.tar.gz")
	err = ttfautohint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ttfautohint_cmd_direct := exec.Command("./binary")
	err = ttfautohint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
