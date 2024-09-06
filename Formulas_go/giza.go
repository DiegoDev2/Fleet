package main

// Giza - Scientific plotting library for C/Fortran built on cairo
// Homepage: https://danieljprice.github.io/giza/

import (
	"fmt"
	
	"os/exec"
)

func installGiza() {
	// Método 1: Descargar y extraer .tar.gz
	giza_tar_url := "https://github.com/danieljprice/giza/archive/refs/tags/v1.4.1.tar.gz"
	giza_cmd_tar := exec.Command("curl", "-L", giza_tar_url, "-o", "package.tar.gz")
	err := giza_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	giza_zip_url := "https://github.com/danieljprice/giza/archive/refs/tags/v1.4.1.zip"
	giza_cmd_zip := exec.Command("curl", "-L", giza_zip_url, "-o", "package.zip")
	err = giza_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	giza_bin_url := "https://github.com/danieljprice/giza/archive/refs/tags/v1.4.1.bin"
	giza_cmd_bin := exec.Command("curl", "-L", giza_bin_url, "-o", "binary.bin")
	err = giza_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	giza_src_url := "https://github.com/danieljprice/giza/archive/refs/tags/v1.4.1.src.tar.gz"
	giza_cmd_src := exec.Command("curl", "-L", giza_src_url, "-o", "source.tar.gz")
	err = giza_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	giza_cmd_direct := exec.Command("./binary")
	err = giza_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: gcc")
exec.Command("latte", "install", "gcc")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
}
