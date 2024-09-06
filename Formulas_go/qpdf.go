package main

// Qpdf - Tools for and transforming and inspecting PDF files
// Homepage: https://github.com/qpdf/qpdf

import (
	"fmt"
	
	"os/exec"
)

func installQpdf() {
	// Método 1: Descargar y extraer .tar.gz
	qpdf_tar_url := "https://github.com/qpdf/qpdf/releases/download/v11.9.1/qpdf-11.9.1.tar.gz"
	qpdf_cmd_tar := exec.Command("curl", "-L", qpdf_tar_url, "-o", "package.tar.gz")
	err := qpdf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qpdf_zip_url := "https://github.com/qpdf/qpdf/releases/download/v11.9.1/qpdf-11.9.1.zip"
	qpdf_cmd_zip := exec.Command("curl", "-L", qpdf_zip_url, "-o", "package.zip")
	err = qpdf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qpdf_bin_url := "https://github.com/qpdf/qpdf/releases/download/v11.9.1/qpdf-11.9.1.bin"
	qpdf_cmd_bin := exec.Command("curl", "-L", qpdf_bin_url, "-o", "binary.bin")
	err = qpdf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qpdf_src_url := "https://github.com/qpdf/qpdf/releases/download/v11.9.1/qpdf-11.9.1.src.tar.gz"
	qpdf_cmd_src := exec.Command("curl", "-L", qpdf_src_url, "-o", "source.tar.gz")
	err = qpdf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qpdf_cmd_direct := exec.Command("./binary")
	err = qpdf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
