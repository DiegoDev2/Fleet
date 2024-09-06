package main

// BasisUniversal - Basis Universal GPU texture codec command-line compression tool
// Homepage: https://github.com/BinomialLLC/basis_universal

import (
	"fmt"
	
	"os/exec"
)

func installBasisUniversal() {
	// Método 1: Descargar y extraer .tar.gz
	basisuniversal_tar_url := "https://github.com/BinomialLLC/basis_universal/archive/refs/tags/1.16.4.tar.gz"
	basisuniversal_cmd_tar := exec.Command("curl", "-L", basisuniversal_tar_url, "-o", "package.tar.gz")
	err := basisuniversal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	basisuniversal_zip_url := "https://github.com/BinomialLLC/basis_universal/archive/refs/tags/1.16.4.zip"
	basisuniversal_cmd_zip := exec.Command("curl", "-L", basisuniversal_zip_url, "-o", "package.zip")
	err = basisuniversal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	basisuniversal_bin_url := "https://github.com/BinomialLLC/basis_universal/archive/refs/tags/1.16.4.bin"
	basisuniversal_cmd_bin := exec.Command("curl", "-L", basisuniversal_bin_url, "-o", "binary.bin")
	err = basisuniversal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	basisuniversal_src_url := "https://github.com/BinomialLLC/basis_universal/archive/refs/tags/1.16.4.src.tar.gz"
	basisuniversal_cmd_src := exec.Command("curl", "-L", basisuniversal_src_url, "-o", "source.tar.gz")
	err = basisuniversal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	basisuniversal_cmd_direct := exec.Command("./binary")
	err = basisuniversal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
