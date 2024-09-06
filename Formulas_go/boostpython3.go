package main

// BoostPython3 - C++ library for C++/Python3 interoperability
// Homepage: https://www.boost.org/

import (
	"fmt"
	
	"os/exec"
)

func installBoostPython3() {
	// Método 1: Descargar y extraer .tar.gz
	boostpython3_tar_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boostpython3_cmd_tar := exec.Command("curl", "-L", boostpython3_tar_url, "-o", "package.tar.gz")
	err := boostpython3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	boostpython3_zip_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boostpython3_cmd_zip := exec.Command("curl", "-L", boostpython3_zip_url, "-o", "package.zip")
	err = boostpython3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	boostpython3_bin_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boostpython3_cmd_bin := exec.Command("curl", "-L", boostpython3_bin_url, "-o", "binary.bin")
	err = boostpython3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	boostpython3_src_url := "https://github.com/boostorg/boost/releases/download/boost-1.86.0/boost-1.86.0-b2-nodocs.tar.xz"
	boostpython3_cmd_src := exec.Command("curl", "-L", boostpython3_src_url, "-o", "source.tar.gz")
	err = boostpython3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	boostpython3_cmd_direct := exec.Command("./binary")
	err = boostpython3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
