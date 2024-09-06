package main

// Libtiff - TIFF library and utilities
// Homepage: https://libtiff.gitlab.io/libtiff/

import (
	"fmt"
	
	"os/exec"
)

func installLibtiff() {
	// Método 1: Descargar y extraer .tar.gz
	libtiff_tar_url := "https://download.osgeo.org/libtiff/tiff-4.6.0.tar.gz"
	libtiff_cmd_tar := exec.Command("curl", "-L", libtiff_tar_url, "-o", "package.tar.gz")
	err := libtiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtiff_zip_url := "https://download.osgeo.org/libtiff/tiff-4.6.0.zip"
	libtiff_cmd_zip := exec.Command("curl", "-L", libtiff_zip_url, "-o", "package.zip")
	err = libtiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtiff_bin_url := "https://download.osgeo.org/libtiff/tiff-4.6.0.bin"
	libtiff_cmd_bin := exec.Command("curl", "-L", libtiff_bin_url, "-o", "binary.bin")
	err = libtiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtiff_src_url := "https://download.osgeo.org/libtiff/tiff-4.6.0.src.tar.gz"
	libtiff_cmd_src := exec.Command("curl", "-L", libtiff_src_url, "-o", "source.tar.gz")
	err = libtiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtiff_cmd_direct := exec.Command("./binary")
	err = libtiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
