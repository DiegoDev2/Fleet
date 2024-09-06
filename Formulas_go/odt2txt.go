package main

// Odt2txt - Convert OpenDocument files to plain text
// Homepage: https://github.com/dstosberg/odt2txt/

import (
	"fmt"
	
	"os/exec"
)

func installOdt2txt() {
	// Método 1: Descargar y extraer .tar.gz
	odt2txt_tar_url := "https://github.com/dstosberg/odt2txt/archive/refs/tags/v0.5.tar.gz"
	odt2txt_cmd_tar := exec.Command("curl", "-L", odt2txt_tar_url, "-o", "package.tar.gz")
	err := odt2txt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	odt2txt_zip_url := "https://github.com/dstosberg/odt2txt/archive/refs/tags/v0.5.zip"
	odt2txt_cmd_zip := exec.Command("curl", "-L", odt2txt_zip_url, "-o", "package.zip")
	err = odt2txt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	odt2txt_bin_url := "https://github.com/dstosberg/odt2txt/archive/refs/tags/v0.5.bin"
	odt2txt_cmd_bin := exec.Command("curl", "-L", odt2txt_bin_url, "-o", "binary.bin")
	err = odt2txt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	odt2txt_src_url := "https://github.com/dstosberg/odt2txt/archive/refs/tags/v0.5.src.tar.gz"
	odt2txt_cmd_src := exec.Command("curl", "-L", odt2txt_src_url, "-o", "source.tar.gz")
	err = odt2txt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	odt2txt_cmd_direct := exec.Command("./binary")
	err = odt2txt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
