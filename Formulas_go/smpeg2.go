package main

// Smpeg2 - SDL MPEG Player Library
// Homepage: https://icculus.org/smpeg/

import (
	"fmt"
	
	"os/exec"
)

func installSmpeg2() {
	// Método 1: Descargar y extraer .tar.gz
	smpeg2_tar_url := "https://github.com/icculus/smpeg/archive/refs/tags/release_2_0_0.tar.gz"
	smpeg2_cmd_tar := exec.Command("curl", "-L", smpeg2_tar_url, "-o", "package.tar.gz")
	err := smpeg2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smpeg2_zip_url := "https://github.com/icculus/smpeg/archive/refs/tags/release_2_0_0.zip"
	smpeg2_cmd_zip := exec.Command("curl", "-L", smpeg2_zip_url, "-o", "package.zip")
	err = smpeg2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smpeg2_bin_url := "https://github.com/icculus/smpeg/archive/refs/tags/release_2_0_0.bin"
	smpeg2_cmd_bin := exec.Command("curl", "-L", smpeg2_bin_url, "-o", "binary.bin")
	err = smpeg2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smpeg2_src_url := "https://github.com/icculus/smpeg/archive/refs/tags/release_2_0_0.src.tar.gz"
	smpeg2_cmd_src := exec.Command("curl", "-L", smpeg2_src_url, "-o", "source.tar.gz")
	err = smpeg2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smpeg2_cmd_direct := exec.Command("./binary")
	err = smpeg2_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: sdl2")
exec.Command("latte", "install", "sdl2")
}
