package main

// Highs - Linear optimization software
// Homepage: https://www.maths.ed.ac.uk/hall/HiGHS/

import (
	"fmt"
	
	"os/exec"
)

func installHighs() {
	// Método 1: Descargar y extraer .tar.gz
	highs_tar_url := "https://github.com/ERGO-Code/HiGHS/archive/refs/tags/v1.7.2.tar.gz"
	highs_cmd_tar := exec.Command("curl", "-L", highs_tar_url, "-o", "package.tar.gz")
	err := highs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	highs_zip_url := "https://github.com/ERGO-Code/HiGHS/archive/refs/tags/v1.7.2.zip"
	highs_cmd_zip := exec.Command("curl", "-L", highs_zip_url, "-o", "package.zip")
	err = highs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	highs_bin_url := "https://github.com/ERGO-Code/HiGHS/archive/refs/tags/v1.7.2.bin"
	highs_cmd_bin := exec.Command("curl", "-L", highs_bin_url, "-o", "binary.bin")
	err = highs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	highs_src_url := "https://github.com/ERGO-Code/HiGHS/archive/refs/tags/v1.7.2.src.tar.gz"
	highs_cmd_src := exec.Command("curl", "-L", highs_src_url, "-o", "source.tar.gz")
	err = highs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	highs_cmd_direct := exec.Command("./binary")
	err = highs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
