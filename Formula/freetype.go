package main

// Freetype - Software library to render fonts
// Homepage: https://www.freetype.org/

import (
	"fmt"
	
	"os/exec"
)

func installFreetype() {
	// Método 1: Descargar y extraer .tar.gz
	freetype_tar_url := "https://downloads.sourceforge.net/project/freetype/freetype2/2.13.3/freetype-2.13.3.tar.xz"
	freetype_cmd_tar := exec.Command("curl", "-L", freetype_tar_url, "-o", "package.tar.gz")
	err := freetype_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freetype_zip_url := "https://downloads.sourceforge.net/project/freetype/freetype2/2.13.3/freetype-2.13.3.tar.xz"
	freetype_cmd_zip := exec.Command("curl", "-L", freetype_zip_url, "-o", "package.zip")
	err = freetype_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freetype_bin_url := "https://downloads.sourceforge.net/project/freetype/freetype2/2.13.3/freetype-2.13.3.tar.xz"
	freetype_cmd_bin := exec.Command("curl", "-L", freetype_bin_url, "-o", "binary.bin")
	err = freetype_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freetype_src_url := "https://downloads.sourceforge.net/project/freetype/freetype2/2.13.3/freetype-2.13.3.tar.xz"
	freetype_cmd_src := exec.Command("curl", "-L", freetype_src_url, "-o", "source.tar.gz")
	err = freetype_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freetype_cmd_direct := exec.Command("./binary")
	err = freetype_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
