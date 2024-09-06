package main

// Pugixml - Light-weight C++ XML processing library
// Homepage: https://pugixml.org/

import (
	"fmt"
	
	"os/exec"
)

func installPugixml() {
	// Método 1: Descargar y extraer .tar.gz
	pugixml_tar_url := "https://github.com/zeux/pugixml/releases/download/v1.14/pugixml-1.14.tar.gz"
	pugixml_cmd_tar := exec.Command("curl", "-L", pugixml_tar_url, "-o", "package.tar.gz")
	err := pugixml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pugixml_zip_url := "https://github.com/zeux/pugixml/releases/download/v1.14/pugixml-1.14.zip"
	pugixml_cmd_zip := exec.Command("curl", "-L", pugixml_zip_url, "-o", "package.zip")
	err = pugixml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pugixml_bin_url := "https://github.com/zeux/pugixml/releases/download/v1.14/pugixml-1.14.bin"
	pugixml_cmd_bin := exec.Command("curl", "-L", pugixml_bin_url, "-o", "binary.bin")
	err = pugixml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pugixml_src_url := "https://github.com/zeux/pugixml/releases/download/v1.14/pugixml-1.14.src.tar.gz"
	pugixml_cmd_src := exec.Command("curl", "-L", pugixml_src_url, "-o", "source.tar.gz")
	err = pugixml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pugixml_cmd_direct := exec.Command("./binary")
	err = pugixml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
