package main

// ShibbolethSp - Shibboleth 2 Service Provider daemon
// Homepage: https://wiki.shibboleth.net/confluence/display/SHIB2

import (
	"fmt"
	
	"os/exec"
)

func installShibbolethSp() {
	// Método 1: Descargar y extraer .tar.gz
	shibbolethsp_tar_url := "https://shibboleth.net/downloads/service-provider/3.4.1/shibboleth-sp-3.4.1.tar.bz2"
	shibbolethsp_cmd_tar := exec.Command("curl", "-L", shibbolethsp_tar_url, "-o", "package.tar.gz")
	err := shibbolethsp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shibbolethsp_zip_url := "https://shibboleth.net/downloads/service-provider/3.4.1/shibboleth-sp-3.4.1.tar.bz2"
	shibbolethsp_cmd_zip := exec.Command("curl", "-L", shibbolethsp_zip_url, "-o", "package.zip")
	err = shibbolethsp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shibbolethsp_bin_url := "https://shibboleth.net/downloads/service-provider/3.4.1/shibboleth-sp-3.4.1.tar.bz2"
	shibbolethsp_cmd_bin := exec.Command("curl", "-L", shibbolethsp_bin_url, "-o", "binary.bin")
	err = shibbolethsp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shibbolethsp_src_url := "https://shibboleth.net/downloads/service-provider/3.4.1/shibboleth-sp-3.4.1.tar.bz2"
	shibbolethsp_cmd_src := exec.Command("curl", "-L", shibbolethsp_src_url, "-o", "source.tar.gz")
	err = shibbolethsp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shibbolethsp_cmd_direct := exec.Command("./binary")
	err = shibbolethsp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: apr")
exec.Command("latte", "install", "apr")
	fmt.Println("Instalando dependencia: apr-util")
exec.Command("latte", "install", "apr-util")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: httpd")
exec.Command("latte", "install", "httpd")
	fmt.Println("Instalando dependencia: log4shib")
exec.Command("latte", "install", "log4shib")
	fmt.Println("Instalando dependencia: opensaml")
exec.Command("latte", "install", "opensaml")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: unixodbc")
exec.Command("latte", "install", "unixodbc")
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
	fmt.Println("Instalando dependencia: xml-security-c")
exec.Command("latte", "install", "xml-security-c")
	fmt.Println("Instalando dependencia: xml-tooling-c")
exec.Command("latte", "install", "xml-tooling-c")
}
