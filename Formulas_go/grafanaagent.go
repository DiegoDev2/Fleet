package main

// GrafanaAgent - Exporter for Prometheus Metrics, Loki Logs, and Tempo Traces
// Homepage: https://grafana.com/docs/agent/

import (
	"fmt"
	
	"os/exec"
)

func installGrafanaAgent() {
	// Método 1: Descargar y extraer .tar.gz
	grafanaagent_tar_url := "https://github.com/grafana/agent/archive/refs/tags/v0.42.0.tar.gz"
	grafanaagent_cmd_tar := exec.Command("curl", "-L", grafanaagent_tar_url, "-o", "package.tar.gz")
	err := grafanaagent_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grafanaagent_zip_url := "https://github.com/grafana/agent/archive/refs/tags/v0.42.0.zip"
	grafanaagent_cmd_zip := exec.Command("curl", "-L", grafanaagent_zip_url, "-o", "package.zip")
	err = grafanaagent_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grafanaagent_bin_url := "https://github.com/grafana/agent/archive/refs/tags/v0.42.0.bin"
	grafanaagent_cmd_bin := exec.Command("curl", "-L", grafanaagent_bin_url, "-o", "binary.bin")
	err = grafanaagent_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grafanaagent_src_url := "https://github.com/grafana/agent/archive/refs/tags/v0.42.0.src.tar.gz"
	grafanaagent_cmd_src := exec.Command("curl", "-L", grafanaagent_src_url, "-o", "source.tar.gz")
	err = grafanaagent_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grafanaagent_cmd_direct := exec.Command("./binary")
	err = grafanaagent_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go@1.22")
exec.Command("latte", "install", "go@1.22")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: yarn")
exec.Command("latte", "install", "yarn")
	fmt.Println("Instalando dependencia: systemd")
exec.Command("latte", "install", "systemd")
}
