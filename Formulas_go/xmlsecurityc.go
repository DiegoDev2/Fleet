package main

// XmlSecurityC - Implementation of primary security standards for XML
// Homepage: https://santuario.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installXmlSecurityC() {
	// Método 1: Descargar y extraer .tar.gz
	xmlsecurityc_tar_url := "https://www.apache.org/dyn/closer.lua?path=santuario/c-library/xml-security-c-2.0.4.tar.bz2"
	xmlsecurityc_cmd_tar := exec.Command("curl", "-L", xmlsecurityc_tar_url, "-o", "package.tar.gz")
	err := xmlsecurityc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmlsecurityc_zip_url := "https://www.apache.org/dyn/closer.lua?path=santuario/c-library/xml-security-c-2.0.4.tar.bz2"
	xmlsecurityc_cmd_zip := exec.Command("curl", "-L", xmlsecurityc_zip_url, "-o", "package.zip")
	err = xmlsecurityc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmlsecurityc_bin_url := "https://www.apache.org/dyn/closer.lua?path=santuario/c-library/xml-security-c-2.0.4.tar.bz2"
	xmlsecurityc_cmd_bin := exec.Command("curl", "-L", xmlsecurityc_bin_url, "-o", "binary.bin")
	err = xmlsecurityc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmlsecurityc_src_url := "https://www.apache.org/dyn/closer.lua?path=santuario/c-library/xml-security-c-2.0.4.tar.bz2"
	xmlsecurityc_cmd_src := exec.Command("curl", "-L", xmlsecurityc_src_url, "-o", "source.tar.gz")
	err = xmlsecurityc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmlsecurityc_cmd_direct := exec.Command("./binary")
	err = xmlsecurityc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
}
