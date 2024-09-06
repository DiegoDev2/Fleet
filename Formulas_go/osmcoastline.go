package main

// Osmcoastline - Extracts coastline data from OpenStreetMap planet file
// Homepage: https://osmcode.org/osmcoastline/

import (
	"fmt"
	
	"os/exec"
)

func installOsmcoastline() {
	// Método 1: Descargar y extraer .tar.gz
	osmcoastline_tar_url := "https://github.com/osmcode/osmcoastline/archive/refs/tags/v2.4.0.tar.gz"
	osmcoastline_cmd_tar := exec.Command("curl", "-L", osmcoastline_tar_url, "-o", "package.tar.gz")
	err := osmcoastline_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osmcoastline_zip_url := "https://github.com/osmcode/osmcoastline/archive/refs/tags/v2.4.0.zip"
	osmcoastline_cmd_zip := exec.Command("curl", "-L", osmcoastline_zip_url, "-o", "package.zip")
	err = osmcoastline_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osmcoastline_bin_url := "https://github.com/osmcode/osmcoastline/archive/refs/tags/v2.4.0.bin"
	osmcoastline_cmd_bin := exec.Command("curl", "-L", osmcoastline_bin_url, "-o", "binary.bin")
	err = osmcoastline_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osmcoastline_src_url := "https://github.com/osmcode/osmcoastline/archive/refs/tags/v2.4.0.src.tar.gz"
	osmcoastline_cmd_src := exec.Command("curl", "-L", osmcoastline_src_url, "-o", "source.tar.gz")
	err = osmcoastline_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osmcoastline_cmd_direct := exec.Command("./binary")
	err = osmcoastline_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libosmium")
exec.Command("latte", "install", "libosmium")
	fmt.Println("Instalando dependencia: expat")
exec.Command("latte", "install", "expat")
	fmt.Println("Instalando dependencia: gdal")
exec.Command("latte", "install", "gdal")
	fmt.Println("Instalando dependencia: geos")
exec.Command("latte", "install", "geos")
	fmt.Println("Instalando dependencia: libspatialite")
exec.Command("latte", "install", "libspatialite")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
}
