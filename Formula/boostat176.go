package main

// BoostAT176 - Collection of portable C++ source libraries
// Homepage: https://www.boost.org/

import (
	"fmt"
	
	"os/exec"
)

func installBoostAT176() {
	// Método 1: Descargar y extraer .tar.gz
	boostat176_tar_url := "https://boostorg.jfrog.io/artifactory/main/release/1.76.0/source/boost_1_76_0.tar.bz2"
	boostat176_cmd_tar := exec.Command("curl", "-L", boostat176_tar_url, "-o", "package.tar.gz")
	err := boostat176_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boostat176_zip_url := "https://boostorg.jfrog.io/artifactory/main/release/1.76.0/source/boost_1_76_0.tar.bz2"
	boostat176_cmd_zip := exec.Command("curl", "-L", boostat176_zip_url, "-o", "package.zip")
	err = boostat176_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boostat176_bin_url := "https://boostorg.jfrog.io/artifactory/main/release/1.76.0/source/boost_1_76_0.tar.bz2"
	boostat176_cmd_bin := exec.Command("curl", "-L", boostat176_bin_url, "-o", "binary.bin")
	err = boostat176_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boostat176_src_url := "https://boostorg.jfrog.io/artifactory/main/release/1.76.0/source/boost_1_76_0.tar.bz2"
	boostat176_cmd_src := exec.Command("curl", "-L", boostat176_src_url, "-o", "source.tar.gz")
	err = boostat176_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boostat176_cmd_direct := exec.Command("./binary")
	err = boostat176_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
}
