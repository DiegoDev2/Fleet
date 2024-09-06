package main

// Geos - Geometry Engine
// Homepage: https://libgeos.org/

import (
	"fmt"
	
	"os/exec"
)

func installGeos() {
	// Método 1: Descargar y extraer .tar.gz
	geos_tar_url := "https://download.osgeo.org/geos/geos-3.12.2.tar.bz2"
	geos_cmd_tar := exec.Command("curl", "-L", geos_tar_url, "-o", "package.tar.gz")
	err := geos_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	geos_zip_url := "https://download.osgeo.org/geos/geos-3.12.2.tar.bz2"
	geos_cmd_zip := exec.Command("curl", "-L", geos_zip_url, "-o", "package.zip")
	err = geos_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	geos_bin_url := "https://download.osgeo.org/geos/geos-3.12.2.tar.bz2"
	geos_cmd_bin := exec.Command("curl", "-L", geos_bin_url, "-o", "binary.bin")
	err = geos_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	geos_src_url := "https://download.osgeo.org/geos/geos-3.12.2.tar.bz2"
	geos_cmd_src := exec.Command("curl", "-L", geos_src_url, "-o", "source.tar.gz")
	err = geos_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	geos_cmd_direct := exec.Command("./binary")
	err = geos_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
