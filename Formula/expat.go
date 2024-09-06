package main

// Expat - XML 1.0 parser
// Homepage: https://libexpat.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installExpat() {
	// Método 1: Descargar y extraer .tar.gz
	expat_tar_url := "https://github.com/libexpat/libexpat/releases/download/R_2_6_3/expat-2.6.3.tar.lz"
	expat_cmd_tar := exec.Command("curl", "-L", expat_tar_url, "-o", "package.tar.gz")
	err := expat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	expat_zip_url := "https://github.com/libexpat/libexpat/releases/download/R_2_6_3/expat-2.6.3.tar.lz"
	expat_cmd_zip := exec.Command("curl", "-L", expat_zip_url, "-o", "package.zip")
	err = expat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	expat_bin_url := "https://github.com/libexpat/libexpat/releases/download/R_2_6_3/expat-2.6.3.tar.lz"
	expat_cmd_bin := exec.Command("curl", "-L", expat_bin_url, "-o", "binary.bin")
	err = expat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	expat_src_url := "https://github.com/libexpat/libexpat/releases/download/R_2_6_3/expat-2.6.3.tar.lz"
	expat_cmd_src := exec.Command("curl", "-L", expat_src_url, "-o", "source.tar.gz")
	err = expat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	expat_cmd_direct := exec.Command("./binary")
	err = expat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: docbook2x")
	exec.Command("latte", "install", "docbook2x").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
