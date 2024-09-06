package main

// Tgui - GUI library for use with sfml
// Homepage: https://tgui.eu

import (
	"fmt"
	
	"os/exec"
)

func installTgui() {
	// Método 1: Descargar y extraer .tar.gz
	tgui_tar_url := "https://github.com/texus/TGUI/archive/refs/tags/v1.5.0.tar.gz"
	tgui_cmd_tar := exec.Command("curl", "-L", tgui_tar_url, "-o", "package.tar.gz")
	err := tgui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tgui_zip_url := "https://github.com/texus/TGUI/archive/refs/tags/v1.5.0.zip"
	tgui_cmd_zip := exec.Command("curl", "-L", tgui_zip_url, "-o", "package.zip")
	err = tgui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tgui_bin_url := "https://github.com/texus/TGUI/archive/refs/tags/v1.5.0.bin"
	tgui_cmd_bin := exec.Command("curl", "-L", tgui_bin_url, "-o", "binary.bin")
	err = tgui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tgui_src_url := "https://github.com/texus/TGUI/archive/refs/tags/v1.5.0.src.tar.gz"
	tgui_cmd_src := exec.Command("curl", "-L", tgui_src_url, "-o", "source.tar.gz")
	err = tgui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tgui_cmd_direct := exec.Command("./binary")
	err = tgui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: sfml")
	exec.Command("latte", "install", "sfml").Run()
}
