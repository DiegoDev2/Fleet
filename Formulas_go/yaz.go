package main

// Yaz - Toolkit for Z39.50/SRW/SRU clients/servers
// Homepage: https://www.indexdata.com/resources/software/yaz/

import (
	"fmt"
	
	"os/exec"
)

func installYaz() {
	// Método 1: Descargar y extraer .tar.gz
	yaz_tar_url := "http://deb.debian.org/debian/pool/main/y/yaz/yaz_5.34.1.orig.tar.gz"
	yaz_cmd_tar := exec.Command("curl", "-L", yaz_tar_url, "-o", "package.tar.gz")
	err := yaz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yaz_zip_url := "http://deb.debian.org/debian/pool/main/y/yaz/yaz_5.34.1.orig.zip"
	yaz_cmd_zip := exec.Command("curl", "-L", yaz_zip_url, "-o", "package.zip")
	err = yaz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yaz_bin_url := "http://deb.debian.org/debian/pool/main/y/yaz/yaz_5.34.1.orig.bin"
	yaz_cmd_bin := exec.Command("curl", "-L", yaz_bin_url, "-o", "binary.bin")
	err = yaz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yaz_src_url := "http://deb.debian.org/debian/pool/main/y/yaz/yaz_5.34.1.orig.src.tar.gz"
	yaz_cmd_src := exec.Command("curl", "-L", yaz_src_url, "-o", "source.tar.gz")
	err = yaz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yaz_cmd_direct := exec.Command("./binary")
	err = yaz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: docbook-xsl")
exec.Command("latte", "install", "docbook-xsl")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
