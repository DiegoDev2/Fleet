package main

// Attr - Manipulate filesystem extended attributes
// Homepage: https://savannah.nongnu.org/projects/attr

import (
	"fmt"
	
	"os/exec"
)

func installAttr() {
	// Método 1: Descargar y extraer .tar.gz
	attr_tar_url := "https://download.savannah.nongnu.org/releases/attr/attr-2.5.2.tar.gz"
	attr_cmd_tar := exec.Command("curl", "-L", attr_tar_url, "-o", "package.tar.gz")
	err := attr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	attr_zip_url := "https://download.savannah.nongnu.org/releases/attr/attr-2.5.2.zip"
	attr_cmd_zip := exec.Command("curl", "-L", attr_zip_url, "-o", "package.zip")
	err = attr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	attr_bin_url := "https://download.savannah.nongnu.org/releases/attr/attr-2.5.2.bin"
	attr_cmd_bin := exec.Command("curl", "-L", attr_bin_url, "-o", "binary.bin")
	err = attr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	attr_src_url := "https://download.savannah.nongnu.org/releases/attr/attr-2.5.2.src.tar.gz"
	attr_cmd_src := exec.Command("curl", "-L", attr_src_url, "-o", "source.tar.gz")
	err = attr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	attr_cmd_direct := exec.Command("./binary")
	err = attr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
