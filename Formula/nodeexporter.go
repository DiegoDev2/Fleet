package main

// NodeExporter - Prometheus exporter for machine metrics
// Homepage: https://prometheus.io/

import (
	"fmt"
	
	"os/exec"
)

func installNodeExporter() {
	// Método 1: Descargar y extraer .tar.gz
	nodeexporter_tar_url := "https://github.com/prometheus/node_exporter/archive/refs/tags/v1.8.2.tar.gz"
	nodeexporter_cmd_tar := exec.Command("curl", "-L", nodeexporter_tar_url, "-o", "package.tar.gz")
	err := nodeexporter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nodeexporter_zip_url := "https://github.com/prometheus/node_exporter/archive/refs/tags/v1.8.2.zip"
	nodeexporter_cmd_zip := exec.Command("curl", "-L", nodeexporter_zip_url, "-o", "package.zip")
	err = nodeexporter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nodeexporter_bin_url := "https://github.com/prometheus/node_exporter/archive/refs/tags/v1.8.2.bin"
	nodeexporter_cmd_bin := exec.Command("curl", "-L", nodeexporter_bin_url, "-o", "binary.bin")
	err = nodeexporter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nodeexporter_src_url := "https://github.com/prometheus/node_exporter/archive/refs/tags/v1.8.2.src.tar.gz"
	nodeexporter_cmd_src := exec.Command("curl", "-L", nodeexporter_src_url, "-o", "source.tar.gz")
	err = nodeexporter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nodeexporter_cmd_direct := exec.Command("./binary")
	err = nodeexporter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
