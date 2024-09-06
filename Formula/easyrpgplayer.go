package main

// EasyrpgPlayer - RPG Maker 2000/2003 games interpreter
// Homepage: https://easyrpg.org/

import (
	"fmt"
	
	"os/exec"
)

func installEasyrpgPlayer() {
	// Método 1: Descargar y extraer .tar.gz
	easyrpgplayer_tar_url := "https://easyrpg.org/downloads/player/0.8/easyrpg-player-0.8.tar.xz"
	easyrpgplayer_cmd_tar := exec.Command("curl", "-L", easyrpgplayer_tar_url, "-o", "package.tar.gz")
	err := easyrpgplayer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	easyrpgplayer_zip_url := "https://easyrpg.org/downloads/player/0.8/easyrpg-player-0.8.tar.xz"
	easyrpgplayer_cmd_zip := exec.Command("curl", "-L", easyrpgplayer_zip_url, "-o", "package.zip")
	err = easyrpgplayer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	easyrpgplayer_bin_url := "https://easyrpg.org/downloads/player/0.8/easyrpg-player-0.8.tar.xz"
	easyrpgplayer_cmd_bin := exec.Command("curl", "-L", easyrpgplayer_bin_url, "-o", "binary.bin")
	err = easyrpgplayer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	easyrpgplayer_src_url := "https://easyrpg.org/downloads/player/0.8/easyrpg-player-0.8.tar.xz"
	easyrpgplayer_cmd_src := exec.Command("curl", "-L", easyrpgplayer_src_url, "-o", "source.tar.gz")
	err = easyrpgplayer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	easyrpgplayer_cmd_direct := exec.Command("./binary")
	err = easyrpgplayer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: expat")
	exec.Command("latte", "install", "expat").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: harfbuzz")
	exec.Command("latte", "install", "harfbuzz").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: liblcf")
	exec.Command("latte", "install", "liblcf").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: libxmp")
	exec.Command("latte", "install", "libxmp").Run()
	fmt.Println("Instalando dependencia: mpg123")
	exec.Command("latte", "install", "mpg123").Run()
	fmt.Println("Instalando dependencia: pixman")
	exec.Command("latte", "install", "pixman").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: speexdsp")
	exec.Command("latte", "install", "speexdsp").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
}
