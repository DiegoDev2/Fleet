package main

// Twoping - Ping utility to determine directional packet loss
// Homepage: https://www.finnie.org/software/2ping/

import (
	"fmt"
	
	"os/exec"
)

func installTwoping() {
	// Método 1: Descargar y extraer .tar.gz
	twoping_tar_url := "https://www.finnie.org/software/2ping/2ping-4.5.1.tar.gz"
	twoping_cmd_tar := exec.Command("curl", "-L", twoping_tar_url, "-o", "package.tar.gz")
	err := twoping_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	twoping_zip_url := "https://www.finnie.org/software/2ping/2ping-4.5.1.zip"
	twoping_cmd_zip := exec.Command("curl", "-L", twoping_zip_url, "-o", "package.zip")
	err = twoping_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	twoping_bin_url := "https://www.finnie.org/software/2ping/2ping-4.5.1.bin"
	twoping_cmd_bin := exec.Command("curl", "-L", twoping_bin_url, "-o", "binary.bin")
	err = twoping_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	twoping_src_url := "https://www.finnie.org/software/2ping/2ping-4.5.1.src.tar.gz"
	twoping_cmd_src := exec.Command("curl", "-L", twoping_src_url, "-o", "source.tar.gz")
	err = twoping_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	twoping_cmd_direct := exec.Command("./binary")
	err = twoping_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
