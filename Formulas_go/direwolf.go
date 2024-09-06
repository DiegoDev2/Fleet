package main

// Direwolf - Software \
// Homepage: https://github.com/wb2osz/direwolf

import (
	"fmt"
	
	"os/exec"
)

func installDirewolf() {
	// Método 1: Descargar y extraer .tar.gz
	direwolf_tar_url := "https://github.com/wb2osz/direwolf/archive/refs/tags/1.7.tar.gz"
	direwolf_cmd_tar := exec.Command("curl", "-L", direwolf_tar_url, "-o", "package.tar.gz")
	err := direwolf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	direwolf_zip_url := "https://github.com/wb2osz/direwolf/archive/refs/tags/1.7.zip"
	direwolf_cmd_zip := exec.Command("curl", "-L", direwolf_zip_url, "-o", "package.zip")
	err = direwolf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	direwolf_bin_url := "https://github.com/wb2osz/direwolf/archive/refs/tags/1.7.bin"
	direwolf_cmd_bin := exec.Command("curl", "-L", direwolf_bin_url, "-o", "binary.bin")
	err = direwolf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	direwolf_src_url := "https://github.com/wb2osz/direwolf/archive/refs/tags/1.7.src.tar.gz"
	direwolf_cmd_src := exec.Command("curl", "-L", direwolf_src_url, "-o", "source.tar.gz")
	err = direwolf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	direwolf_cmd_direct := exec.Command("./binary")
	err = direwolf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: gpsd")
exec.Command("latte", "install", "gpsd")
	fmt.Println("Instalando dependencia: hamlib")
exec.Command("latte", "install", "hamlib")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
