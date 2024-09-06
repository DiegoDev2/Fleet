package main

// Huexpress - PC Engine emulator
// Homepage: https://github.com/kallisti5/huexpress

import (
	"fmt"
	
	"os/exec"
)

func installHuexpress() {
	// Método 1: Descargar y extraer .tar.gz
	huexpress_tar_url := "https://github.com/kallisti5/huexpress/archive/refs/tags/3.0.4.tar.gz"
	huexpress_cmd_tar := exec.Command("curl", "-L", huexpress_tar_url, "-o", "package.tar.gz")
	err := huexpress_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	huexpress_zip_url := "https://github.com/kallisti5/huexpress/archive/refs/tags/3.0.4.zip"
	huexpress_cmd_zip := exec.Command("curl", "-L", huexpress_zip_url, "-o", "package.zip")
	err = huexpress_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	huexpress_bin_url := "https://github.com/kallisti5/huexpress/archive/refs/tags/3.0.4.bin"
	huexpress_cmd_bin := exec.Command("curl", "-L", huexpress_bin_url, "-o", "binary.bin")
	err = huexpress_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	huexpress_src_url := "https://github.com/kallisti5/huexpress/archive/refs/tags/3.0.4.src.tar.gz"
	huexpress_cmd_src := exec.Command("curl", "-L", huexpress_src_url, "-o", "source.tar.gz")
	err = huexpress_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	huexpress_cmd_direct := exec.Command("./binary")
	err = huexpress_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: scons")
	exec.Command("latte", "install", "scons").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: libzip")
	exec.Command("latte", "install", "libzip").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: sdl2_mixer")
	exec.Command("latte", "install", "sdl2_mixer").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: mesa-glu")
	exec.Command("latte", "install", "mesa-glu").Run()
}
