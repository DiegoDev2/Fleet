package main

// Gauge - Test automation tool that supports executable documentation
// Homepage: https://gauge.org

import (
	"fmt"
	
	"os/exec"
)

func installGauge() {
	// Método 1: Descargar y extraer .tar.gz
	gauge_tar_url := "https://github.com/getgauge/gauge/archive/refs/tags/v1.6.8.tar.gz"
	gauge_cmd_tar := exec.Command("curl", "-L", gauge_tar_url, "-o", "package.tar.gz")
	err := gauge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gauge_zip_url := "https://github.com/getgauge/gauge/archive/refs/tags/v1.6.8.zip"
	gauge_cmd_zip := exec.Command("curl", "-L", gauge_zip_url, "-o", "package.zip")
	err = gauge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gauge_bin_url := "https://github.com/getgauge/gauge/archive/refs/tags/v1.6.8.bin"
	gauge_cmd_bin := exec.Command("curl", "-L", gauge_bin_url, "-o", "binary.bin")
	err = gauge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gauge_src_url := "https://github.com/getgauge/gauge/archive/refs/tags/v1.6.8.src.tar.gz"
	gauge_cmd_src := exec.Command("curl", "-L", gauge_src_url, "-o", "source.tar.gz")
	err = gauge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gauge_cmd_direct := exec.Command("./binary")
	err = gauge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
