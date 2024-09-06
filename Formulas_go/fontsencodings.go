package main

// FontsEncodings - Font encoding tables for libfontenc
// Homepage: https://gitlab.freedesktop.org/xorg/font/encodings

import (
	"fmt"
	
	"os/exec"
)

func installFontsEncodings() {
	// Método 1: Descargar y extraer .tar.gz
	fontsencodings_tar_url := "https://www.x.org/archive/individual/font/encodings-1.1.0.tar.xz"
	fontsencodings_cmd_tar := exec.Command("curl", "-L", fontsencodings_tar_url, "-o", "package.tar.gz")
	err := fontsencodings_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fontsencodings_zip_url := "https://www.x.org/archive/individual/font/encodings-1.1.0.tar.xz"
	fontsencodings_cmd_zip := exec.Command("curl", "-L", fontsencodings_zip_url, "-o", "package.zip")
	err = fontsencodings_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fontsencodings_bin_url := "https://www.x.org/archive/individual/font/encodings-1.1.0.tar.xz"
	fontsencodings_cmd_bin := exec.Command("curl", "-L", fontsencodings_bin_url, "-o", "binary.bin")
	err = fontsencodings_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fontsencodings_src_url := "https://www.x.org/archive/individual/font/encodings-1.1.0.tar.xz"
	fontsencodings_cmd_src := exec.Command("curl", "-L", fontsencodings_src_url, "-o", "source.tar.gz")
	err = fontsencodings_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fontsencodings_cmd_direct := exec.Command("./binary")
	err = fontsencodings_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: font-util")
exec.Command("latte", "install", "font-util")
	fmt.Println("Instalando dependencia: mkfontscale")
exec.Command("latte", "install", "mkfontscale")
	fmt.Println("Instalando dependencia: util-macros")
exec.Command("latte", "install", "util-macros")
}
