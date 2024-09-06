package main

// Flac123 - Command-line program for playing FLAC audio files
// Homepage: https://github.com/flac123/flac123

import (
	"fmt"
	
	"os/exec"
)

func installFlac123() {
	// Método 1: Descargar y extraer .tar.gz
	flac123_tar_url := "https://github.com/flac123/flac123/archive/refs/tags/v2.1.1.tar.gz"
	flac123_cmd_tar := exec.Command("curl", "-L", flac123_tar_url, "-o", "package.tar.gz")
	err := flac123_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flac123_zip_url := "https://github.com/flac123/flac123/archive/refs/tags/v2.1.1.zip"
	flac123_cmd_zip := exec.Command("curl", "-L", flac123_zip_url, "-o", "package.zip")
	err = flac123_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flac123_bin_url := "https://github.com/flac123/flac123/archive/refs/tags/v2.1.1.bin"
	flac123_cmd_bin := exec.Command("curl", "-L", flac123_bin_url, "-o", "binary.bin")
	err = flac123_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flac123_src_url := "https://github.com/flac123/flac123/archive/refs/tags/v2.1.1.src.tar.gz"
	flac123_cmd_src := exec.Command("curl", "-L", flac123_src_url, "-o", "source.tar.gz")
	err = flac123_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flac123_cmd_direct := exec.Command("./binary")
	err = flac123_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: libao")
	exec.Command("latte", "install", "libao").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: popt")
	exec.Command("latte", "install", "popt").Run()
}
