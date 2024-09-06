package main

// Tfel - Code generation tool dedicated to material knowledge for numerical mechanics
// Homepage: https://thelfer.github.io/tfel/web/index.html

import (
	"fmt"
	
	"os/exec"
)

func installTfel() {
	// Método 1: Descargar y extraer .tar.gz
	tfel_tar_url := "https://github.com/thelfer/tfel/archive/refs/tags/TFEL-4.2.1.tar.gz"
	tfel_cmd_tar := exec.Command("curl", "-L", tfel_tar_url, "-o", "package.tar.gz")
	err := tfel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tfel_zip_url := "https://github.com/thelfer/tfel/archive/refs/tags/TFEL-4.2.1.zip"
	tfel_cmd_zip := exec.Command("curl", "-L", tfel_zip_url, "-o", "package.zip")
	err = tfel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tfel_bin_url := "https://github.com/thelfer/tfel/archive/refs/tags/TFEL-4.2.1.bin"
	tfel_cmd_bin := exec.Command("curl", "-L", tfel_bin_url, "-o", "binary.bin")
	err = tfel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tfel_src_url := "https://github.com/thelfer/tfel/archive/refs/tags/TFEL-4.2.1.src.tar.gz"
	tfel_cmd_src := exec.Command("curl", "-L", tfel_src_url, "-o", "source.tar.gz")
	err = tfel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tfel_cmd_direct := exec.Command("./binary")
	err = tfel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: boost-python3")
	exec.Command("latte", "install", "boost-python3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
