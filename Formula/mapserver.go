package main

// Mapserver - Publish spatial data and interactive mapping apps to the web
// Homepage: https://mapserver.org/

import (
	"fmt"
	
	"os/exec"
)

func installMapserver() {
	// Método 1: Descargar y extraer .tar.gz
	mapserver_tar_url := "https://download.osgeo.org/mapserver/mapserver-8.2.2.tar.gz"
	mapserver_cmd_tar := exec.Command("curl", "-L", mapserver_tar_url, "-o", "package.tar.gz")
	err := mapserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mapserver_zip_url := "https://download.osgeo.org/mapserver/mapserver-8.2.2.zip"
	mapserver_cmd_zip := exec.Command("curl", "-L", mapserver_zip_url, "-o", "package.zip")
	err = mapserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mapserver_bin_url := "https://download.osgeo.org/mapserver/mapserver-8.2.2.bin"
	mapserver_cmd_bin := exec.Command("curl", "-L", mapserver_bin_url, "-o", "binary.bin")
	err = mapserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mapserver_src_url := "https://download.osgeo.org/mapserver/mapserver-8.2.2.src.tar.gz"
	mapserver_cmd_src := exec.Command("curl", "-L", mapserver_src_url, "-o", "source.tar.gz")
	err = mapserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mapserver_cmd_direct := exec.Command("./binary")
	err = mapserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fcgi")
	exec.Command("latte", "install", "fcgi").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gd")
	exec.Command("latte", "install", "gd").Run()
	fmt.Println("Instalando dependencia: gdal")
	exec.Command("latte", "install", "gdal").Run()
	fmt.Println("Instalando dependencia: geos")
	exec.Command("latte", "install", "geos").Run()
	fmt.Println("Instalando dependencia: giflib")
	exec.Command("latte", "install", "giflib").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libpq")
	exec.Command("latte", "install", "libpq").Run()
	fmt.Println("Instalando dependencia: libxml2")
	exec.Command("latte", "install", "libxml2").Run()
	fmt.Println("Instalando dependencia: proj")
	exec.Command("latte", "install", "proj").Run()
	fmt.Println("Instalando dependencia: protobuf-c")
	exec.Command("latte", "install", "protobuf-c").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
