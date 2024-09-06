package main

// InfluxdbAT1 - Time series, events, and metrics database
// Homepage: https://influxdata.com/time-series-platform/influxdb/

import (
	"fmt"
	
	"os/exec"
)

func installInfluxdbAT1() {
	// Método 1: Descargar y extraer .tar.gz
	influxdbat1_tar_url := "https://github.com/influxdata/influxdb/archive/refs/tags/v1.11.5.tar.gz"
	influxdbat1_cmd_tar := exec.Command("curl", "-L", influxdbat1_tar_url, "-o", "package.tar.gz")
	err := influxdbat1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	influxdbat1_zip_url := "https://github.com/influxdata/influxdb/archive/refs/tags/v1.11.5.zip"
	influxdbat1_cmd_zip := exec.Command("curl", "-L", influxdbat1_zip_url, "-o", "package.zip")
	err = influxdbat1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	influxdbat1_bin_url := "https://github.com/influxdata/influxdb/archive/refs/tags/v1.11.5.bin"
	influxdbat1_cmd_bin := exec.Command("curl", "-L", influxdbat1_bin_url, "-o", "binary.bin")
	err = influxdbat1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	influxdbat1_src_url := "https://github.com/influxdata/influxdb/archive/refs/tags/v1.11.5.src.tar.gz"
	influxdbat1_cmd_src := exec.Command("curl", "-L", influxdbat1_src_url, "-o", "source.tar.gz")
	err = influxdbat1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	influxdbat1_cmd_direct := exec.Command("./binary")
	err = influxdbat1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
