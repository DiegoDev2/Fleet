package main

// Frotz - Infocom-style interactive fiction player
// Homepage: https://661.org/proj/if/frotz/

import (
	"fmt"
	
	"os/exec"
)

func installFrotz() {
	// Método 1: Descargar y extraer .tar.gz
	frotz_tar_url := "https://gitlab.com/DavidGriffith/frotz/-/archive/2.54/frotz-2.54.tar.bz2"
	frotz_cmd_tar := exec.Command("curl", "-L", frotz_tar_url, "-o", "package.tar.gz")
	err := frotz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	frotz_zip_url := "https://gitlab.com/DavidGriffith/frotz/-/archive/2.54/frotz-2.54.tar.bz2"
	frotz_cmd_zip := exec.Command("curl", "-L", frotz_zip_url, "-o", "package.zip")
	err = frotz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	frotz_bin_url := "https://gitlab.com/DavidGriffith/frotz/-/archive/2.54/frotz-2.54.tar.bz2"
	frotz_cmd_bin := exec.Command("curl", "-L", frotz_bin_url, "-o", "binary.bin")
	err = frotz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	frotz_src_url := "https://gitlab.com/DavidGriffith/frotz/-/archive/2.54/frotz-2.54.tar.bz2"
	frotz_cmd_src := exec.Command("curl", "-L", frotz_src_url, "-o", "source.tar.gz")
	err = frotz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	frotz_cmd_direct := exec.Command("./binary")
	err = frotz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libao")
exec.Command("latte", "install", "libao")
	fmt.Println("Instalando dependencia: libmodplug")
exec.Command("latte", "install", "libmodplug")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_mixer")
exec.Command("latte", "install", "sdl2_mixer")
}
