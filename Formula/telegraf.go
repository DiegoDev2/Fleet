package main

// Telegraf - Plugin-driven server agent for collecting & reporting metrics
// Homepage: https://www.influxdata.com/time-series-platform/telegraf/

import (
	"fmt"
	
	"os/exec"
)

func installTelegraf() {
	// Método 1: Descargar y extraer .tar.gz
	telegraf_tar_url := "https://github.com/influxdata/telegraf/archive/refs/tags/v1.31.3.tar.gz"
	telegraf_cmd_tar := exec.Command("curl", "-L", telegraf_tar_url, "-o", "package.tar.gz")
	err := telegraf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	telegraf_zip_url := "https://github.com/influxdata/telegraf/archive/refs/tags/v1.31.3.zip"
	telegraf_cmd_zip := exec.Command("curl", "-L", telegraf_zip_url, "-o", "package.zip")
	err = telegraf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	telegraf_bin_url := "https://github.com/influxdata/telegraf/archive/refs/tags/v1.31.3.bin"
	telegraf_cmd_bin := exec.Command("curl", "-L", telegraf_bin_url, "-o", "binary.bin")
	err = telegraf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	telegraf_src_url := "https://github.com/influxdata/telegraf/archive/refs/tags/v1.31.3.src.tar.gz"
	telegraf_cmd_src := exec.Command("curl", "-L", telegraf_src_url, "-o", "source.tar.gz")
	err = telegraf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	telegraf_cmd_direct := exec.Command("./binary")
	err = telegraf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
