package main

// Osmosis - Command-line OpenStreetMap data processor
// Homepage: https://wiki.openstreetmap.org/wiki/Osmosis

import (
	"fmt"
	
	"os/exec"
)

func installOsmosis() {
	// Método 1: Descargar y extraer .tar.gz
	osmosis_tar_url := "https://github.com/openstreetmap/osmosis/releases/download/0.49.2/osmosis-0.49.2.tar"
	osmosis_cmd_tar := exec.Command("curl", "-L", osmosis_tar_url, "-o", "package.tar.gz")
	err := osmosis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osmosis_zip_url := "https://github.com/openstreetmap/osmosis/releases/download/0.49.2/osmosis-0.49.2.tar"
	osmosis_cmd_zip := exec.Command("curl", "-L", osmosis_zip_url, "-o", "package.zip")
	err = osmosis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osmosis_bin_url := "https://github.com/openstreetmap/osmosis/releases/download/0.49.2/osmosis-0.49.2.tar"
	osmosis_cmd_bin := exec.Command("curl", "-L", osmosis_bin_url, "-o", "binary.bin")
	err = osmosis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osmosis_src_url := "https://github.com/openstreetmap/osmosis/releases/download/0.49.2/osmosis-0.49.2.tar"
	osmosis_cmd_src := exec.Command("curl", "-L", osmosis_src_url, "-o", "source.tar.gz")
	err = osmosis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osmosis_cmd_direct := exec.Command("./binary")
	err = osmosis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
