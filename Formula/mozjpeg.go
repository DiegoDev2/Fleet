package main

// Mozjpeg - Improved JPEG encoder
// Homepage: https://github.com/mozilla/mozjpeg

import (
	"fmt"
	
	"os/exec"
)

func installMozjpeg() {
	// Método 1: Descargar y extraer .tar.gz
	mozjpeg_tar_url := "https://github.com/mozilla/mozjpeg/archive/refs/tags/v4.1.5.tar.gz"
	mozjpeg_cmd_tar := exec.Command("curl", "-L", mozjpeg_tar_url, "-o", "package.tar.gz")
	err := mozjpeg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mozjpeg_zip_url := "https://github.com/mozilla/mozjpeg/archive/refs/tags/v4.1.5.zip"
	mozjpeg_cmd_zip := exec.Command("curl", "-L", mozjpeg_zip_url, "-o", "package.zip")
	err = mozjpeg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mozjpeg_bin_url := "https://github.com/mozilla/mozjpeg/archive/refs/tags/v4.1.5.bin"
	mozjpeg_cmd_bin := exec.Command("curl", "-L", mozjpeg_bin_url, "-o", "binary.bin")
	err = mozjpeg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mozjpeg_src_url := "https://github.com/mozilla/mozjpeg/archive/refs/tags/v4.1.5.src.tar.gz"
	mozjpeg_cmd_src := exec.Command("curl", "-L", mozjpeg_src_url, "-o", "source.tar.gz")
	err = mozjpeg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mozjpeg_cmd_direct := exec.Command("./binary")
	err = mozjpeg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: nasm")
	exec.Command("latte", "install", "nasm").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
