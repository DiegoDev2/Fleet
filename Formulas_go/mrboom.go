package main

// Mrboom - Eight player Bomberman clone
// Homepage: http://mrboom.mumblecore.org/

import (
	"fmt"
	
	"os/exec"
)

func installMrboom() {
	// Método 1: Descargar y extraer .tar.gz
	mrboom_tar_url := "https://github.com/Javanaise/mrboom-libretro/releases/download/5.5/MrBoom-src-5.5.tar.gz"
	mrboom_cmd_tar := exec.Command("curl", "-L", mrboom_tar_url, "-o", "package.tar.gz")
	err := mrboom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mrboom_zip_url := "https://github.com/Javanaise/mrboom-libretro/releases/download/5.5/MrBoom-src-5.5.zip"
	mrboom_cmd_zip := exec.Command("curl", "-L", mrboom_zip_url, "-o", "package.zip")
	err = mrboom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mrboom_bin_url := "https://github.com/Javanaise/mrboom-libretro/releases/download/5.5/MrBoom-src-5.5.bin"
	mrboom_cmd_bin := exec.Command("curl", "-L", mrboom_bin_url, "-o", "binary.bin")
	err = mrboom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mrboom_src_url := "https://github.com/Javanaise/mrboom-libretro/releases/download/5.5/MrBoom-src-5.5.src.tar.gz"
	mrboom_cmd_src := exec.Command("curl", "-L", mrboom_src_url, "-o", "source.tar.gz")
	err = mrboom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mrboom_cmd_direct := exec.Command("./binary")
	err = mrboom_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libmodplug")
exec.Command("latte", "install", "libmodplug")
	fmt.Println("Instalando dependencia: minizip")
exec.Command("latte", "install", "minizip")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
	fmt.Println("Instalando dependencia: sdl2_mixer")
exec.Command("latte", "install", "sdl2_mixer")
}
