package main

// Iredis - Terminal Client for Redis with AutoCompletion and Syntax Highlighting
// Homepage: https://iredis.xbin.io/

import (
	"fmt"
	
	"os/exec"
)

func installIredis() {
	// Método 1: Descargar y extraer .tar.gz
	iredis_tar_url := "https://files.pythonhosted.org/packages/7a/80/ac86d397fa0b931cfa0121ed23549a245e706b4b34e4bfc491bcd4123acf/iredis-1.15.0.tar.gz"
	iredis_cmd_tar := exec.Command("curl", "-L", iredis_tar_url, "-o", "package.tar.gz")
	err := iredis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iredis_zip_url := "https://files.pythonhosted.org/packages/7a/80/ac86d397fa0b931cfa0121ed23549a245e706b4b34e4bfc491bcd4123acf/iredis-1.15.0.zip"
	iredis_cmd_zip := exec.Command("curl", "-L", iredis_zip_url, "-o", "package.zip")
	err = iredis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iredis_bin_url := "https://files.pythonhosted.org/packages/7a/80/ac86d397fa0b931cfa0121ed23549a245e706b4b34e4bfc491bcd4123acf/iredis-1.15.0.bin"
	iredis_cmd_bin := exec.Command("curl", "-L", iredis_bin_url, "-o", "binary.bin")
	err = iredis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iredis_src_url := "https://files.pythonhosted.org/packages/7a/80/ac86d397fa0b931cfa0121ed23549a245e706b4b34e4bfc491bcd4123acf/iredis-1.15.0.src.tar.gz"
	iredis_cmd_src := exec.Command("curl", "-L", iredis_src_url, "-o", "source.tar.gz")
	err = iredis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iredis_cmd_direct := exec.Command("./binary")
	err = iredis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
