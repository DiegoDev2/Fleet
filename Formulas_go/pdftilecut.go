package main

// Pdftilecut - Sub-divide a PDF page(s) into smaller pages so you can print them
// Homepage: https://github.com/oxplot/pdftilecut

import (
	"fmt"
	
	"os/exec"
)

func installPdftilecut() {
	// Método 1: Descargar y extraer .tar.gz
	pdftilecut_tar_url := "https://github.com/oxplot/pdftilecut/archive/refs/tags/v0.6.tar.gz"
	pdftilecut_cmd_tar := exec.Command("curl", "-L", pdftilecut_tar_url, "-o", "package.tar.gz")
	err := pdftilecut_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdftilecut_zip_url := "https://github.com/oxplot/pdftilecut/archive/refs/tags/v0.6.zip"
	pdftilecut_cmd_zip := exec.Command("curl", "-L", pdftilecut_zip_url, "-o", "package.zip")
	err = pdftilecut_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdftilecut_bin_url := "https://github.com/oxplot/pdftilecut/archive/refs/tags/v0.6.bin"
	pdftilecut_cmd_bin := exec.Command("curl", "-L", pdftilecut_bin_url, "-o", "binary.bin")
	err = pdftilecut_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdftilecut_src_url := "https://github.com/oxplot/pdftilecut/archive/refs/tags/v0.6.src.tar.gz"
	pdftilecut_cmd_src := exec.Command("curl", "-L", pdftilecut_src_url, "-o", "source.tar.gz")
	err = pdftilecut_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdftilecut_cmd_direct := exec.Command("./binary")
	err = pdftilecut_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: qpdf")
exec.Command("latte", "install", "qpdf")
}
