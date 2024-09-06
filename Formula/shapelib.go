package main

// Shapelib - Library for reading and writing ArcView Shapefiles
// Homepage: http://shapelib.maptools.org/

import (
	"fmt"
	
	"os/exec"
)

func installShapelib() {
	// Método 1: Descargar y extraer .tar.gz
	shapelib_tar_url := "https://download.osgeo.org/shapelib/shapelib-1.6.1.tar.gz"
	shapelib_cmd_tar := exec.Command("curl", "-L", shapelib_tar_url, "-o", "package.tar.gz")
	err := shapelib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shapelib_zip_url := "https://download.osgeo.org/shapelib/shapelib-1.6.1.zip"
	shapelib_cmd_zip := exec.Command("curl", "-L", shapelib_zip_url, "-o", "package.zip")
	err = shapelib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shapelib_bin_url := "https://download.osgeo.org/shapelib/shapelib-1.6.1.bin"
	shapelib_cmd_bin := exec.Command("curl", "-L", shapelib_bin_url, "-o", "binary.bin")
	err = shapelib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shapelib_src_url := "https://download.osgeo.org/shapelib/shapelib-1.6.1.src.tar.gz"
	shapelib_cmd_src := exec.Command("curl", "-L", shapelib_src_url, "-o", "source.tar.gz")
	err = shapelib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shapelib_cmd_direct := exec.Command("./binary")
	err = shapelib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
