package main

// Spatialindex - General framework for developing spatial indices
// Homepage: https://libspatialindex.org/

import (
	"fmt"
	
	"os/exec"
)

func installSpatialindex() {
	// Método 1: Descargar y extraer .tar.gz
	spatialindex_tar_url := "https://github.com/libspatialindex/libspatialindex/releases/download/2.0.0/spatialindex-src-2.0.0.tar.bz2"
	spatialindex_cmd_tar := exec.Command("curl", "-L", spatialindex_tar_url, "-o", "package.tar.gz")
	err := spatialindex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	spatialindex_zip_url := "https://github.com/libspatialindex/libspatialindex/releases/download/2.0.0/spatialindex-src-2.0.0.tar.bz2"
	spatialindex_cmd_zip := exec.Command("curl", "-L", spatialindex_zip_url, "-o", "package.zip")
	err = spatialindex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	spatialindex_bin_url := "https://github.com/libspatialindex/libspatialindex/releases/download/2.0.0/spatialindex-src-2.0.0.tar.bz2"
	spatialindex_cmd_bin := exec.Command("curl", "-L", spatialindex_bin_url, "-o", "binary.bin")
	err = spatialindex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	spatialindex_src_url := "https://github.com/libspatialindex/libspatialindex/releases/download/2.0.0/spatialindex-src-2.0.0.tar.bz2"
	spatialindex_cmd_src := exec.Command("curl", "-L", spatialindex_src_url, "-o", "source.tar.gz")
	err = spatialindex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	spatialindex_cmd_direct := exec.Command("./binary")
	err = spatialindex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
