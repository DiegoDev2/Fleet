package main

// Fierce - DNS reconnaissance tool for locating non-contiguous IP space
// Homepage: https://github.com/mschwager/fierce

import (
	"fmt"
	
	"os/exec"
)

func installFierce() {
	// Método 1: Descargar y extraer .tar.gz
	fierce_tar_url := "https://github.com/mschwager/fierce/archive/refs/tags/1.6.0.tar.gz"
	fierce_cmd_tar := exec.Command("curl", "-L", fierce_tar_url, "-o", "package.tar.gz")
	err := fierce_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fierce_zip_url := "https://github.com/mschwager/fierce/archive/refs/tags/1.6.0.zip"
	fierce_cmd_zip := exec.Command("curl", "-L", fierce_zip_url, "-o", "package.zip")
	err = fierce_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fierce_bin_url := "https://github.com/mschwager/fierce/archive/refs/tags/1.6.0.bin"
	fierce_cmd_bin := exec.Command("curl", "-L", fierce_bin_url, "-o", "binary.bin")
	err = fierce_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fierce_src_url := "https://github.com/mschwager/fierce/archive/refs/tags/1.6.0.src.tar.gz"
	fierce_cmd_src := exec.Command("curl", "-L", fierce_src_url, "-o", "source.tar.gz")
	err = fierce_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fierce_cmd_direct := exec.Command("./binary")
	err = fierce_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
