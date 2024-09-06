package main

// Docx2txt - Converts Microsoft Office docx documents to equivalent text documents
// Homepage: https://docx2txt.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDocx2txt() {
	// Método 1: Descargar y extraer .tar.gz
	docx2txt_tar_url := "https://downloads.sourceforge.net/project/docx2txt/docx2txt/v1.4/docx2txt-1.4.tgz"
	docx2txt_cmd_tar := exec.Command("curl", "-L", docx2txt_tar_url, "-o", "package.tar.gz")
	err := docx2txt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	docx2txt_zip_url := "https://downloads.sourceforge.net/project/docx2txt/docx2txt/v1.4/docx2txt-1.4.tgz"
	docx2txt_cmd_zip := exec.Command("curl", "-L", docx2txt_zip_url, "-o", "package.zip")
	err = docx2txt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	docx2txt_bin_url := "https://downloads.sourceforge.net/project/docx2txt/docx2txt/v1.4/docx2txt-1.4.tgz"
	docx2txt_cmd_bin := exec.Command("curl", "-L", docx2txt_bin_url, "-o", "binary.bin")
	err = docx2txt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	docx2txt_src_url := "https://downloads.sourceforge.net/project/docx2txt/docx2txt/v1.4/docx2txt-1.4.tgz"
	docx2txt_cmd_src := exec.Command("curl", "-L", docx2txt_src_url, "-o", "source.tar.gz")
	err = docx2txt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	docx2txt_cmd_direct := exec.Command("./binary")
	err = docx2txt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
