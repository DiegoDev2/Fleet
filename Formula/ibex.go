package main

// Ibex - C++ library for constraint processing over real numbers
// Homepage: https://github.com/ibex-team/ibex-lib

import (
	"fmt"
	
	"os/exec"
)

func installIbex() {
	// Método 1: Descargar y extraer .tar.gz
	ibex_tar_url := "https://github.com/ibex-team/ibex-lib/archive/refs/tags/ibex-2.8.9.tar.gz"
	ibex_cmd_tar := exec.Command("curl", "-L", ibex_tar_url, "-o", "package.tar.gz")
	err := ibex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ibex_zip_url := "https://github.com/ibex-team/ibex-lib/archive/refs/tags/ibex-2.8.9.zip"
	ibex_cmd_zip := exec.Command("curl", "-L", ibex_zip_url, "-o", "package.zip")
	err = ibex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ibex_bin_url := "https://github.com/ibex-team/ibex-lib/archive/refs/tags/ibex-2.8.9.bin"
	ibex_cmd_bin := exec.Command("curl", "-L", ibex_bin_url, "-o", "binary.bin")
	err = ibex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ibex_src_url := "https://github.com/ibex-team/ibex-lib/archive/refs/tags/ibex-2.8.9.src.tar.gz"
	ibex_cmd_src := exec.Command("curl", "-L", ibex_src_url, "-o", "source.tar.gz")
	err = ibex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ibex_cmd_direct := exec.Command("./binary")
	err = ibex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
