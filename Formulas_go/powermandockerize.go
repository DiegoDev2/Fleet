package main

// PowermanDockerize - Utility to simplify running applications in docker containers
// Homepage: https://github.com/powerman/dockerize

import (
	"fmt"
	
	"os/exec"
)

func installPowermanDockerize() {
	// Método 1: Descargar y extraer .tar.gz
	powermandockerize_tar_url := "https://github.com/powerman/dockerize/archive/refs/tags/v0.20.0.tar.gz"
	powermandockerize_cmd_tar := exec.Command("curl", "-L", powermandockerize_tar_url, "-o", "package.tar.gz")
	err := powermandockerize_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	powermandockerize_zip_url := "https://github.com/powerman/dockerize/archive/refs/tags/v0.20.0.zip"
	powermandockerize_cmd_zip := exec.Command("curl", "-L", powermandockerize_zip_url, "-o", "package.zip")
	err = powermandockerize_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	powermandockerize_bin_url := "https://github.com/powerman/dockerize/archive/refs/tags/v0.20.0.bin"
	powermandockerize_cmd_bin := exec.Command("curl", "-L", powermandockerize_bin_url, "-o", "binary.bin")
	err = powermandockerize_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	powermandockerize_src_url := "https://github.com/powerman/dockerize/archive/refs/tags/v0.20.0.src.tar.gz"
	powermandockerize_cmd_src := exec.Command("curl", "-L", powermandockerize_src_url, "-o", "source.tar.gz")
	err = powermandockerize_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	powermandockerize_cmd_direct := exec.Command("./binary")
	err = powermandockerize_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
