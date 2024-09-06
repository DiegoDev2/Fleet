package main

// Libmpeg2 - Library to decode mpeg-2 and mpeg-1 video streams
// Homepage: https://libmpeg2.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installLibmpeg2() {
	// Método 1: Descargar y extraer .tar.gz
	libmpeg2_tar_url := "https://libmpeg2.sourceforge.io/files/libmpeg2-0.5.1.tar.gz"
	libmpeg2_cmd_tar := exec.Command("curl", "-L", libmpeg2_tar_url, "-o", "package.tar.gz")
	err := libmpeg2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmpeg2_zip_url := "https://libmpeg2.sourceforge.io/files/libmpeg2-0.5.1.zip"
	libmpeg2_cmd_zip := exec.Command("curl", "-L", libmpeg2_zip_url, "-o", "package.zip")
	err = libmpeg2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmpeg2_bin_url := "https://libmpeg2.sourceforge.io/files/libmpeg2-0.5.1.bin"
	libmpeg2_cmd_bin := exec.Command("curl", "-L", libmpeg2_bin_url, "-o", "binary.bin")
	err = libmpeg2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmpeg2_src_url := "https://libmpeg2.sourceforge.io/files/libmpeg2-0.5.1.src.tar.gz"
	libmpeg2_cmd_src := exec.Command("curl", "-L", libmpeg2_src_url, "-o", "source.tar.gz")
	err = libmpeg2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmpeg2_cmd_direct := exec.Command("./binary")
	err = libmpeg2_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
}
