package main

// Imposm3 - Imports OpenStreetMap data into PostgreSQL/PostGIS databases
// Homepage: https://imposm.org

import (
	"fmt"
	
	"os/exec"
)

func installImposm3() {
	// Método 1: Descargar y extraer .tar.gz
	imposm3_tar_url := "https://github.com/omniscale/imposm3/archive/refs/tags/v0.14.0.tar.gz"
	imposm3_cmd_tar := exec.Command("curl", "-L", imposm3_tar_url, "-o", "package.tar.gz")
	err := imposm3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	imposm3_zip_url := "https://github.com/omniscale/imposm3/archive/refs/tags/v0.14.0.zip"
	imposm3_cmd_zip := exec.Command("curl", "-L", imposm3_zip_url, "-o", "package.zip")
	err = imposm3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	imposm3_bin_url := "https://github.com/omniscale/imposm3/archive/refs/tags/v0.14.0.bin"
	imposm3_cmd_bin := exec.Command("curl", "-L", imposm3_bin_url, "-o", "binary.bin")
	err = imposm3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	imposm3_src_url := "https://github.com/omniscale/imposm3/archive/refs/tags/v0.14.0.src.tar.gz"
	imposm3_cmd_src := exec.Command("curl", "-L", imposm3_src_url, "-o", "source.tar.gz")
	err = imposm3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	imposm3_cmd_direct := exec.Command("./binary")
	err = imposm3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: osmium-tool")
exec.Command("latte", "install", "osmium-tool")
	fmt.Println("Instalando dependencia: geos")
exec.Command("latte", "install", "geos")
	fmt.Println("Instalando dependencia: leveldb")
exec.Command("latte", "install", "leveldb")
}
