package main

// Burst - Radix sort, lazy ranges and iterators, and more. Boost-like header-only library
// Homepage: https://github.com/izvolov/burst

import (
	"fmt"
	
	"os/exec"
)

func installBurst() {
	// Método 1: Descargar y extraer .tar.gz
	burst_tar_url := "https://github.com/izvolov/burst/archive/refs/tags/v3.1.1.tar.gz"
	burst_cmd_tar := exec.Command("curl", "-L", burst_tar_url, "-o", "package.tar.gz")
	err := burst_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	burst_zip_url := "https://github.com/izvolov/burst/archive/refs/tags/v3.1.1.zip"
	burst_cmd_zip := exec.Command("curl", "-L", burst_zip_url, "-o", "package.zip")
	err = burst_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	burst_bin_url := "https://github.com/izvolov/burst/archive/refs/tags/v3.1.1.bin"
	burst_cmd_bin := exec.Command("curl", "-L", burst_bin_url, "-o", "binary.bin")
	err = burst_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	burst_src_url := "https://github.com/izvolov/burst/archive/refs/tags/v3.1.1.src.tar.gz"
	burst_cmd_src := exec.Command("curl", "-L", burst_src_url, "-o", "source.tar.gz")
	err = burst_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	burst_cmd_direct := exec.Command("./binary")
	err = burst_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
}
