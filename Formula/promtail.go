package main

// Promtail - Log agent for Loki
// Homepage: https://grafana.com/loki

import (
	"fmt"
	
	"os/exec"
)

func installPromtail() {
	// Método 1: Descargar y extraer .tar.gz
	promtail_tar_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.tar.gz"
	promtail_cmd_tar := exec.Command("curl", "-L", promtail_tar_url, "-o", "package.tar.gz")
	err := promtail_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	promtail_zip_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.zip"
	promtail_cmd_zip := exec.Command("curl", "-L", promtail_zip_url, "-o", "package.zip")
	err = promtail_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	promtail_bin_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.bin"
	promtail_cmd_bin := exec.Command("curl", "-L", promtail_bin_url, "-o", "binary.bin")
	err = promtail_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	promtail_src_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.src.tar.gz"
	promtail_cmd_src := exec.Command("curl", "-L", promtail_src_url, "-o", "source.tar.gz")
	err = promtail_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	promtail_cmd_direct := exec.Command("./binary")
	err = promtail_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: systemd")
	exec.Command("latte", "install", "systemd").Run()
}
