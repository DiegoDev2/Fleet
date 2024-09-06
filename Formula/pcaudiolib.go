package main

// Pcaudiolib - Portable C Audio Library
// Homepage: https://github.com/espeak-ng/pcaudiolib

import (
	"fmt"
	
	"os/exec"
)

func installPcaudiolib() {
	// Método 1: Descargar y extraer .tar.gz
	pcaudiolib_tar_url := "https://github.com/espeak-ng/pcaudiolib/releases/download/1.2/pcaudiolib-1.2.tar.gz"
	pcaudiolib_cmd_tar := exec.Command("curl", "-L", pcaudiolib_tar_url, "-o", "package.tar.gz")
	err := pcaudiolib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcaudiolib_zip_url := "https://github.com/espeak-ng/pcaudiolib/releases/download/1.2/pcaudiolib-1.2.zip"
	pcaudiolib_cmd_zip := exec.Command("curl", "-L", pcaudiolib_zip_url, "-o", "package.zip")
	err = pcaudiolib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcaudiolib_bin_url := "https://github.com/espeak-ng/pcaudiolib/releases/download/1.2/pcaudiolib-1.2.bin"
	pcaudiolib_cmd_bin := exec.Command("curl", "-L", pcaudiolib_bin_url, "-o", "binary.bin")
	err = pcaudiolib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcaudiolib_src_url := "https://github.com/espeak-ng/pcaudiolib/releases/download/1.2/pcaudiolib-1.2.src.tar.gz"
	pcaudiolib_cmd_src := exec.Command("curl", "-L", pcaudiolib_src_url, "-o", "source.tar.gz")
	err = pcaudiolib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcaudiolib_cmd_direct := exec.Command("./binary")
	err = pcaudiolib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: alsa-lib")
	exec.Command("latte", "install", "alsa-lib").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
}
