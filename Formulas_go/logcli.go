package main

// Logcli - Run LogQL queries against a Loki server
// Homepage: https://grafana.com/loki

import (
	"fmt"
	
	"os/exec"
)

func installLogcli() {
	// Método 1: Descargar y extraer .tar.gz
	logcli_tar_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.tar.gz"
	logcli_cmd_tar := exec.Command("curl", "-L", logcli_tar_url, "-o", "package.tar.gz")
	err := logcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	logcli_zip_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.zip"
	logcli_cmd_zip := exec.Command("curl", "-L", logcli_zip_url, "-o", "package.zip")
	err = logcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	logcli_bin_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.bin"
	logcli_cmd_bin := exec.Command("curl", "-L", logcli_bin_url, "-o", "binary.bin")
	err = logcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	logcli_src_url := "https://github.com/grafana/loki/archive/refs/tags/v3.1.1.src.tar.gz"
	logcli_cmd_src := exec.Command("curl", "-L", logcli_src_url, "-o", "source.tar.gz")
	err = logcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	logcli_cmd_direct := exec.Command("./binary")
	err = logcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: loki")
exec.Command("latte", "install", "loki")
}
