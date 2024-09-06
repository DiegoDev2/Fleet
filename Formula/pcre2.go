package main

// Pcre2 - Perl compatible regular expressions library with a new API
// Homepage: https://www.pcre.org/

import (
	"fmt"
	
	"os/exec"
)

func installPcre2() {
	// Método 1: Descargar y extraer .tar.gz
	pcre2_tar_url := "https://github.com/PCRE2Project/pcre2/releases/download/pcre2-10.44/pcre2-10.44.tar.bz2"
	pcre2_cmd_tar := exec.Command("curl", "-L", pcre2_tar_url, "-o", "package.tar.gz")
	err := pcre2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pcre2_zip_url := "https://github.com/PCRE2Project/pcre2/releases/download/pcre2-10.44/pcre2-10.44.tar.bz2"
	pcre2_cmd_zip := exec.Command("curl", "-L", pcre2_zip_url, "-o", "package.zip")
	err = pcre2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pcre2_bin_url := "https://github.com/PCRE2Project/pcre2/releases/download/pcre2-10.44/pcre2-10.44.tar.bz2"
	pcre2_cmd_bin := exec.Command("curl", "-L", pcre2_bin_url, "-o", "binary.bin")
	err = pcre2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pcre2_src_url := "https://github.com/PCRE2Project/pcre2/releases/download/pcre2-10.44/pcre2-10.44.tar.bz2"
	pcre2_cmd_src := exec.Command("curl", "-L", pcre2_src_url, "-o", "source.tar.gz")
	err = pcre2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pcre2_cmd_direct := exec.Command("./binary")
	err = pcre2_cmd_direct.Run()
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
