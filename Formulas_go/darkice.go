package main

// Darkice - Live audio streamer
// Homepage: http://www.darkice.org/

import (
	"fmt"
	
	"os/exec"
)

func installDarkice() {
	// Método 1: Descargar y extraer .tar.gz
	darkice_tar_url := "https://github.com/rafael2k/darkice/releases/download/v1.5/darkice-1.5.tar.gz"
	darkice_cmd_tar := exec.Command("curl", "-L", darkice_tar_url, "-o", "package.tar.gz")
	err := darkice_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	darkice_zip_url := "https://github.com/rafael2k/darkice/releases/download/v1.5/darkice-1.5.zip"
	darkice_cmd_zip := exec.Command("curl", "-L", darkice_zip_url, "-o", "package.zip")
	err = darkice_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	darkice_bin_url := "https://github.com/rafael2k/darkice/releases/download/v1.5/darkice-1.5.bin"
	darkice_cmd_bin := exec.Command("curl", "-L", darkice_bin_url, "-o", "binary.bin")
	err = darkice_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	darkice_src_url := "https://github.com/rafael2k/darkice/releases/download/v1.5/darkice-1.5.src.tar.gz"
	darkice_cmd_src := exec.Command("curl", "-L", darkice_src_url, "-o", "source.tar.gz")
	err = darkice_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	darkice_cmd_direct := exec.Command("./binary")
	err = darkice_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: faac")
exec.Command("latte", "install", "faac")
	fmt.Println("Instalando dependencia: jack")
exec.Command("latte", "install", "jack")
	fmt.Println("Instalando dependencia: lame")
exec.Command("latte", "install", "lame")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: two-lame")
exec.Command("latte", "install", "two-lame")
	fmt.Println("Instalando dependencia: alsa-lib")
exec.Command("latte", "install", "alsa-lib")
}
