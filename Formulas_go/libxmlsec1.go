package main

// Libxmlsec1 - XML security library
// Homepage: https://www.aleksey.com/xmlsec/

import (
	"fmt"
	
	"os/exec"
)

func installLibxmlsec1() {
	// Método 1: Descargar y extraer .tar.gz
	libxmlsec1_tar_url := "https://www.aleksey.com/xmlsec/download/xmlsec1-1.3.5.tar.gz"
	libxmlsec1_cmd_tar := exec.Command("curl", "-L", libxmlsec1_tar_url, "-o", "package.tar.gz")
	err := libxmlsec1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxmlsec1_zip_url := "https://www.aleksey.com/xmlsec/download/xmlsec1-1.3.5.zip"
	libxmlsec1_cmd_zip := exec.Command("curl", "-L", libxmlsec1_zip_url, "-o", "package.zip")
	err = libxmlsec1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxmlsec1_bin_url := "https://www.aleksey.com/xmlsec/download/xmlsec1-1.3.5.bin"
	libxmlsec1_cmd_bin := exec.Command("curl", "-L", libxmlsec1_bin_url, "-o", "binary.bin")
	err = libxmlsec1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxmlsec1_src_url := "https://www.aleksey.com/xmlsec/download/xmlsec1-1.3.5.src.tar.gz"
	libxmlsec1_cmd_src := exec.Command("curl", "-L", libxmlsec1_src_url, "-o", "source.tar.gz")
	err = libxmlsec1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxmlsec1_cmd_direct := exec.Command("./binary")
	err = libxmlsec1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: libxml2")
exec.Command("latte", "install", "libxml2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
