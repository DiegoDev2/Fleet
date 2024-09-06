package main

// Loki - Horizontally-scalable, highly-available log aggregation system
// Homepage: https://grafana.com/loki

import (
	"fmt"
	
	"os/exec"
)

func installLoki() {
	// Método 1: Descargar y extraer .tar.gz
	loki_tar_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.tar.gz"
	loki_cmd_tar := exec.Command("curl", "-L", loki_tar_url, "-o", "package.tar.gz")
	err := loki_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	loki_zip_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.zip"
	loki_cmd_zip := exec.Command("curl", "-L", loki_zip_url, "-o", "package.zip")
	err = loki_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	loki_bin_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.bin"
	loki_cmd_bin := exec.Command("curl", "-L", loki_bin_url, "-o", "binary.bin")
	err = loki_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	loki_src_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.src.tar.gz"
	loki_cmd_src := exec.Command("curl", "-L", loki_src_url, "-o", "source.tar.gz")
	err = loki_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	loki_cmd_direct := exec.Command("./binary")
	err = loki_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
