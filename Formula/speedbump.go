package main

// Speedbump - TCP proxy for simulating variable, yet predictable network latency
// Homepage: https://github.com/kffl/speedbump

import (
	"fmt"
	
	"os/exec"
)

func installSpeedbump() {
	// Método 1: Descargar y extraer .tar.gz
	speedbump_tar_url := "https://github.com/kffl/speedbump/archive/refs/tags/v1.1.0.tar.gz"
	speedbump_cmd_tar := exec.Command("curl", "-L", speedbump_tar_url, "-o", "package.tar.gz")
	err := speedbump_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	speedbump_zip_url := "https://github.com/kffl/speedbump/archive/refs/tags/v1.1.0.zip"
	speedbump_cmd_zip := exec.Command("curl", "-L", speedbump_zip_url, "-o", "package.zip")
	err = speedbump_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	speedbump_bin_url := "https://github.com/kffl/speedbump/archive/refs/tags/v1.1.0.bin"
	speedbump_cmd_bin := exec.Command("curl", "-L", speedbump_bin_url, "-o", "binary.bin")
	err = speedbump_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	speedbump_src_url := "https://github.com/kffl/speedbump/archive/refs/tags/v1.1.0.src.tar.gz"
	speedbump_cmd_src := exec.Command("curl", "-L", speedbump_src_url, "-o", "source.tar.gz")
	err = speedbump_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	speedbump_cmd_direct := exec.Command("./binary")
	err = speedbump_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
