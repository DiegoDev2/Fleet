package main

// Djvu2pdf - Small tool to convert Djvu files to PDF files
// Homepage: https://0x2a.at/site/projects/djvu2pdf/

import (
	"fmt"
	
	"os/exec"
)

func installDjvu2pdf() {
	// Método 1: Descargar y extraer .tar.gz
	djvu2pdf_tar_url := "https://0x2a.at/site/projects/djvu2pdf/djvu2pdf-0.9.2.tar.gz"
	djvu2pdf_cmd_tar := exec.Command("curl", "-L", djvu2pdf_tar_url, "-o", "package.tar.gz")
	err := djvu2pdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	djvu2pdf_zip_url := "https://0x2a.at/site/projects/djvu2pdf/djvu2pdf-0.9.2.zip"
	djvu2pdf_cmd_zip := exec.Command("curl", "-L", djvu2pdf_zip_url, "-o", "package.zip")
	err = djvu2pdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	djvu2pdf_bin_url := "https://0x2a.at/site/projects/djvu2pdf/djvu2pdf-0.9.2.bin"
	djvu2pdf_cmd_bin := exec.Command("curl", "-L", djvu2pdf_bin_url, "-o", "binary.bin")
	err = djvu2pdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	djvu2pdf_src_url := "https://0x2a.at/site/projects/djvu2pdf/djvu2pdf-0.9.2.src.tar.gz"
	djvu2pdf_cmd_src := exec.Command("curl", "-L", djvu2pdf_src_url, "-o", "source.tar.gz")
	err = djvu2pdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	djvu2pdf_cmd_direct := exec.Command("./binary")
	err = djvu2pdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: djvulibre")
	exec.Command("latte", "install", "djvulibre").Run()
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
}
