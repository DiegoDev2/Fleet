package main

// XmlToolingC - Provides a higher level interface to XML processing
// Homepage: https://wiki.shibboleth.net/confluence/display/OpenSAML/XMLTooling-C

import (
	"fmt"
	
	"os/exec"
)

func installXmlToolingC() {
	// Método 1: Descargar y extraer .tar.gz
	xmltoolingc_tar_url := "https://shibboleth.net/downloads/c++-opensaml/3.2.1/xmltooling-3.2.4.tar.bz2"
	xmltoolingc_cmd_tar := exec.Command("curl", "-L", xmltoolingc_tar_url, "-o", "package.tar.gz")
	err := xmltoolingc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmltoolingc_zip_url := "https://shibboleth.net/downloads/c++-opensaml/3.2.1/xmltooling-3.2.4.tar.bz2"
	xmltoolingc_cmd_zip := exec.Command("curl", "-L", xmltoolingc_zip_url, "-o", "package.zip")
	err = xmltoolingc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmltoolingc_bin_url := "https://shibboleth.net/downloads/c++-opensaml/3.2.1/xmltooling-3.2.4.tar.bz2"
	xmltoolingc_cmd_bin := exec.Command("curl", "-L", xmltoolingc_bin_url, "-o", "binary.bin")
	err = xmltoolingc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmltoolingc_src_url := "https://shibboleth.net/downloads/c++-opensaml/3.2.1/xmltooling-3.2.4.tar.bz2"
	xmltoolingc_cmd_src := exec.Command("curl", "-L", xmltoolingc_src_url, "-o", "source.tar.gz")
	err = xmltoolingc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmltoolingc_cmd_direct := exec.Command("./binary")
	err = xmltoolingc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: log4shib")
exec.Command("latte", "install", "log4shib")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
	fmt.Println("Instalando dependencia: xml-security-c")
exec.Command("latte", "install", "xml-security-c")
}
