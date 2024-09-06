package main

// Xwin - Microsoft CRT and Windows SDK headers and libraries loader
// Homepage: https://github.com/Jake-Shadle/xwin

import (
	"fmt"
	
	"os/exec"
)

func installXwin() {
	// Método 1: Descargar y extraer .tar.gz
	xwin_tar_url := "https://github.com/Jake-Shadle/xwin/archive/refs/tags/0.6.5.tar.gz"
	xwin_cmd_tar := exec.Command("curl", "-L", xwin_tar_url, "-o", "package.tar.gz")
	err := xwin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xwin_zip_url := "https://github.com/Jake-Shadle/xwin/archive/refs/tags/0.6.5.zip"
	xwin_cmd_zip := exec.Command("curl", "-L", xwin_zip_url, "-o", "package.zip")
	err = xwin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xwin_bin_url := "https://github.com/Jake-Shadle/xwin/archive/refs/tags/0.6.5.bin"
	xwin_cmd_bin := exec.Command("curl", "-L", xwin_bin_url, "-o", "binary.bin")
	err = xwin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xwin_src_url := "https://github.com/Jake-Shadle/xwin/archive/refs/tags/0.6.5.src.tar.gz"
	xwin_cmd_src := exec.Command("curl", "-L", xwin_src_url, "-o", "source.tar.gz")
	err = xwin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xwin_cmd_direct := exec.Command("./binary")
	err = xwin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
