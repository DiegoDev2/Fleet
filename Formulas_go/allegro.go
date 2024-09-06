package main

// Allegro - C/C++ multimedia library for cross-platform game development
// Homepage: https://liballeg.org/

import (
	"fmt"
	
	"os/exec"
)

func installAllegro() {
	// Método 1: Descargar y extraer .tar.gz
	allegro_tar_url := "https://github.com/liballeg/allegro5/releases/download/5.2.9.1/allegro-5.2.9.1.tar.gz"
	allegro_cmd_tar := exec.Command("curl", "-L", allegro_tar_url, "-o", "package.tar.gz")
	err := allegro_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	allegro_zip_url := "https://github.com/liballeg/allegro5/releases/download/5.2.9.1/allegro-5.2.9.1.zip"
	allegro_cmd_zip := exec.Command("curl", "-L", allegro_zip_url, "-o", "package.zip")
	err = allegro_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	allegro_bin_url := "https://github.com/liballeg/allegro5/releases/download/5.2.9.1/allegro-5.2.9.1.bin"
	allegro_cmd_bin := exec.Command("curl", "-L", allegro_bin_url, "-o", "binary.bin")
	err = allegro_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	allegro_src_url := "https://github.com/liballeg/allegro5/releases/download/5.2.9.1/allegro-5.2.9.1.src.tar.gz"
	allegro_cmd_src := exec.Command("curl", "-L", allegro_src_url, "-o", "source.tar.gz")
	err = allegro_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	allegro_cmd_direct := exec.Command("./binary")
	err = allegro_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: dumb")
exec.Command("latte", "install", "dumb")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: opusfile")
exec.Command("latte", "install", "opusfile")
	fmt.Println("Instalando dependencia: physfs")
exec.Command("latte", "install", "physfs")
	fmt.Println("Instalando dependencia: theora")
exec.Command("latte", "install", "theora")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
	fmt.Println("Instalando dependencia: opus")
exec.Command("latte", "install", "opus")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libx11")
exec.Command("latte", "install", "libx11")
	fmt.Println("Instalando dependencia: libxcursor")
exec.Command("latte", "install", "libxcursor")
	fmt.Println("Instalando dependencia: libxi")
exec.Command("latte", "install", "libxi")
	fmt.Println("Instalando dependencia: libxinerama")
exec.Command("latte", "install", "libxinerama")
	fmt.Println("Instalando dependencia: libxrandr")
exec.Command("latte", "install", "libxrandr")
	fmt.Println("Instalando dependencia: libxscrnsaver")
exec.Command("latte", "install", "libxscrnsaver")
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
	fmt.Println("Instalando dependencia: mesa-glu")
exec.Command("latte", "install", "mesa-glu")
}
