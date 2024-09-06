package main

// Scrcpy - Display and control your Android device
// Homepage: https://github.com/Genymobile/scrcpy

import (
	"fmt"
	
	"os/exec"
)

func installScrcpy() {
	// Método 1: Descargar y extraer .tar.gz
	scrcpy_tar_url := "https://github.com/Genymobile/scrcpy/archive/refs/tags/v2.6.1.tar.gz"
	scrcpy_cmd_tar := exec.Command("curl", "-L", scrcpy_tar_url, "-o", "package.tar.gz")
	err := scrcpy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scrcpy_zip_url := "https://github.com/Genymobile/scrcpy/archive/refs/tags/v2.6.1.zip"
	scrcpy_cmd_zip := exec.Command("curl", "-L", scrcpy_zip_url, "-o", "package.zip")
	err = scrcpy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scrcpy_bin_url := "https://github.com/Genymobile/scrcpy/archive/refs/tags/v2.6.1.bin"
	scrcpy_cmd_bin := exec.Command("curl", "-L", scrcpy_bin_url, "-o", "binary.bin")
	err = scrcpy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scrcpy_src_url := "https://github.com/Genymobile/scrcpy/archive/refs/tags/v2.6.1.src.tar.gz"
	scrcpy_cmd_src := exec.Command("curl", "-L", scrcpy_src_url, "-o", "source.tar.gz")
	err = scrcpy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scrcpy_cmd_direct := exec.Command("./binary")
	err = scrcpy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
