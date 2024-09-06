package main

// Victoriametrics - Cost-effective and scalable monitoring solution and time series database
// Homepage: https://victoriametrics.com/

import (
	"fmt"
	
	"os/exec"
)

func installVictoriametrics() {
	// Método 1: Descargar y extraer .tar.gz
	victoriametrics_tar_url := "https://github.com/VictoriaMetrics/VictoriaMetrics/archive/refs/tags/v1.103.0.tar.gz"
	victoriametrics_cmd_tar := exec.Command("curl", "-L", victoriametrics_tar_url, "-o", "package.tar.gz")
	err := victoriametrics_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	victoriametrics_zip_url := "https://github.com/VictoriaMetrics/VictoriaMetrics/archive/refs/tags/v1.103.0.zip"
	victoriametrics_cmd_zip := exec.Command("curl", "-L", victoriametrics_zip_url, "-o", "package.zip")
	err = victoriametrics_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	victoriametrics_bin_url := "https://github.com/VictoriaMetrics/VictoriaMetrics/archive/refs/tags/v1.103.0.bin"
	victoriametrics_cmd_bin := exec.Command("curl", "-L", victoriametrics_bin_url, "-o", "binary.bin")
	err = victoriametrics_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	victoriametrics_src_url := "https://github.com/VictoriaMetrics/VictoriaMetrics/archive/refs/tags/v1.103.0.src.tar.gz"
	victoriametrics_cmd_src := exec.Command("curl", "-L", victoriametrics_src_url, "-o", "source.tar.gz")
	err = victoriametrics_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	victoriametrics_cmd_direct := exec.Command("./binary")
	err = victoriametrics_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
