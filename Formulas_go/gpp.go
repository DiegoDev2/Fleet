package main

// Gpp - General-purpose preprocessor with customizable syntax
// Homepage: https://logological.org/gpp

import (
	"fmt"
	
	"os/exec"
)

func installGpp() {
	// Método 1: Descargar y extraer .tar.gz
	gpp_tar_url := "https://files.nothingisreal.com/software/gpp/gpp-2.28.tar.bz2"
	gpp_cmd_tar := exec.Command("curl", "-L", gpp_tar_url, "-o", "package.tar.gz")
	err := gpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpp_zip_url := "https://files.nothingisreal.com/software/gpp/gpp-2.28.tar.bz2"
	gpp_cmd_zip := exec.Command("curl", "-L", gpp_zip_url, "-o", "package.zip")
	err = gpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpp_bin_url := "https://files.nothingisreal.com/software/gpp/gpp-2.28.tar.bz2"
	gpp_cmd_bin := exec.Command("curl", "-L", gpp_bin_url, "-o", "binary.bin")
	err = gpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpp_src_url := "https://files.nothingisreal.com/software/gpp/gpp-2.28.tar.bz2"
	gpp_cmd_src := exec.Command("curl", "-L", gpp_src_url, "-o", "source.tar.gz")
	err = gpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpp_cmd_direct := exec.Command("./binary")
	err = gpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
}
