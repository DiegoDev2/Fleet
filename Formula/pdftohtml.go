package main

// Pdftohtml - Utility which converts PDF files into HTML and XML formats
// Homepage: https://pdftohtml.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPdftohtml() {
	// Método 1: Descargar y extraer .tar.gz
	pdftohtml_tar_url := "https://downloads.sourceforge.net/project/pdftohtml/Experimental%20Versions/pdftohtml%200.40/pdftohtml-0.40a.tar.gz"
	pdftohtml_cmd_tar := exec.Command("curl", "-L", pdftohtml_tar_url, "-o", "package.tar.gz")
	err := pdftohtml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdftohtml_zip_url := "https://downloads.sourceforge.net/project/pdftohtml/Experimental%20Versions/pdftohtml%200.40/pdftohtml-0.40a.zip"
	pdftohtml_cmd_zip := exec.Command("curl", "-L", pdftohtml_zip_url, "-o", "package.zip")
	err = pdftohtml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdftohtml_bin_url := "https://downloads.sourceforge.net/project/pdftohtml/Experimental%20Versions/pdftohtml%200.40/pdftohtml-0.40a.bin"
	pdftohtml_cmd_bin := exec.Command("curl", "-L", pdftohtml_bin_url, "-o", "binary.bin")
	err = pdftohtml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdftohtml_src_url := "https://downloads.sourceforge.net/project/pdftohtml/Experimental%20Versions/pdftohtml%200.40/pdftohtml-0.40a.src.tar.gz"
	pdftohtml_cmd_src := exec.Command("curl", "-L", pdftohtml_src_url, "-o", "source.tar.gz")
	err = pdftohtml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdftohtml_cmd_direct := exec.Command("./binary")
	err = pdftohtml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
