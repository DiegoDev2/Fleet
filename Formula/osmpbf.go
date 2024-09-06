package main

// OsmPbf - Tools related to PBF (an alternative to XML format)
// Homepage: https://wiki.openstreetmap.org/wiki/PBF_Format

import (
	"fmt"
	
	"os/exec"
)

func installOsmPbf() {
	// Método 1: Descargar y extraer .tar.gz
	osmpbf_tar_url := "https://github.com/openstreetmap/OSM-binary/archive/refs/tags/v1.5.1.tar.gz"
	osmpbf_cmd_tar := exec.Command("curl", "-L", osmpbf_tar_url, "-o", "package.tar.gz")
	err := osmpbf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osmpbf_zip_url := "https://github.com/openstreetmap/OSM-binary/archive/refs/tags/v1.5.1.zip"
	osmpbf_cmd_zip := exec.Command("curl", "-L", osmpbf_zip_url, "-o", "package.zip")
	err = osmpbf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osmpbf_bin_url := "https://github.com/openstreetmap/OSM-binary/archive/refs/tags/v1.5.1.bin"
	osmpbf_cmd_bin := exec.Command("curl", "-L", osmpbf_bin_url, "-o", "binary.bin")
	err = osmpbf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osmpbf_src_url := "https://github.com/openstreetmap/OSM-binary/archive/refs/tags/v1.5.1.src.tar.gz"
	osmpbf_cmd_src := exec.Command("curl", "-L", osmpbf_src_url, "-o", "source.tar.gz")
	err = osmpbf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osmpbf_cmd_direct := exec.Command("./binary")
	err = osmpbf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
}
