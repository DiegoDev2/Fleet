package main

// BoostBcp - Utility for extracting subsets of the Boost library
// Homepage: https://www.boost.org/doc/tools/bcp/

import (
	"fmt"
	
	"os/exec"
)

func installBoostBcp() {
	// Método 1: Descargar y extraer .tar.gz
	boostbcp_tar_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boostbcp_cmd_tar := exec.Command("curl", "-L", boostbcp_tar_url, "-o", "package.tar.gz")
	err := boostbcp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boostbcp_zip_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boostbcp_cmd_zip := exec.Command("curl", "-L", boostbcp_zip_url, "-o", "package.zip")
	err = boostbcp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boostbcp_bin_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boostbcp_cmd_bin := exec.Command("curl", "-L", boostbcp_bin_url, "-o", "binary.bin")
	err = boostbcp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boostbcp_src_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boostbcp_cmd_src := exec.Command("curl", "-L", boostbcp_src_url, "-o", "source.tar.gz")
	err = boostbcp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boostbcp_cmd_direct := exec.Command("./binary")
	err = boostbcp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: boost-build")
	exec.Command("latte", "install", "boost-build").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
}
