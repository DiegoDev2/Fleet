package main

// Xsd - XML Data Binding for C++
// Homepage: https://www.codesynthesis.com/products/xsd/

import (
	"fmt"
	
	"os/exec"
)

func installXsd() {
	// Método 1: Descargar y extraer .tar.gz
	xsd_tar_url := "https://www.codesynthesis.com/download/xsd/4.0/xsd-4.0.0+dep.tar.bz2"
	xsd_cmd_tar := exec.Command("curl", "-L", xsd_tar_url, "-o", "package.tar.gz")
	err := xsd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xsd_zip_url := "https://www.codesynthesis.com/download/xsd/4.0/xsd-4.0.0+dep.tar.bz2"
	xsd_cmd_zip := exec.Command("curl", "-L", xsd_zip_url, "-o", "package.zip")
	err = xsd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xsd_bin_url := "https://www.codesynthesis.com/download/xsd/4.0/xsd-4.0.0+dep.tar.bz2"
	xsd_cmd_bin := exec.Command("curl", "-L", xsd_bin_url, "-o", "binary.bin")
	err = xsd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xsd_src_url := "https://www.codesynthesis.com/download/xsd/4.0/xsd-4.0.0+dep.tar.bz2"
	xsd_cmd_src := exec.Command("curl", "-L", xsd_src_url, "-o", "source.tar.gz")
	err = xsd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xsd_cmd_direct := exec.Command("./binary")
	err = xsd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: xerces-c")
	exec.Command("latte", "install", "xerces-c").Run()
}
