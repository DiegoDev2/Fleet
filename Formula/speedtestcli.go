package main

// SpeedtestCli - Command-line interface for https://speedtest.net bandwidth tests
// Homepage: https://github.com/sivel/speedtest-cli

import (
	"fmt"
	
	"os/exec"
)

func installSpeedtestCli() {
	// Método 1: Descargar y extraer .tar.gz
	speedtestcli_tar_url := "https://github.com/sivel/speedtest-cli/archive/refs/tags/v2.1.3.tar.gz"
	speedtestcli_cmd_tar := exec.Command("curl", "-L", speedtestcli_tar_url, "-o", "package.tar.gz")
	err := speedtestcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	speedtestcli_zip_url := "https://github.com/sivel/speedtest-cli/archive/refs/tags/v2.1.3.zip"
	speedtestcli_cmd_zip := exec.Command("curl", "-L", speedtestcli_zip_url, "-o", "package.zip")
	err = speedtestcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	speedtestcli_bin_url := "https://github.com/sivel/speedtest-cli/archive/refs/tags/v2.1.3.bin"
	speedtestcli_cmd_bin := exec.Command("curl", "-L", speedtestcli_bin_url, "-o", "binary.bin")
	err = speedtestcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	speedtestcli_src_url := "https://github.com/sivel/speedtest-cli/archive/refs/tags/v2.1.3.src.tar.gz"
	speedtestcli_cmd_src := exec.Command("curl", "-L", speedtestcli_src_url, "-o", "source.tar.gz")
	err = speedtestcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	speedtestcli_cmd_direct := exec.Command("./binary")
	err = speedtestcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
