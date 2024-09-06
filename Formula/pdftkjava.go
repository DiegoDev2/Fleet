package main

// PdftkJava - Port of pdftk in java
// Homepage: https://gitlab.com/pdftk-java/pdftk

import (
	"fmt"
	
	"os/exec"
)

func installPdftkJava() {
	// Método 1: Descargar y extraer .tar.gz
	pdftkjava_tar_url := "https://gitlab.com/pdftk-java/pdftk/-/archive/v3.3.3/pdftk-v3.3.3.tar.gz"
	pdftkjava_cmd_tar := exec.Command("curl", "-L", pdftkjava_tar_url, "-o", "package.tar.gz")
	err := pdftkjava_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdftkjava_zip_url := "https://gitlab.com/pdftk-java/pdftk/-/archive/v3.3.3/pdftk-v3.3.3.zip"
	pdftkjava_cmd_zip := exec.Command("curl", "-L", pdftkjava_zip_url, "-o", "package.zip")
	err = pdftkjava_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdftkjava_bin_url := "https://gitlab.com/pdftk-java/pdftk/-/archive/v3.3.3/pdftk-v3.3.3.bin"
	pdftkjava_cmd_bin := exec.Command("curl", "-L", pdftkjava_bin_url, "-o", "binary.bin")
	err = pdftkjava_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdftkjava_src_url := "https://gitlab.com/pdftk-java/pdftk/-/archive/v3.3.3/pdftk-v3.3.3.src.tar.gz"
	pdftkjava_cmd_src := exec.Command("curl", "-L", pdftkjava_src_url, "-o", "source.tar.gz")
	err = pdftkjava_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdftkjava_cmd_direct := exec.Command("./binary")
	err = pdftkjava_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
	exec.Command("latte", "install", "gradle").Run()
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
