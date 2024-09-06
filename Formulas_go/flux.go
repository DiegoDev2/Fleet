package main

// Flux - Lightweight scripting language for querying databases
// Homepage: https://www.influxdata.com/products/flux/

import (
	"fmt"
	
	"os/exec"
)

func installFlux() {
	// Método 1: Descargar y extraer .tar.gz
	flux_tar_url := "https://github.com/influxdata/flux.git"
	flux_cmd_tar := exec.Command("curl", "-L", flux_tar_url, "-o", "package.tar.gz")
	err := flux_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flux_zip_url := "https://github.com/influxdata/flux.git"
	flux_cmd_zip := exec.Command("curl", "-L", flux_zip_url, "-o", "package.zip")
	err = flux_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flux_bin_url := "https://github.com/influxdata/flux.git"
	flux_cmd_bin := exec.Command("curl", "-L", flux_bin_url, "-o", "binary.bin")
	err = flux_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flux_src_url := "https://github.com/influxdata/flux.git"
	flux_cmd_src := exec.Command("curl", "-L", flux_src_url, "-o", "source.tar.gz")
	err = flux_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flux_cmd_direct := exec.Command("./binary")
	err = flux_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
