package main

// DocbookXsl - XML vocabulary to create presentation-neutral documents
// Homepage: https://github.com/docbook/xslt10-stylesheets

import (
	"fmt"
	
	"os/exec"
)

func installDocbookXsl() {
	// Método 1: Descargar y extraer .tar.gz
	docbookxsl_tar_url := "https://github.com/docbook/xslt10-stylesheets/releases/download/release%2F1.79.2/docbook-xsl-nons-1.79.2.tar.bz2"
	docbookxsl_cmd_tar := exec.Command("curl", "-L", docbookxsl_tar_url, "-o", "package.tar.gz")
	err := docbookxsl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	docbookxsl_zip_url := "https://github.com/docbook/xslt10-stylesheets/releases/download/release%2F1.79.2/docbook-xsl-nons-1.79.2.tar.bz2"
	docbookxsl_cmd_zip := exec.Command("curl", "-L", docbookxsl_zip_url, "-o", "package.zip")
	err = docbookxsl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	docbookxsl_bin_url := "https://github.com/docbook/xslt10-stylesheets/releases/download/release%2F1.79.2/docbook-xsl-nons-1.79.2.tar.bz2"
	docbookxsl_cmd_bin := exec.Command("curl", "-L", docbookxsl_bin_url, "-o", "binary.bin")
	err = docbookxsl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	docbookxsl_src_url := "https://github.com/docbook/xslt10-stylesheets/releases/download/release%2F1.79.2/docbook-xsl-nons-1.79.2.tar.bz2"
	docbookxsl_cmd_src := exec.Command("curl", "-L", docbookxsl_src_url, "-o", "source.tar.gz")
	err = docbookxsl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	docbookxsl_cmd_direct := exec.Command("./binary")
	err = docbookxsl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook")
exec.Command("latte", "install", "docbook")
}
