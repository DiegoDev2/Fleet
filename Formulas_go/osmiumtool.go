package main

// OsmiumTool - Libosmium-based command-line tool for processing OpenStreetMap data
// Homepage: https://osmcode.org/osmium-tool/

import (
	"fmt"
	
	"os/exec"
)

func installOsmiumTool() {
	// Método 1: Descargar y extraer .tar.gz
	osmiumtool_tar_url := "https://github.com/osmcode/osmium-tool/archive/refs/tags/v1.16.0.tar.gz"
	osmiumtool_cmd_tar := exec.Command("curl", "-L", osmiumtool_tar_url, "-o", "package.tar.gz")
	err := osmiumtool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osmiumtool_zip_url := "https://github.com/osmcode/osmium-tool/archive/refs/tags/v1.16.0.zip"
	osmiumtool_cmd_zip := exec.Command("curl", "-L", osmiumtool_zip_url, "-o", "package.zip")
	err = osmiumtool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osmiumtool_bin_url := "https://github.com/osmcode/osmium-tool/archive/refs/tags/v1.16.0.bin"
	osmiumtool_cmd_bin := exec.Command("curl", "-L", osmiumtool_bin_url, "-o", "binary.bin")
	err = osmiumtool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osmiumtool_src_url := "https://github.com/osmcode/osmium-tool/archive/refs/tags/v1.16.0.src.tar.gz"
	osmiumtool_cmd_src := exec.Command("curl", "-L", osmiumtool_src_url, "-o", "source.tar.gz")
	err = osmiumtool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osmiumtool_cmd_direct := exec.Command("./binary")
	err = osmiumtool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libosmium")
exec.Command("latte", "install", "libosmium")
	fmt.Println("Instalando dependencia: pandoc")
exec.Command("latte", "install", "pandoc")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
}
