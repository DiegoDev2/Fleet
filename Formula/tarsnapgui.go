package main

// TarsnapGui - Cross-platform GUI for the Tarsnap command-line client
// Homepage: https://github.com/Tarsnap/tarsnap-gui/wiki

import (
	"fmt"
	
	"os/exec"
)

func installTarsnapGui() {
	// Método 1: Descargar y extraer .tar.gz
	tarsnapgui_tar_url := "https://github.com/Tarsnap/tarsnap-gui/archive/refs/tags/v1.0.2.tar.gz"
	tarsnapgui_cmd_tar := exec.Command("curl", "-L", tarsnapgui_tar_url, "-o", "package.tar.gz")
	err := tarsnapgui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tarsnapgui_zip_url := "https://github.com/Tarsnap/tarsnap-gui/archive/refs/tags/v1.0.2.zip"
	tarsnapgui_cmd_zip := exec.Command("curl", "-L", tarsnapgui_zip_url, "-o", "package.zip")
	err = tarsnapgui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tarsnapgui_bin_url := "https://github.com/Tarsnap/tarsnap-gui/archive/refs/tags/v1.0.2.bin"
	tarsnapgui_cmd_bin := exec.Command("curl", "-L", tarsnapgui_bin_url, "-o", "binary.bin")
	err = tarsnapgui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tarsnapgui_src_url := "https://github.com/Tarsnap/tarsnap-gui/archive/refs/tags/v1.0.2.src.tar.gz"
	tarsnapgui_cmd_src := exec.Command("curl", "-L", tarsnapgui_src_url, "-o", "source.tar.gz")
	err = tarsnapgui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tarsnapgui_cmd_direct := exec.Command("./binary")
	err = tarsnapgui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: qt@5")
	exec.Command("latte", "install", "qt@5").Run()
	fmt.Println("Instalando dependencia: tarsnap")
	exec.Command("latte", "install", "tarsnap").Run()
}
