package main

// Atari800 - Atari 8-bit machine emulator
// Homepage: https://atari800.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installAtari800() {
	// Método 1: Descargar y extraer .tar.gz
	atari800_tar_url := "https://github.com/atari800/atari800/releases/download/ATARI800_5_2_0/atari800-5.2.0-src.tgz"
	atari800_cmd_tar := exec.Command("curl", "-L", atari800_tar_url, "-o", "package.tar.gz")
	err := atari800_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	atari800_zip_url := "https://github.com/atari800/atari800/releases/download/ATARI800_5_2_0/atari800-5.2.0-src.tgz"
	atari800_cmd_zip := exec.Command("curl", "-L", atari800_zip_url, "-o", "package.zip")
	err = atari800_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	atari800_bin_url := "https://github.com/atari800/atari800/releases/download/ATARI800_5_2_0/atari800-5.2.0-src.tgz"
	atari800_cmd_bin := exec.Command("curl", "-L", atari800_bin_url, "-o", "binary.bin")
	err = atari800_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	atari800_src_url := "https://github.com/atari800/atari800/releases/download/ATARI800_5_2_0/atari800-5.2.0-src.tgz"
	atari800_cmd_src := exec.Command("curl", "-L", atari800_src_url, "-o", "source.tar.gz")
	err = atari800_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	atari800_cmd_direct := exec.Command("./binary")
	err = atari800_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
