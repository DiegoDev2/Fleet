package main

// Gsoap - SOAP stub and skeleton compiler for C and C++
// Homepage: https://www.genivia.com/products.html

import (
	"fmt"
	
	"os/exec"
)

func installGsoap() {
	// Método 1: Descargar y extraer .tar.gz
	gsoap_tar_url := "https://downloads.sourceforge.net/project/gsoap2/gsoap_2.8.135.zip"
	gsoap_cmd_tar := exec.Command("curl", "-L", gsoap_tar_url, "-o", "package.tar.gz")
	err := gsoap_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gsoap_zip_url := "https://downloads.sourceforge.net/project/gsoap2/gsoap_2.8.135.zip"
	gsoap_cmd_zip := exec.Command("curl", "-L", gsoap_zip_url, "-o", "package.zip")
	err = gsoap_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gsoap_bin_url := "https://downloads.sourceforge.net/project/gsoap2/gsoap_2.8.135.zip"
	gsoap_cmd_bin := exec.Command("curl", "-L", gsoap_bin_url, "-o", "binary.bin")
	err = gsoap_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gsoap_src_url := "https://downloads.sourceforge.net/project/gsoap2/gsoap_2.8.135.zip"
	gsoap_cmd_src := exec.Command("curl", "-L", gsoap_src_url, "-o", "source.tar.gz")
	err = gsoap_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gsoap_cmd_direct := exec.Command("./binary")
	err = gsoap_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
