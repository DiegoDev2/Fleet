package main

// Libkate - Overlay codec for multiplexed audio/video in Ogg
// Homepage: https://code.google.com/archive/p/libkate/

import (
	"fmt"
	
	"os/exec"
)

func installLibkate() {
	// Método 1: Descargar y extraer .tar.gz
	libkate_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/libkate/libkate-0.4.1.tar.gz"
	libkate_cmd_tar := exec.Command("curl", "-L", libkate_tar_url, "-o", "package.tar.gz")
	err := libkate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libkate_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/libkate/libkate-0.4.1.zip"
	libkate_cmd_zip := exec.Command("curl", "-L", libkate_zip_url, "-o", "package.zip")
	err = libkate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libkate_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/libkate/libkate-0.4.1.bin"
	libkate_cmd_bin := exec.Command("curl", "-L", libkate_bin_url, "-o", "binary.bin")
	err = libkate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libkate_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/libkate/libkate-0.4.1.src.tar.gz"
	libkate_cmd_src := exec.Command("curl", "-L", libkate_src_url, "-o", "source.tar.gz")
	err = libkate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libkate_cmd_direct := exec.Command("./binary")
	err = libkate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libogg")
exec.Command("latte", "install", "libogg")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
