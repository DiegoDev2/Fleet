package main

// Chronograf - Open source monitoring and visualization UI for the TICK stack
// Homepage: https://docs.influxdata.com/chronograf/latest/

import (
	"fmt"
	
	"os/exec"
)

func installChronograf() {
	// Método 1: Descargar y extraer .tar.gz
	chronograf_tar_url := "https://github.com/influxdata/chronograf/archive/refs/tags/1.10.5.tar.gz"
	chronograf_cmd_tar := exec.Command("curl", "-L", chronograf_tar_url, "-o", "package.tar.gz")
	err := chronograf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chronograf_zip_url := "https://github.com/influxdata/chronograf/archive/refs/tags/1.10.5.zip"
	chronograf_cmd_zip := exec.Command("curl", "-L", chronograf_zip_url, "-o", "package.zip")
	err = chronograf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chronograf_bin_url := "https://github.com/influxdata/chronograf/archive/refs/tags/1.10.5.bin"
	chronograf_cmd_bin := exec.Command("curl", "-L", chronograf_bin_url, "-o", "binary.bin")
	err = chronograf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chronograf_src_url := "https://github.com/influxdata/chronograf/archive/refs/tags/1.10.5.src.tar.gz"
	chronograf_cmd_src := exec.Command("curl", "-L", chronograf_src_url, "-o", "source.tar.gz")
	err = chronograf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chronograf_cmd_direct := exec.Command("./binary")
	err = chronograf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: yarn")
	exec.Command("latte", "install", "yarn").Run()
	fmt.Println("Instalando dependencia: influxdb")
	exec.Command("latte", "install", "influxdb").Run()
	fmt.Println("Instalando dependencia: kapacitor")
	exec.Command("latte", "install", "kapacitor").Run()
}
