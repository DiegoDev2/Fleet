package main

// Systemc - Core SystemC language and examples
// Homepage: https://accellera.org/

import (
	"fmt"
	
	"os/exec"
)

func installSystemc() {
	// Método 1: Descargar y extraer .tar.gz
	systemc_tar_url := "https://github.com/accellera-official/systemc/archive/refs/tags/3.0.0.tar.gz"
	systemc_cmd_tar := exec.Command("curl", "-L", systemc_tar_url, "-o", "package.tar.gz")
	err := systemc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	systemc_zip_url := "https://github.com/accellera-official/systemc/archive/refs/tags/3.0.0.zip"
	systemc_cmd_zip := exec.Command("curl", "-L", systemc_zip_url, "-o", "package.zip")
	err = systemc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	systemc_bin_url := "https://github.com/accellera-official/systemc/archive/refs/tags/3.0.0.bin"
	systemc_cmd_bin := exec.Command("curl", "-L", systemc_bin_url, "-o", "binary.bin")
	err = systemc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	systemc_src_url := "https://github.com/accellera-official/systemc/archive/refs/tags/3.0.0.src.tar.gz"
	systemc_cmd_src := exec.Command("curl", "-L", systemc_src_url, "-o", "source.tar.gz")
	err = systemc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	systemc_cmd_direct := exec.Command("./binary")
	err = systemc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
