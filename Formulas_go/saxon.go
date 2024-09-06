package main

// Saxon - XSLT and XQuery processor
// Homepage: https://github.com/Saxonica/Saxon-HE

import (
	"fmt"
	
	"os/exec"
)

func installSaxon() {
	// Método 1: Descargar y extraer .tar.gz
	saxon_tar_url := "https://github.com/Saxonica/Saxon-HE/releases/download/SaxonHE12-5/SaxonHE12-5J.zip"
	saxon_cmd_tar := exec.Command("curl", "-L", saxon_tar_url, "-o", "package.tar.gz")
	err := saxon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	saxon_zip_url := "https://github.com/Saxonica/Saxon-HE/releases/download/SaxonHE12-5/SaxonHE12-5J.zip"
	saxon_cmd_zip := exec.Command("curl", "-L", saxon_zip_url, "-o", "package.zip")
	err = saxon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	saxon_bin_url := "https://github.com/Saxonica/Saxon-HE/releases/download/SaxonHE12-5/SaxonHE12-5J.zip"
	saxon_cmd_bin := exec.Command("curl", "-L", saxon_bin_url, "-o", "binary.bin")
	err = saxon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	saxon_src_url := "https://github.com/Saxonica/Saxon-HE/releases/download/SaxonHE12-5/SaxonHE12-5J.zip"
	saxon_cmd_src := exec.Command("curl", "-L", saxon_src_url, "-o", "source.tar.gz")
	err = saxon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	saxon_cmd_direct := exec.Command("./binary")
	err = saxon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
