package main

// Itpp - Library of math, signal, and communication classes and functions
// Homepage: https://itpp.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installItpp() {
	// Método 1: Descargar y extraer .tar.gz
	itpp_tar_url := "https://downloads.sourceforge.net/project/itpp/itpp/4.3.1/itpp-4.3.1.tar.bz2"
	itpp_cmd_tar := exec.Command("curl", "-L", itpp_tar_url, "-o", "package.tar.gz")
	err := itpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	itpp_zip_url := "https://downloads.sourceforge.net/project/itpp/itpp/4.3.1/itpp-4.3.1.tar.bz2"
	itpp_cmd_zip := exec.Command("curl", "-L", itpp_zip_url, "-o", "package.zip")
	err = itpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	itpp_bin_url := "https://downloads.sourceforge.net/project/itpp/itpp/4.3.1/itpp-4.3.1.tar.bz2"
	itpp_cmd_bin := exec.Command("curl", "-L", itpp_bin_url, "-o", "binary.bin")
	err = itpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	itpp_src_url := "https://downloads.sourceforge.net/project/itpp/itpp/4.3.1/itpp-4.3.1.tar.bz2"
	itpp_cmd_src := exec.Command("curl", "-L", itpp_src_url, "-o", "source.tar.gz")
	err = itpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	itpp_cmd_direct := exec.Command("./binary")
	err = itpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
}
