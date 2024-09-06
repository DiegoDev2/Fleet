package main

// Libavif - Library for encoding and decoding .avif files
// Homepage: https://github.com/AOMediaCodec/libavif

import (
	"fmt"
	
	"os/exec"
)

func installLibavif() {
	// Método 1: Descargar y extraer .tar.gz
	libavif_tar_url := "https://github.com/AOMediaCodec/libavif/archive/refs/tags/v1.1.1.tar.gz"
	libavif_cmd_tar := exec.Command("curl", "-L", libavif_tar_url, "-o", "package.tar.gz")
	err := libavif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libavif_zip_url := "https://github.com/AOMediaCodec/libavif/archive/refs/tags/v1.1.1.zip"
	libavif_cmd_zip := exec.Command("curl", "-L", libavif_zip_url, "-o", "package.zip")
	err = libavif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libavif_bin_url := "https://github.com/AOMediaCodec/libavif/archive/refs/tags/v1.1.1.bin"
	libavif_cmd_bin := exec.Command("curl", "-L", libavif_bin_url, "-o", "binary.bin")
	err = libavif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libavif_src_url := "https://github.com/AOMediaCodec/libavif/archive/refs/tags/v1.1.1.src.tar.gz"
	libavif_cmd_src := exec.Command("curl", "-L", libavif_src_url, "-o", "source.tar.gz")
	err = libavif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libavif_cmd_direct := exec.Command("./binary")
	err = libavif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: nasm")
exec.Command("latte", "install", "nasm")
	fmt.Println("Instalando dependencia: aom")
exec.Command("latte", "install", "aom")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
