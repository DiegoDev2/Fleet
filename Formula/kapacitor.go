package main

// Kapacitor - Open source time series data processor
// Homepage: https://github.com/influxdata/kapacitor

import (
	"fmt"
	
	"os/exec"
)

func installKapacitor() {
	// Método 1: Descargar y extraer .tar.gz
	kapacitor_tar_url := "https://github.com/influxdata/kapacitor.git"
	kapacitor_cmd_tar := exec.Command("curl", "-L", kapacitor_tar_url, "-o", "package.tar.gz")
	err := kapacitor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kapacitor_zip_url := "https://github.com/influxdata/kapacitor.git"
	kapacitor_cmd_zip := exec.Command("curl", "-L", kapacitor_zip_url, "-o", "package.zip")
	err = kapacitor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kapacitor_bin_url := "https://github.com/influxdata/kapacitor.git"
	kapacitor_cmd_bin := exec.Command("curl", "-L", kapacitor_bin_url, "-o", "binary.bin")
	err = kapacitor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kapacitor_src_url := "https://github.com/influxdata/kapacitor.git"
	kapacitor_cmd_src := exec.Command("curl", "-L", kapacitor_src_url, "-o", "source.tar.gz")
	err = kapacitor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kapacitor_cmd_direct := exec.Command("./binary")
	err = kapacitor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
