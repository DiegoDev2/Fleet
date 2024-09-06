package main

// Terminalimageviewer - Display images in a terminal using block graphic characters
// Homepage: https://github.com/stefanhaustein/TerminalImageViewer

import (
	"fmt"
	
	"os/exec"
)

func installTerminalimageviewer() {
	// Método 1: Descargar y extraer .tar.gz
	terminalimageviewer_tar_url := "https://github.com/stefanhaustein/TerminalImageViewer/archive/refs/tags/v1.2.1.tar.gz"
	terminalimageviewer_cmd_tar := exec.Command("curl", "-L", terminalimageviewer_tar_url, "-o", "package.tar.gz")
	err := terminalimageviewer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terminalimageviewer_zip_url := "https://github.com/stefanhaustein/TerminalImageViewer/archive/refs/tags/v1.2.1.zip"
	terminalimageviewer_cmd_zip := exec.Command("curl", "-L", terminalimageviewer_zip_url, "-o", "package.zip")
	err = terminalimageviewer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terminalimageviewer_bin_url := "https://github.com/stefanhaustein/TerminalImageViewer/archive/refs/tags/v1.2.1.bin"
	terminalimageviewer_cmd_bin := exec.Command("curl", "-L", terminalimageviewer_bin_url, "-o", "binary.bin")
	err = terminalimageviewer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terminalimageviewer_src_url := "https://github.com/stefanhaustein/TerminalImageViewer/archive/refs/tags/v1.2.1.src.tar.gz"
	terminalimageviewer_cmd_src := exec.Command("curl", "-L", terminalimageviewer_src_url, "-o", "source.tar.gz")
	err = terminalimageviewer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terminalimageviewer_cmd_direct := exec.Command("./binary")
	err = terminalimageviewer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
}
