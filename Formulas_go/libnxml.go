package main

// Libnxml - C library for parsing, writing, and creating XML files
// Homepage: https://github.com/bakulf/libnxml

import (
	"fmt"
	
	"os/exec"
)

func installLibnxml() {
	// Método 1: Descargar y extraer .tar.gz
	libnxml_tar_url := "https://github.com/bakulf/libnxml/archive/refs/tags/0.18.5.tar.gz"
	libnxml_cmd_tar := exec.Command("curl", "-L", libnxml_tar_url, "-o", "package.tar.gz")
	err := libnxml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libnxml_zip_url := "https://github.com/bakulf/libnxml/archive/refs/tags/0.18.5.zip"
	libnxml_cmd_zip := exec.Command("curl", "-L", libnxml_zip_url, "-o", "package.zip")
	err = libnxml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libnxml_bin_url := "https://github.com/bakulf/libnxml/archive/refs/tags/0.18.5.bin"
	libnxml_cmd_bin := exec.Command("curl", "-L", libnxml_bin_url, "-o", "binary.bin")
	err = libnxml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libnxml_src_url := "https://github.com/bakulf/libnxml/archive/refs/tags/0.18.5.src.tar.gz"
	libnxml_cmd_src := exec.Command("curl", "-L", libnxml_src_url, "-o", "source.tar.gz")
	err = libnxml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libnxml_cmd_direct := exec.Command("./binary")
	err = libnxml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
