package main

// Osmfilter - Command-line tool to filter OpenStreetMap files for specific tags
// Homepage: https://wiki.openstreetmap.org/wiki/Osmfilter

import (
	"fmt"
	
	"os/exec"
)

func installOsmfilter() {
	// Método 1: Descargar y extraer .tar.gz
	osmfilter_tar_url := "https://gitlab.com/osm-c-tools/osmctools.git"
	osmfilter_cmd_tar := exec.Command("curl", "-L", osmfilter_tar_url, "-o", "package.tar.gz")
	err := osmfilter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osmfilter_zip_url := "https://gitlab.com/osm-c-tools/osmctools.git"
	osmfilter_cmd_zip := exec.Command("curl", "-L", osmfilter_zip_url, "-o", "package.zip")
	err = osmfilter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osmfilter_bin_url := "https://gitlab.com/osm-c-tools/osmctools.git"
	osmfilter_cmd_bin := exec.Command("curl", "-L", osmfilter_bin_url, "-o", "binary.bin")
	err = osmfilter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osmfilter_src_url := "https://gitlab.com/osm-c-tools/osmctools.git"
	osmfilter_cmd_src := exec.Command("curl", "-L", osmfilter_src_url, "-o", "source.tar.gz")
	err = osmfilter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osmfilter_cmd_direct := exec.Command("./binary")
	err = osmfilter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
