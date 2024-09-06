package main

// Fceux - All-in-one NES/Famicom Emulator
// Homepage: https://fceux.com/

import (
	"fmt"
	
	"os/exec"
)

func installFceux() {
	// Método 1: Descargar y extraer .tar.gz
	fceux_tar_url := "https://github.com/TASEmulators/fceux.git"
	fceux_cmd_tar := exec.Command("curl", "-L", fceux_tar_url, "-o", "package.tar.gz")
	err := fceux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fceux_zip_url := "https://github.com/TASEmulators/fceux.git"
	fceux_cmd_zip := exec.Command("curl", "-L", fceux_zip_url, "-o", "package.zip")
	err = fceux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fceux_bin_url := "https://github.com/TASEmulators/fceux.git"
	fceux_cmd_bin := exec.Command("curl", "-L", fceux_bin_url, "-o", "binary.bin")
	err = fceux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fceux_src_url := "https://github.com/TASEmulators/fceux.git"
	fceux_cmd_src := exec.Command("curl", "-L", fceux_src_url, "-o", "source.tar.gz")
	err = fceux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fceux_cmd_direct := exec.Command("./binary")
	err = fceux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: libarchive")
	exec.Command("latte", "install", "libarchive").Run()
	fmt.Println("Instalando dependencia: minizip")
	exec.Command("latte", "install", "minizip").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: x264")
	exec.Command("latte", "install", "x264").Run()
	fmt.Println("Instalando dependencia: x265")
	exec.Command("latte", "install", "x265").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
}
