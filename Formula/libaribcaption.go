package main

// Libaribcaption - Portable ARIB STD-B24 Caption Decoder/Renderer
// Homepage: https://github.com/xqq/libaribcaption

import (
	"fmt"
	
	"os/exec"
)

func installLibaribcaption() {
	// Método 1: Descargar y extraer .tar.gz
	libaribcaption_tar_url := "https://github.com/xqq/libaribcaption/archive/refs/tags/v1.1.1.tar.gz"
	libaribcaption_cmd_tar := exec.Command("curl", "-L", libaribcaption_tar_url, "-o", "package.tar.gz")
	err := libaribcaption_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libaribcaption_zip_url := "https://github.com/xqq/libaribcaption/archive/refs/tags/v1.1.1.zip"
	libaribcaption_cmd_zip := exec.Command("curl", "-L", libaribcaption_zip_url, "-o", "package.zip")
	err = libaribcaption_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libaribcaption_bin_url := "https://github.com/xqq/libaribcaption/archive/refs/tags/v1.1.1.bin"
	libaribcaption_cmd_bin := exec.Command("curl", "-L", libaribcaption_bin_url, "-o", "binary.bin")
	err = libaribcaption_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libaribcaption_src_url := "https://github.com/xqq/libaribcaption/archive/refs/tags/v1.1.1.src.tar.gz"
	libaribcaption_cmd_src := exec.Command("curl", "-L", libaribcaption_src_url, "-o", "source.tar.gz")
	err = libaribcaption_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libaribcaption_cmd_direct := exec.Command("./binary")
	err = libaribcaption_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
}
