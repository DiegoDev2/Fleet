package main

// Alerter - Send User Alert Notification on macOS from the command-line
// Homepage: https://github.com/vjeantet/alerter

import (
	"fmt"
	
	"os/exec"
)

func installAlerter() {
	// Método 1: Descargar y extraer .tar.gz
	alerter_tar_url := "https://github.com/vjeantet/alerter/archive/refs/tags/004.tar.gz"
	alerter_cmd_tar := exec.Command("curl", "-L", alerter_tar_url, "-o", "package.tar.gz")
	err := alerter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alerter_zip_url := "https://github.com/vjeantet/alerter/archive/refs/tags/004.zip"
	alerter_cmd_zip := exec.Command("curl", "-L", alerter_zip_url, "-o", "package.zip")
	err = alerter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alerter_bin_url := "https://github.com/vjeantet/alerter/archive/refs/tags/004.bin"
	alerter_cmd_bin := exec.Command("curl", "-L", alerter_bin_url, "-o", "binary.bin")
	err = alerter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alerter_src_url := "https://github.com/vjeantet/alerter/archive/refs/tags/004.src.tar.gz"
	alerter_cmd_src := exec.Command("curl", "-L", alerter_src_url, "-o", "source.tar.gz")
	err = alerter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alerter_cmd_direct := exec.Command("./binary")
	err = alerter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
