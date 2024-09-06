package main

// Influxdb - Time series, events, and metrics database
// Homepage: https://influxdata.com/time-series-platform/influxdb/

import (
	"fmt"
	
	"os/exec"
)

func installInfluxdb() {
	// Método 1: Descargar y extraer .tar.gz
	influxdb_tar_url := "https://github.com/influxdata/influxdb.git"
	influxdb_cmd_tar := exec.Command("curl", "-L", influxdb_tar_url, "-o", "package.tar.gz")
	err := influxdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	influxdb_zip_url := "https://github.com/influxdata/influxdb.git"
	influxdb_cmd_zip := exec.Command("curl", "-L", influxdb_zip_url, "-o", "package.zip")
	err = influxdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	influxdb_bin_url := "https://github.com/influxdata/influxdb.git"
	influxdb_cmd_bin := exec.Command("curl", "-L", influxdb_bin_url, "-o", "binary.bin")
	err = influxdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	influxdb_src_url := "https://github.com/influxdata/influxdb.git"
	influxdb_cmd_src := exec.Command("curl", "-L", influxdb_src_url, "-o", "source.tar.gz")
	err = influxdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	influxdb_cmd_direct := exec.Command("./binary")
	err = influxdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: breezy")
	exec.Command("latte", "install", "breezy").Run()
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
