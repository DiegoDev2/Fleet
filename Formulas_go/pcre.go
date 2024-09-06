package main

// Pcre - Perl compatible regular expressions library
// Homepage: https://www.pcre.org/

import (
	"fmt"
	
	"os/exec"
)

func installPcre() {
	// Método 1: Descargar y extraer .tar.gz
	pcre_tar_url := "https://downloads.sourceforge.net/project/pcre/pcre/8.45/pcre-8.45.tar.bz2"
	pcre_cmd_tar := exec.Command("curl", "-L", pcre_tar_url, "-o", "package.tar.gz")
	err := pcre_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcre_zip_url := "https://downloads.sourceforge.net/project/pcre/pcre/8.45/pcre-8.45.tar.bz2"
	pcre_cmd_zip := exec.Command("curl", "-L", pcre_zip_url, "-o", "package.zip")
	err = pcre_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcre_bin_url := "https://downloads.sourceforge.net/project/pcre/pcre/8.45/pcre-8.45.tar.bz2"
	pcre_cmd_bin := exec.Command("curl", "-L", pcre_bin_url, "-o", "binary.bin")
	err = pcre_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcre_src_url := "https://downloads.sourceforge.net/project/pcre/pcre/8.45/pcre-8.45.tar.bz2"
	pcre_cmd_src := exec.Command("curl", "-L", pcre_src_url, "-o", "source.tar.gz")
	err = pcre_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcre_cmd_direct := exec.Command("./binary")
	err = pcre_cmd_direct.Run()
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
}
