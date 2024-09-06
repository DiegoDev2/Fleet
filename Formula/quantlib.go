package main

// Quantlib - Library for quantitative finance
// Homepage: https://www.quantlib.org/

import (
	"fmt"
	
	"os/exec"
)

func installQuantlib() {
	// Método 1: Descargar y extraer .tar.gz
	quantlib_tar_url := "https://github.com/lballabio/QuantLib/releases/download/v1.35/QuantLib-1.35.tar.gz"
	quantlib_cmd_tar := exec.Command("curl", "-L", quantlib_tar_url, "-o", "package.tar.gz")
	err := quantlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	quantlib_zip_url := "https://github.com/lballabio/QuantLib/releases/download/v1.35/QuantLib-1.35.zip"
	quantlib_cmd_zip := exec.Command("curl", "-L", quantlib_zip_url, "-o", "package.zip")
	err = quantlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	quantlib_bin_url := "https://github.com/lballabio/QuantLib/releases/download/v1.35/QuantLib-1.35.bin"
	quantlib_cmd_bin := exec.Command("curl", "-L", quantlib_bin_url, "-o", "binary.bin")
	err = quantlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	quantlib_src_url := "https://github.com/lballabio/QuantLib/releases/download/v1.35/QuantLib-1.35.src.tar.gz"
	quantlib_cmd_src := exec.Command("curl", "-L", quantlib_src_url, "-o", "source.tar.gz")
	err = quantlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	quantlib_cmd_direct := exec.Command("./binary")
	err = quantlib_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
}
