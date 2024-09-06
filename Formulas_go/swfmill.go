package main

// Swfmill - Processor of xml2swf and swf2xml
// Homepage: https://www.swfmill.org/

import (
	"fmt"
	
	"os/exec"
)

func installSwfmill() {
	// Método 1: Descargar y extraer .tar.gz
	swfmill_tar_url := "https://www.swfmill.org/releases/swfmill-0.3.6.tar.gz"
	swfmill_cmd_tar := exec.Command("curl", "-L", swfmill_tar_url, "-o", "package.tar.gz")
	err := swfmill_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swfmill_zip_url := "https://www.swfmill.org/releases/swfmill-0.3.6.zip"
	swfmill_cmd_zip := exec.Command("curl", "-L", swfmill_zip_url, "-o", "package.zip")
	err = swfmill_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swfmill_bin_url := "https://www.swfmill.org/releases/swfmill-0.3.6.bin"
	swfmill_cmd_bin := exec.Command("curl", "-L", swfmill_bin_url, "-o", "binary.bin")
	err = swfmill_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swfmill_src_url := "https://www.swfmill.org/releases/swfmill-0.3.6.src.tar.gz"
	swfmill_cmd_src := exec.Command("curl", "-L", swfmill_src_url, "-o", "source.tar.gz")
	err = swfmill_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swfmill_cmd_direct := exec.Command("./binary")
	err = swfmill_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
