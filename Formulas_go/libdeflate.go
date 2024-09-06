package main

// Libdeflate - Heavily optimized DEFLATE/zlib/gzip compression and decompression
// Homepage: https://github.com/ebiggers/libdeflate

import (
	"fmt"
	
	"os/exec"
)

func installLibdeflate() {
	// Método 1: Descargar y extraer .tar.gz
	libdeflate_tar_url := "https://github.com/ebiggers/libdeflate/archive/refs/tags/v1.21.tar.gz"
	libdeflate_cmd_tar := exec.Command("curl", "-L", libdeflate_tar_url, "-o", "package.tar.gz")
	err := libdeflate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdeflate_zip_url := "https://github.com/ebiggers/libdeflate/archive/refs/tags/v1.21.zip"
	libdeflate_cmd_zip := exec.Command("curl", "-L", libdeflate_zip_url, "-o", "package.zip")
	err = libdeflate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdeflate_bin_url := "https://github.com/ebiggers/libdeflate/archive/refs/tags/v1.21.bin"
	libdeflate_cmd_bin := exec.Command("curl", "-L", libdeflate_bin_url, "-o", "binary.bin")
	err = libdeflate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdeflate_src_url := "https://github.com/ebiggers/libdeflate/archive/refs/tags/v1.21.src.tar.gz"
	libdeflate_cmd_src := exec.Command("curl", "-L", libdeflate_src_url, "-o", "source.tar.gz")
	err = libdeflate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdeflate_cmd_direct := exec.Command("./binary")
	err = libdeflate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
