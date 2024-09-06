package main

// Aribb24 - Library for ARIB STD-B24, decoding JIS 8 bit characters and parsing MPEG-TS
// Homepage: https://code.videolan.org/jeeb/aribb24

import (
	"fmt"
	
	"os/exec"
)

func installAribb24() {
	// Método 1: Descargar y extraer .tar.gz
	aribb24_tar_url := "https://code.videolan.org/jeeb/aribb24/-/archive/v1.0.4/aribb24-v1.0.4.tar.bz2"
	aribb24_cmd_tar := exec.Command("curl", "-L", aribb24_tar_url, "-o", "package.tar.gz")
	err := aribb24_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aribb24_zip_url := "https://code.videolan.org/jeeb/aribb24/-/archive/v1.0.4/aribb24-v1.0.4.tar.bz2"
	aribb24_cmd_zip := exec.Command("curl", "-L", aribb24_zip_url, "-o", "package.zip")
	err = aribb24_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aribb24_bin_url := "https://code.videolan.org/jeeb/aribb24/-/archive/v1.0.4/aribb24-v1.0.4.tar.bz2"
	aribb24_cmd_bin := exec.Command("curl", "-L", aribb24_bin_url, "-o", "binary.bin")
	err = aribb24_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aribb24_src_url := "https://code.videolan.org/jeeb/aribb24/-/archive/v1.0.4/aribb24-v1.0.4.tar.bz2"
	aribb24_cmd_src := exec.Command("curl", "-L", aribb24_src_url, "-o", "source.tar.gz")
	err = aribb24_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aribb24_cmd_direct := exec.Command("./binary")
	err = aribb24_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
