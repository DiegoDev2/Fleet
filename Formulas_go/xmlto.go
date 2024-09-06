package main

// Xmlto - Convert XML to another format (based on XSL or other tools)
// Homepage: https://pagure.io/xmlto/

import (
	"fmt"
	
	"os/exec"
)

func installXmlto() {
	// Método 1: Descargar y extraer .tar.gz
	xmlto_tar_url := "https://releases.pagure.org/xmlto/xmlto-0.0.28.tar.bz2"
	xmlto_cmd_tar := exec.Command("curl", "-L", xmlto_tar_url, "-o", "package.tar.gz")
	err := xmlto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmlto_zip_url := "https://releases.pagure.org/xmlto/xmlto-0.0.28.tar.bz2"
	xmlto_cmd_zip := exec.Command("curl", "-L", xmlto_zip_url, "-o", "package.zip")
	err = xmlto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmlto_bin_url := "https://releases.pagure.org/xmlto/xmlto-0.0.28.tar.bz2"
	xmlto_cmd_bin := exec.Command("curl", "-L", xmlto_bin_url, "-o", "binary.bin")
	err = xmlto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmlto_src_url := "https://releases.pagure.org/xmlto/xmlto-0.0.28.tar.bz2"
	xmlto_cmd_src := exec.Command("curl", "-L", xmlto_src_url, "-o", "source.tar.gz")
	err = xmlto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmlto_cmd_direct := exec.Command("./binary")
	err = xmlto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: docbook")
exec.Command("latte", "install", "docbook")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: gnu-getopt")
exec.Command("latte", "install", "gnu-getopt")
}
