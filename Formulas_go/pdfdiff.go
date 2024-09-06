package main

// PdfDiff - Tool for visualizing differences between two pdf files
// Homepage: https://github.com/serhack/pdf-diff

import (
	"fmt"
	
	"os/exec"
)

func installPdfDiff() {
	// Método 1: Descargar y extraer .tar.gz
	pdfdiff_tar_url := "https://github.com/serhack/pdf-diff/archive/refs/tags/v0.0.1.tar.gz"
	pdfdiff_cmd_tar := exec.Command("curl", "-L", pdfdiff_tar_url, "-o", "package.tar.gz")
	err := pdfdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdfdiff_zip_url := "https://github.com/serhack/pdf-diff/archive/refs/tags/v0.0.1.zip"
	pdfdiff_cmd_zip := exec.Command("curl", "-L", pdfdiff_zip_url, "-o", "package.zip")
	err = pdfdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdfdiff_bin_url := "https://github.com/serhack/pdf-diff/archive/refs/tags/v0.0.1.bin"
	pdfdiff_cmd_bin := exec.Command("curl", "-L", pdfdiff_bin_url, "-o", "binary.bin")
	err = pdfdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdfdiff_src_url := "https://github.com/serhack/pdf-diff/archive/refs/tags/v0.0.1.src.tar.gz"
	pdfdiff_cmd_src := exec.Command("curl", "-L", pdfdiff_src_url, "-o", "source.tar.gz")
	err = pdfdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdfdiff_cmd_direct := exec.Command("./binary")
	err = pdfdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
}
