package main

// Opensaml - Library for Security Assertion Markup Language
// Homepage: https://wiki.shibboleth.net/confluence/display/OpenSAML/Home

import (
	"fmt"
	
	"os/exec"
)

func installOpensaml() {
	// Método 1: Descargar y extraer .tar.gz
	opensaml_tar_url := "https://shibboleth.net/downloads/c++-opensaml/3.2.1/opensaml-3.2.1.tar.bz2"
	opensaml_cmd_tar := exec.Command("curl", "-L", opensaml_tar_url, "-o", "package.tar.gz")
	err := opensaml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opensaml_zip_url := "https://shibboleth.net/downloads/c++-opensaml/3.2.1/opensaml-3.2.1.tar.bz2"
	opensaml_cmd_zip := exec.Command("curl", "-L", opensaml_zip_url, "-o", "package.zip")
	err = opensaml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opensaml_bin_url := "https://shibboleth.net/downloads/c++-opensaml/3.2.1/opensaml-3.2.1.tar.bz2"
	opensaml_cmd_bin := exec.Command("curl", "-L", opensaml_bin_url, "-o", "binary.bin")
	err = opensaml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opensaml_src_url := "https://shibboleth.net/downloads/c++-opensaml/3.2.1/opensaml-3.2.1.tar.bz2"
	opensaml_cmd_src := exec.Command("curl", "-L", opensaml_src_url, "-o", "source.tar.gz")
	err = opensaml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opensaml_cmd_direct := exec.Command("./binary")
	err = opensaml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: log4shib")
exec.Command("latte", "install", "log4shib")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
	fmt.Println("Instalando dependencia: xml-security-c")
exec.Command("latte", "install", "xml-security-c")
	fmt.Println("Instalando dependencia: xml-tooling-c")
exec.Command("latte", "install", "xml-tooling-c")
}
