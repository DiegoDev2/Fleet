package main

// Libtextcat - N-gram-based text categorization library
// Homepage: https://software.wise-guys.nl/libtextcat/

import (
	"fmt"
	
	"os/exec"
)

func installLibtextcat() {
	// Método 1: Descargar y extraer .tar.gz
	libtextcat_tar_url := "https://software.wise-guys.nl/download/libtextcat-2.2.tar.gz"
	libtextcat_cmd_tar := exec.Command("curl", "-L", libtextcat_tar_url, "-o", "package.tar.gz")
	err := libtextcat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtextcat_zip_url := "https://software.wise-guys.nl/download/libtextcat-2.2.zip"
	libtextcat_cmd_zip := exec.Command("curl", "-L", libtextcat_zip_url, "-o", "package.zip")
	err = libtextcat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtextcat_bin_url := "https://software.wise-guys.nl/download/libtextcat-2.2.bin"
	libtextcat_cmd_bin := exec.Command("curl", "-L", libtextcat_bin_url, "-o", "binary.bin")
	err = libtextcat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtextcat_src_url := "https://software.wise-guys.nl/download/libtextcat-2.2.src.tar.gz"
	libtextcat_cmd_src := exec.Command("curl", "-L", libtextcat_src_url, "-o", "source.tar.gz")
	err = libtextcat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtextcat_cmd_direct := exec.Command("./binary")
	err = libtextcat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
