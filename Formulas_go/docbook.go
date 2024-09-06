package main

// Docbook - Standard XML representation system for technical documents
// Homepage: https://docbook.org/

import (
	"fmt"
	
	"os/exec"
)

func installDocbook() {
	// Método 1: Descargar y extraer .tar.gz
	docbook_tar_url := "https://docbook.org/xml/5.1/docbook-v5.1-os.zip"
	docbook_cmd_tar := exec.Command("curl", "-L", docbook_tar_url, "-o", "package.tar.gz")
	err := docbook_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	docbook_zip_url := "https://docbook.org/xml/5.1/docbook-v5.1-os.zip"
	docbook_cmd_zip := exec.Command("curl", "-L", docbook_zip_url, "-o", "package.zip")
	err = docbook_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	docbook_bin_url := "https://docbook.org/xml/5.1/docbook-v5.1-os.zip"
	docbook_cmd_bin := exec.Command("curl", "-L", docbook_bin_url, "-o", "binary.bin")
	err = docbook_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	docbook_src_url := "https://docbook.org/xml/5.1/docbook-v5.1-os.zip"
	docbook_cmd_src := exec.Command("curl", "-L", docbook_src_url, "-o", "source.tar.gz")
	err = docbook_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	docbook_cmd_direct := exec.Command("./binary")
	err = docbook_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
