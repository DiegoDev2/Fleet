package main

// ManifestTool - Command-line tool to create and query container image manifest list/indexes
// Homepage: https://github.com/estesp/manifest-tool/

import (
	"fmt"
	
	"os/exec"
)

func installManifestTool() {
	// Método 1: Descargar y extraer .tar.gz
	manifesttool_tar_url := "https://github.com/estesp/manifest-tool/archive/refs/tags/v2.1.7.tar.gz"
	manifesttool_cmd_tar := exec.Command("curl", "-L", manifesttool_tar_url, "-o", "package.tar.gz")
	err := manifesttool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	manifesttool_zip_url := "https://github.com/estesp/manifest-tool/archive/refs/tags/v2.1.7.zip"
	manifesttool_cmd_zip := exec.Command("curl", "-L", manifesttool_zip_url, "-o", "package.zip")
	err = manifesttool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	manifesttool_bin_url := "https://github.com/estesp/manifest-tool/archive/refs/tags/v2.1.7.bin"
	manifesttool_cmd_bin := exec.Command("curl", "-L", manifesttool_bin_url, "-o", "binary.bin")
	err = manifesttool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	manifesttool_src_url := "https://github.com/estesp/manifest-tool/archive/refs/tags/v2.1.7.src.tar.gz"
	manifesttool_cmd_src := exec.Command("curl", "-L", manifesttool_src_url, "-o", "source.tar.gz")
	err = manifesttool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	manifesttool_cmd_direct := exec.Command("./binary")
	err = manifesttool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
