package main

// Gismo - C++ library for isogeometric analysis (IGA)
// Homepage: https://gismo.github.io

import (
	"fmt"
	
	"os/exec"
)

func installGismo() {
	// Método 1: Descargar y extraer .tar.gz
	gismo_tar_url := "https://github.com/gismo/gismo/archive/refs/tags/v24.08.0.tar.gz"
	gismo_cmd_tar := exec.Command("curl", "-L", gismo_tar_url, "-o", "package.tar.gz")
	err := gismo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gismo_zip_url := "https://github.com/gismo/gismo/archive/refs/tags/v24.08.0.zip"
	gismo_cmd_zip := exec.Command("curl", "-L", gismo_zip_url, "-o", "package.zip")
	err = gismo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gismo_bin_url := "https://github.com/gismo/gismo/archive/refs/tags/v24.08.0.bin"
	gismo_cmd_bin := exec.Command("curl", "-L", gismo_bin_url, "-o", "binary.bin")
	err = gismo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gismo_src_url := "https://github.com/gismo/gismo/archive/refs/tags/v24.08.0.src.tar.gz"
	gismo_cmd_src := exec.Command("curl", "-L", gismo_src_url, "-o", "source.tar.gz")
	err = gismo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gismo_cmd_direct := exec.Command("./binary")
	err = gismo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: suite-sparse")
	exec.Command("latte", "install", "suite-sparse").Run()
	fmt.Println("Instalando dependencia: superlu")
	exec.Command("latte", "install", "superlu").Run()
	fmt.Println("Instalando dependencia: libomp")
	exec.Command("latte", "install", "libomp").Run()
}
