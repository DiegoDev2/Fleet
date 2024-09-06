package main

// Xsw - Slide show presentation tool
// Homepage: https://code.google.com/archive/p/xsw/

import (
	"fmt"
	
	"os/exec"
)

func installXsw() {
	// Método 1: Descargar y extraer .tar.gz
	xsw_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/xsw/xsw-0.3.5.tar.gz"
	xsw_cmd_tar := exec.Command("curl", "-L", xsw_tar_url, "-o", "package.tar.gz")
	err := xsw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xsw_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/xsw/xsw-0.3.5.zip"
	xsw_cmd_zip := exec.Command("curl", "-L", xsw_zip_url, "-o", "package.zip")
	err = xsw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xsw_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/xsw/xsw-0.3.5.bin"
	xsw_cmd_bin := exec.Command("curl", "-L", xsw_bin_url, "-o", "binary.bin")
	err = xsw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xsw_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/xsw/xsw-0.3.5.src.tar.gz"
	xsw_cmd_src := exec.Command("curl", "-L", xsw_src_url, "-o", "source.tar.gz")
	err = xsw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xsw_cmd_direct := exec.Command("./binary")
	err = xsw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
	fmt.Println("Instalando dependencia: sdl_gfx")
	exec.Command("latte", "install", "sdl_gfx").Run()
	fmt.Println("Instalando dependencia: sdl_image")
	exec.Command("latte", "install", "sdl_image").Run()
	fmt.Println("Instalando dependencia: sdl_ttf")
	exec.Command("latte", "install", "sdl_ttf").Run()
}
