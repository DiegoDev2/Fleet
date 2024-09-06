package main

// Blahtexml - Converts equations into Math ML
// Homepage: https://github.com/gvanas/blahtexml

import (
	"fmt"
	
	"os/exec"
)

func installBlahtexml() {
	// Método 1: Descargar y extraer .tar.gz
	blahtexml_tar_url := "https://github.com/gvanas/blahtexml/archive/refs/tags/v1.0.tar.gz"
	blahtexml_cmd_tar := exec.Command("curl", "-L", blahtexml_tar_url, "-o", "package.tar.gz")
	err := blahtexml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blahtexml_zip_url := "https://github.com/gvanas/blahtexml/archive/refs/tags/v1.0.zip"
	blahtexml_cmd_zip := exec.Command("curl", "-L", blahtexml_zip_url, "-o", "package.zip")
	err = blahtexml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blahtexml_bin_url := "https://github.com/gvanas/blahtexml/archive/refs/tags/v1.0.bin"
	blahtexml_cmd_bin := exec.Command("curl", "-L", blahtexml_bin_url, "-o", "binary.bin")
	err = blahtexml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blahtexml_src_url := "https://github.com/gvanas/blahtexml/archive/refs/tags/v1.0.src.tar.gz"
	blahtexml_cmd_src := exec.Command("curl", "-L", blahtexml_src_url, "-o", "source.tar.gz")
	err = blahtexml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blahtexml_cmd_direct := exec.Command("./binary")
	err = blahtexml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xerces-c")
	exec.Command("latte", "install", "xerces-c").Run()
}
