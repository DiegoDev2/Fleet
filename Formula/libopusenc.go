package main

// Libopusenc - Convenience library for creating .opus files
// Homepage: https://opus-codec.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibopusenc() {
	// Método 1: Descargar y extraer .tar.gz
	libopusenc_tar_url := "https://archive.mozilla.org/pub/opus/libopusenc-0.2.1.tar.gz"
	libopusenc_cmd_tar := exec.Command("curl", "-L", libopusenc_tar_url, "-o", "package.tar.gz")
	err := libopusenc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libopusenc_zip_url := "https://archive.mozilla.org/pub/opus/libopusenc-0.2.1.zip"
	libopusenc_cmd_zip := exec.Command("curl", "-L", libopusenc_zip_url, "-o", "package.zip")
	err = libopusenc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libopusenc_bin_url := "https://archive.mozilla.org/pub/opus/libopusenc-0.2.1.bin"
	libopusenc_cmd_bin := exec.Command("curl", "-L", libopusenc_bin_url, "-o", "binary.bin")
	err = libopusenc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libopusenc_src_url := "https://archive.mozilla.org/pub/opus/libopusenc-0.2.1.src.tar.gz"
	libopusenc_cmd_src := exec.Command("curl", "-L", libopusenc_src_url, "-o", "source.tar.gz")
	err = libopusenc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libopusenc_cmd_direct := exec.Command("./binary")
	err = libopusenc_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: opus")
	exec.Command("latte", "install", "opus").Run()
}
