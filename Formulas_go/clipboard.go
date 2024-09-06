package main

// Clipboard - Cut, copy, and paste anything, anywhere, all from the terminal
// Homepage: https://getclipboard.app

import (
	"fmt"
	
	"os/exec"
)

func installClipboard() {
	// Método 1: Descargar y extraer .tar.gz
	clipboard_tar_url := "https://github.com/Slackadays/Clipboard/archive/refs/tags/0.9.0.1.tar.gz"
	clipboard_cmd_tar := exec.Command("curl", "-L", clipboard_tar_url, "-o", "package.tar.gz")
	err := clipboard_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clipboard_zip_url := "https://github.com/Slackadays/Clipboard/archive/refs/tags/0.9.0.1.zip"
	clipboard_cmd_zip := exec.Command("curl", "-L", clipboard_zip_url, "-o", "package.zip")
	err = clipboard_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clipboard_bin_url := "https://github.com/Slackadays/Clipboard/archive/refs/tags/0.9.0.1.bin"
	clipboard_cmd_bin := exec.Command("curl", "-L", clipboard_bin_url, "-o", "binary.bin")
	err = clipboard_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clipboard_src_url := "https://github.com/Slackadays/Clipboard/archive/refs/tags/0.9.0.1.src.tar.gz"
	clipboard_cmd_src := exec.Command("curl", "-L", clipboard_src_url, "-o", "source.tar.gz")
	err = clipboard_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clipboard_cmd_direct := exec.Command("./binary")
	err = clipboard_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: wayland-protocols")
exec.Command("latte", "install", "wayland-protocols")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: wayland")
exec.Command("latte", "install", "wayland")
}
