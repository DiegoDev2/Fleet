package main

// Libsamplerate - Library for sample rate conversion of audio data
// Homepage: https://github.com/libsndfile/libsamplerate

import (
	"fmt"
	
	"os/exec"
)

func installLibsamplerate() {
	// Método 1: Descargar y extraer .tar.gz
	libsamplerate_tar_url := "https://github.com/libsndfile/libsamplerate/archive/refs/tags/0.2.2.tar.gz"
	libsamplerate_cmd_tar := exec.Command("curl", "-L", libsamplerate_tar_url, "-o", "package.tar.gz")
	err := libsamplerate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsamplerate_zip_url := "https://github.com/libsndfile/libsamplerate/archive/refs/tags/0.2.2.zip"
	libsamplerate_cmd_zip := exec.Command("curl", "-L", libsamplerate_zip_url, "-o", "package.zip")
	err = libsamplerate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsamplerate_bin_url := "https://github.com/libsndfile/libsamplerate/archive/refs/tags/0.2.2.bin"
	libsamplerate_cmd_bin := exec.Command("curl", "-L", libsamplerate_bin_url, "-o", "binary.bin")
	err = libsamplerate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsamplerate_src_url := "https://github.com/libsndfile/libsamplerate/archive/refs/tags/0.2.2.src.tar.gz"
	libsamplerate_cmd_src := exec.Command("curl", "-L", libsamplerate_src_url, "-o", "source.tar.gz")
	err = libsamplerate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsamplerate_cmd_direct := exec.Command("./binary")
	err = libsamplerate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
