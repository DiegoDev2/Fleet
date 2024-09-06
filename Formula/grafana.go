package main

// Grafana - Gorgeous metric visualizations and dashboards for timeseries databases
// Homepage: https://grafana.com

import (
	"fmt"
	
	"os/exec"
)

func installGrafana() {
	// Método 1: Descargar y extraer .tar.gz
	grafana_tar_url := "https://github.com/grafana/grafana/archive/refs/tags/v11.2.0.tar.gz"
	grafana_cmd_tar := exec.Command("curl", "-L", grafana_tar_url, "-o", "package.tar.gz")
	err := grafana_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grafana_zip_url := "https://github.com/grafana/grafana/archive/refs/tags/v11.2.0.zip"
	grafana_cmd_zip := exec.Command("curl", "-L", grafana_zip_url, "-o", "package.zip")
	err = grafana_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grafana_bin_url := "https://github.com/grafana/grafana/archive/refs/tags/v11.2.0.bin"
	grafana_cmd_bin := exec.Command("curl", "-L", grafana_bin_url, "-o", "binary.bin")
	err = grafana_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grafana_src_url := "https://github.com/grafana/grafana/archive/refs/tags/v11.2.0.src.tar.gz"
	grafana_cmd_src := exec.Command("curl", "-L", grafana_src_url, "-o", "source.tar.gz")
	err = grafana_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grafana_cmd_direct := exec.Command("./binary")
	err = grafana_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: corepack")
	exec.Command("latte", "install", "corepack").Run()
	fmt.Println("Instalando dependencia: go@1.22")
	exec.Command("latte", "install", "go@1.22").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
}
