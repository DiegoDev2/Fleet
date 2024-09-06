package main

// Mad - MPEG audio decoder
// Homepage: https://www.underbit.com/products/mad/

import (
	"fmt"
	
	"os/exec"
)

func installMad() {
	// Método 1: Descargar y extraer .tar.gz
	mad_tar_url := "https://downloads.sourceforge.net/project/mad/libmad/0.15.1b/libmad-0.15.1b.tar.gz"
	mad_cmd_tar := exec.Command("curl", "-L", mad_tar_url, "-o", "package.tar.gz")
	err := mad_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mad_zip_url := "https://downloads.sourceforge.net/project/mad/libmad/0.15.1b/libmad-0.15.1b.zip"
	mad_cmd_zip := exec.Command("curl", "-L", mad_zip_url, "-o", "package.zip")
	err = mad_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mad_bin_url := "https://downloads.sourceforge.net/project/mad/libmad/0.15.1b/libmad-0.15.1b.bin"
	mad_cmd_bin := exec.Command("curl", "-L", mad_bin_url, "-o", "binary.bin")
	err = mad_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mad_src_url := "https://downloads.sourceforge.net/project/mad/libmad/0.15.1b/libmad-0.15.1b.src.tar.gz"
	mad_cmd_src := exec.Command("curl", "-L", mad_src_url, "-o", "source.tar.gz")
	err = mad_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mad_cmd_direct := exec.Command("./binary")
	err = mad_cmd_direct.Run()
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
