package main

// Libopenmpt - Software library to decode tracked music files
// Homepage: https://lib.openmpt.org/libopenmpt/

import (
	"fmt"
	
	"os/exec"
)

func installLibopenmpt() {
	// Método 1: Descargar y extraer .tar.gz
	libopenmpt_tar_url := "https://lib.openmpt.org/files/libopenmpt/src/libopenmpt-0.7.9+release.autotools.tar.gz"
	libopenmpt_cmd_tar := exec.Command("curl", "-L", libopenmpt_tar_url, "-o", "package.tar.gz")
	err := libopenmpt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libopenmpt_zip_url := "https://lib.openmpt.org/files/libopenmpt/src/libopenmpt-0.7.9+release.autotools.zip"
	libopenmpt_cmd_zip := exec.Command("curl", "-L", libopenmpt_zip_url, "-o", "package.zip")
	err = libopenmpt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libopenmpt_bin_url := "https://lib.openmpt.org/files/libopenmpt/src/libopenmpt-0.7.9+release.autotools.bin"
	libopenmpt_cmd_bin := exec.Command("curl", "-L", libopenmpt_bin_url, "-o", "binary.bin")
	err = libopenmpt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libopenmpt_src_url := "https://lib.openmpt.org/files/libopenmpt/src/libopenmpt-0.7.9+release.autotools.src.tar.gz"
	libopenmpt_cmd_src := exec.Command("curl", "-L", libopenmpt_src_url, "-o", "source.tar.gz")
	err = libopenmpt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libopenmpt_cmd_direct := exec.Command("./binary")
	err = libopenmpt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: flac")
exec.Command("latte", "install", "flac")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
	fmt.Println("Instalando dependencia: libvorbis")
exec.Command("latte", "install", "libvorbis")
	fmt.Println("Instalando dependencia: mpg123")
exec.Command("latte", "install", "mpg123")
	fmt.Println("Instalando dependencia: portaudio")
exec.Command("latte", "install", "portaudio")
	fmt.Println("Instalando dependencia: pulseaudio")
exec.Command("latte", "install", "pulseaudio")
}
