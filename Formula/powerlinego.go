package main

// PowerlineGo - Beautiful and useful low-latency prompt for your shell
// Homepage: https://github.com/justjanne/powerline-go

import (
	"fmt"
	
	"os/exec"
)

func installPowerlineGo() {
	// Método 1: Descargar y extraer .tar.gz
	powerlinego_tar_url := "https://github.com/justjanne/powerline-go/archive/refs/tags/v1.24.tar.gz"
	powerlinego_cmd_tar := exec.Command("curl", "-L", powerlinego_tar_url, "-o", "package.tar.gz")
	err := powerlinego_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	powerlinego_zip_url := "https://github.com/justjanne/powerline-go/archive/refs/tags/v1.24.zip"
	powerlinego_cmd_zip := exec.Command("curl", "-L", powerlinego_zip_url, "-o", "package.zip")
	err = powerlinego_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	powerlinego_bin_url := "https://github.com/justjanne/powerline-go/archive/refs/tags/v1.24.bin"
	powerlinego_cmd_bin := exec.Command("curl", "-L", powerlinego_bin_url, "-o", "binary.bin")
	err = powerlinego_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	powerlinego_src_url := "https://github.com/justjanne/powerline-go/archive/refs/tags/v1.24.src.tar.gz"
	powerlinego_cmd_src := exec.Command("curl", "-L", powerlinego_src_url, "-o", "source.tar.gz")
	err = powerlinego_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	powerlinego_cmd_direct := exec.Command("./binary")
	err = powerlinego_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
