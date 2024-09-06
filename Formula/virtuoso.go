package main

// Virtuoso - High-performance object-relational SQL database
// Homepage: https://virtuoso.openlinksw.com/wiki/main/

import (
	"fmt"
	
	"os/exec"
)

func installVirtuoso() {
	// Método 1: Descargar y extraer .tar.gz
	virtuoso_tar_url := "https://github.com/openlink/virtuoso-opensource/releases/download/v7.2.13/virtuoso-opensource-7.2.13.tar.gz"
	virtuoso_cmd_tar := exec.Command("curl", "-L", virtuoso_tar_url, "-o", "package.tar.gz")
	err := virtuoso_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	virtuoso_zip_url := "https://github.com/openlink/virtuoso-opensource/releases/download/v7.2.13/virtuoso-opensource-7.2.13.zip"
	virtuoso_cmd_zip := exec.Command("curl", "-L", virtuoso_zip_url, "-o", "package.zip")
	err = virtuoso_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	virtuoso_bin_url := "https://github.com/openlink/virtuoso-opensource/releases/download/v7.2.13/virtuoso-opensource-7.2.13.bin"
	virtuoso_cmd_bin := exec.Command("curl", "-L", virtuoso_bin_url, "-o", "binary.bin")
	err = virtuoso_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	virtuoso_src_url := "https://github.com/openlink/virtuoso-opensource/releases/download/v7.2.13/virtuoso-opensource-7.2.13.src.tar.gz"
	virtuoso_cmd_src := exec.Command("curl", "-L", virtuoso_src_url, "-o", "source.tar.gz")
	err = virtuoso_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	virtuoso_cmd_direct := exec.Command("./binary")
	err = virtuoso_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gawk")
	exec.Command("latte", "install", "gawk").Run()
	fmt.Println("Instalando dependencia: openssl@3.0")
	exec.Command("latte", "install", "openssl@3.0").Run()
	fmt.Println("Instalando dependencia: net-tools")
	exec.Command("latte", "install", "net-tools").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
}
