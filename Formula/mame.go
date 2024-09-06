package main

// Mame - Multiple Arcade Machine Emulator
// Homepage: https://mamedev.org/

import (
	"fmt"
	
	"os/exec"
)

func installMame() {
	// Método 1: Descargar y extraer .tar.gz
	mame_tar_url := "https://github.com/mamedev/mame/archive/refs/tags/mame0269.tar.gz"
	mame_cmd_tar := exec.Command("curl", "-L", mame_tar_url, "-o", "package.tar.gz")
	err := mame_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mame_zip_url := "https://github.com/mamedev/mame/archive/refs/tags/mame0269.zip"
	mame_cmd_zip := exec.Command("curl", "-L", mame_zip_url, "-o", "package.zip")
	err = mame_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mame_bin_url := "https://github.com/mamedev/mame/archive/refs/tags/mame0269.bin"
	mame_cmd_bin := exec.Command("curl", "-L", mame_bin_url, "-o", "binary.bin")
	err = mame_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mame_src_url := "https://github.com/mamedev/mame/archive/refs/tags/mame0269.src.tar.gz"
	mame_cmd_src := exec.Command("curl", "-L", mame_src_url, "-o", "source.tar.gz")
	err = mame_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mame_cmd_direct := exec.Command("./binary")
	err = mame_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asio")
	exec.Command("latte", "install", "asio").Run()
	fmt.Println("Instalando dependencia: glm")
	exec.Command("latte", "install", "glm").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rapidjson")
	exec.Command("latte", "install", "rapidjson").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: portaudio")
	exec.Command("latte", "install", "portaudio").Run()
	fmt.Println("Instalando dependencia: portmidi")
	exec.Command("latte", "install", "portmidi").Run()
	fmt.Println("Instalando dependencia: pugixml")
	exec.Command("latte", "install", "pugixml").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: utf8proc")
	exec.Command("latte", "install", "utf8proc").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
	fmt.Println("Instalando dependencia: qt@5")
	exec.Command("latte", "install", "qt@5").Run()
	fmt.Println("Instalando dependencia: sdl2_ttf")
	exec.Command("latte", "install", "sdl2_ttf").Run()
}
