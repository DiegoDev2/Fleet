package main

// Doppler - CLI for interacting with Doppler secrets and configuration
// Homepage: https://docs.doppler.com/docs

import (
	"fmt"
	
	"os/exec"
)

func installDoppler() {
	// Método 1: Descargar y extraer .tar.gz
	doppler_tar_url := "https://github.com/DopplerHQ/cli/archive/refs/tags/3.69.0.tar.gz"
	doppler_cmd_tar := exec.Command("curl", "-L", doppler_tar_url, "-o", "package.tar.gz")
	err := doppler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	doppler_zip_url := "https://github.com/DopplerHQ/cli/archive/refs/tags/3.69.0.zip"
	doppler_cmd_zip := exec.Command("curl", "-L", doppler_zip_url, "-o", "package.zip")
	err = doppler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	doppler_bin_url := "https://github.com/DopplerHQ/cli/archive/refs/tags/3.69.0.bin"
	doppler_cmd_bin := exec.Command("curl", "-L", doppler_bin_url, "-o", "binary.bin")
	err = doppler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	doppler_src_url := "https://github.com/DopplerHQ/cli/archive/refs/tags/3.69.0.src.tar.gz"
	doppler_cmd_src := exec.Command("curl", "-L", doppler_src_url, "-o", "source.tar.gz")
	err = doppler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	doppler_cmd_direct := exec.Command("./binary")
	err = doppler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
