package main

// Prometheus - Service monitoring system and time series database
// Homepage: https://prometheus.io/

import (
	"fmt"
	
	"os/exec"
)

func installPrometheus() {
	// Método 1: Descargar y extraer .tar.gz
	prometheus_tar_url := "https://github.com/prometheus/prometheus/archive/refs/tags/v2.54.1.tar.gz"
	prometheus_cmd_tar := exec.Command("curl", "-L", prometheus_tar_url, "-o", "package.tar.gz")
	err := prometheus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prometheus_zip_url := "https://github.com/prometheus/prometheus/archive/refs/tags/v2.54.1.zip"
	prometheus_cmd_zip := exec.Command("curl", "-L", prometheus_zip_url, "-o", "package.zip")
	err = prometheus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prometheus_bin_url := "https://github.com/prometheus/prometheus/archive/refs/tags/v2.54.1.bin"
	prometheus_cmd_bin := exec.Command("curl", "-L", prometheus_bin_url, "-o", "binary.bin")
	err = prometheus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prometheus_src_url := "https://github.com/prometheus/prometheus/archive/refs/tags/v2.54.1.src.tar.gz"
	prometheus_cmd_src := exec.Command("curl", "-L", prometheus_src_url, "-o", "source.tar.gz")
	err = prometheus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prometheus_cmd_direct := exec.Command("./binary")
	err = prometheus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-tar")
exec.Command("latte", "install", "gnu-tar")
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: yarn")
exec.Command("latte", "install", "yarn")
}
