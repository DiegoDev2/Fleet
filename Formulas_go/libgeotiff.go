package main

// Libgeotiff - Library and tools for dealing with GeoTIFF
// Homepage: https://github.com/OSGeo/libgeotiff

import (
	"fmt"
	
	"os/exec"
)

func installLibgeotiff() {
	// Método 1: Descargar y extraer .tar.gz
	libgeotiff_tar_url := "https://github.com/OSGeo/libgeotiff/releases/download/1.7.3/libgeotiff-1.7.3.tar.gz"
	libgeotiff_cmd_tar := exec.Command("curl", "-L", libgeotiff_tar_url, "-o", "package.tar.gz")
	err := libgeotiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libgeotiff_zip_url := "https://github.com/OSGeo/libgeotiff/releases/download/1.7.3/libgeotiff-1.7.3.zip"
	libgeotiff_cmd_zip := exec.Command("curl", "-L", libgeotiff_zip_url, "-o", "package.zip")
	err = libgeotiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libgeotiff_bin_url := "https://github.com/OSGeo/libgeotiff/releases/download/1.7.3/libgeotiff-1.7.3.bin"
	libgeotiff_cmd_bin := exec.Command("curl", "-L", libgeotiff_bin_url, "-o", "binary.bin")
	err = libgeotiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libgeotiff_src_url := "https://github.com/OSGeo/libgeotiff/releases/download/1.7.3/libgeotiff-1.7.3.src.tar.gz"
	libgeotiff_cmd_src := exec.Command("curl", "-L", libgeotiff_src_url, "-o", "source.tar.gz")
	err = libgeotiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libgeotiff_cmd_direct := exec.Command("./binary")
	err = libgeotiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: proj")
exec.Command("latte", "install", "proj")
}
