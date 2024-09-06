package main

// Countdown - Terminal countdown timer
// Homepage: https://github.com/antonmedv/countdown

import (
	"fmt"
	
	"os/exec"
)

func installCountdown() {
	// Método 1: Descargar y extraer .tar.gz
	countdown_tar_url := "https://github.com/antonmedv/countdown/archive/refs/tags/v1.5.0.tar.gz"
	countdown_cmd_tar := exec.Command("curl", "-L", countdown_tar_url, "-o", "package.tar.gz")
	err := countdown_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	countdown_zip_url := "https://github.com/antonmedv/countdown/archive/refs/tags/v1.5.0.zip"
	countdown_cmd_zip := exec.Command("curl", "-L", countdown_zip_url, "-o", "package.zip")
	err = countdown_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	countdown_bin_url := "https://github.com/antonmedv/countdown/archive/refs/tags/v1.5.0.bin"
	countdown_cmd_bin := exec.Command("curl", "-L", countdown_bin_url, "-o", "binary.bin")
	err = countdown_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	countdown_src_url := "https://github.com/antonmedv/countdown/archive/refs/tags/v1.5.0.src.tar.gz"
	countdown_cmd_src := exec.Command("curl", "-L", countdown_src_url, "-o", "source.tar.gz")
	err = countdown_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	countdown_cmd_direct := exec.Command("./binary")
	err = countdown_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
