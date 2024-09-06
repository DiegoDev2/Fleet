package main

// Xmlsectool - Check schema validity and signature of an XML document
// Homepage: https://wiki.shibboleth.net/confluence/display/XSTJ3/xmlsectool+V3+Home

import (
	"fmt"
	
	"os/exec"
)

func installXmlsectool() {
	// Método 1: Descargar y extraer .tar.gz
	xmlsectool_tar_url := "https://shibboleth.net/downloads/tools/xmlsectool/3.0.0/xmlsectool-3.0.0-bin.zip"
	xmlsectool_cmd_tar := exec.Command("curl", "-L", xmlsectool_tar_url, "-o", "package.tar.gz")
	err := xmlsectool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmlsectool_zip_url := "https://shibboleth.net/downloads/tools/xmlsectool/3.0.0/xmlsectool-3.0.0-bin.zip"
	xmlsectool_cmd_zip := exec.Command("curl", "-L", xmlsectool_zip_url, "-o", "package.zip")
	err = xmlsectool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmlsectool_bin_url := "https://shibboleth.net/downloads/tools/xmlsectool/3.0.0/xmlsectool-3.0.0-bin.zip"
	xmlsectool_cmd_bin := exec.Command("curl", "-L", xmlsectool_bin_url, "-o", "binary.bin")
	err = xmlsectool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmlsectool_src_url := "https://shibboleth.net/downloads/tools/xmlsectool/3.0.0/xmlsectool-3.0.0-bin.zip"
	xmlsectool_cmd_src := exec.Command("curl", "-L", xmlsectool_src_url, "-o", "source.tar.gz")
	err = xmlsectool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmlsectool_cmd_direct := exec.Command("./binary")
	err = xmlsectool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
