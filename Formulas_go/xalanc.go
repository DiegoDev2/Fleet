package main

// XalanC - XSLT processor
// Homepage: https://apache.github.io/xalan-c/

import (
	"fmt"
	
	"os/exec"
)

func installXalanC() {
	// Método 1: Descargar y extraer .tar.gz
	xalanc_tar_url := "https://www.apache.org/dyn/closer.lua?path=xalan/xalan-c/sources/xalan_c-1.12.tar.gz"
	xalanc_cmd_tar := exec.Command("curl", "-L", xalanc_tar_url, "-o", "package.tar.gz")
	err := xalanc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xalanc_zip_url := "https://www.apache.org/dyn/closer.lua?path=xalan/xalan-c/sources/xalan_c-1.12.zip"
	xalanc_cmd_zip := exec.Command("curl", "-L", xalanc_zip_url, "-o", "package.zip")
	err = xalanc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xalanc_bin_url := "https://www.apache.org/dyn/closer.lua?path=xalan/xalan-c/sources/xalan_c-1.12.bin"
	xalanc_cmd_bin := exec.Command("curl", "-L", xalanc_bin_url, "-o", "binary.bin")
	err = xalanc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xalanc_src_url := "https://www.apache.org/dyn/closer.lua?path=xalan/xalan-c/sources/xalan_c-1.12.src.tar.gz"
	xalanc_cmd_src := exec.Command("curl", "-L", xalanc_src_url, "-o", "source.tar.gz")
	err = xalanc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xalanc_cmd_direct := exec.Command("./binary")
	err = xalanc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
}
