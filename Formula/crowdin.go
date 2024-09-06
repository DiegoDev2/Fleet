package main

// Crowdin - Command-line tool that allows to manage your resources with crowdin.com
// Homepage: https://support.crowdin.com/cli-tool/

import (
	"fmt"
	
	"os/exec"
)

func installCrowdin() {
	// Método 1: Descargar y extraer .tar.gz
	crowdin_tar_url := "https://github.com/crowdin/crowdin-cli/releases/download/4.1.1/crowdin-cli.zip"
	crowdin_cmd_tar := exec.Command("curl", "-L", crowdin_tar_url, "-o", "package.tar.gz")
	err := crowdin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crowdin_zip_url := "https://github.com/crowdin/crowdin-cli/releases/download/4.1.1/crowdin-cli.zip"
	crowdin_cmd_zip := exec.Command("curl", "-L", crowdin_zip_url, "-o", "package.zip")
	err = crowdin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crowdin_bin_url := "https://github.com/crowdin/crowdin-cli/releases/download/4.1.1/crowdin-cli.zip"
	crowdin_cmd_bin := exec.Command("curl", "-L", crowdin_bin_url, "-o", "binary.bin")
	err = crowdin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crowdin_src_url := "https://github.com/crowdin/crowdin-cli/releases/download/4.1.1/crowdin-cli.zip"
	crowdin_cmd_src := exec.Command("curl", "-L", crowdin_src_url, "-o", "source.tar.gz")
	err = crowdin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crowdin_cmd_direct := exec.Command("./binary")
	err = crowdin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
