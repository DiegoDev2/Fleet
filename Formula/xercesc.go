package main

// XercesC - Validating XML parser
// Homepage: https://xerces.apache.org/xerces-c/

import (
	"fmt"
	
	"os/exec"
)

func installXercesC() {
	// Método 1: Descargar y extraer .tar.gz
	xercesc_tar_url := "https://www.apache.org/dyn/closer.lua?path=xerces/c/3/sources/xerces-c-3.2.5.tar.gz"
	xercesc_cmd_tar := exec.Command("curl", "-L", xercesc_tar_url, "-o", "package.tar.gz")
	err := xercesc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xercesc_zip_url := "https://www.apache.org/dyn/closer.lua?path=xerces/c/3/sources/xerces-c-3.2.5.zip"
	xercesc_cmd_zip := exec.Command("curl", "-L", xercesc_zip_url, "-o", "package.zip")
	err = xercesc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xercesc_bin_url := "https://www.apache.org/dyn/closer.lua?path=xerces/c/3/sources/xerces-c-3.2.5.bin"
	xercesc_cmd_bin := exec.Command("curl", "-L", xercesc_bin_url, "-o", "binary.bin")
	err = xercesc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xercesc_src_url := "https://www.apache.org/dyn/closer.lua?path=xerces/c/3/sources/xerces-c-3.2.5.src.tar.gz"
	xercesc_cmd_src := exec.Command("curl", "-L", xercesc_src_url, "-o", "source.tar.gz")
	err = xercesc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xercesc_cmd_direct := exec.Command("./binary")
	err = xercesc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
