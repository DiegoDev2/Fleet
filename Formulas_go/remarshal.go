package main

// Remarshal - Convert between TOML, YAML and JSON
// Homepage: https://github.com/remarshal-project/remarshal

import (
	"fmt"
	
	"os/exec"
)

func installRemarshal() {
	// Método 1: Descargar y extraer .tar.gz
	remarshal_tar_url := "https://files.pythonhosted.org/packages/ae/e7/98e22a9d62ee2c086da94f6aafeee5e7c3a68197d761cc22c90e7e949afb/remarshal-0.18.0.tar.gz"
	remarshal_cmd_tar := exec.Command("curl", "-L", remarshal_tar_url, "-o", "package.tar.gz")
	err := remarshal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	remarshal_zip_url := "https://files.pythonhosted.org/packages/ae/e7/98e22a9d62ee2c086da94f6aafeee5e7c3a68197d761cc22c90e7e949afb/remarshal-0.18.0.zip"
	remarshal_cmd_zip := exec.Command("curl", "-L", remarshal_zip_url, "-o", "package.zip")
	err = remarshal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	remarshal_bin_url := "https://files.pythonhosted.org/packages/ae/e7/98e22a9d62ee2c086da94f6aafeee5e7c3a68197d761cc22c90e7e949afb/remarshal-0.18.0.bin"
	remarshal_cmd_bin := exec.Command("curl", "-L", remarshal_bin_url, "-o", "binary.bin")
	err = remarshal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	remarshal_src_url := "https://files.pythonhosted.org/packages/ae/e7/98e22a9d62ee2c086da94f6aafeee5e7c3a68197d761cc22c90e7e949afb/remarshal-0.18.0.src.tar.gz"
	remarshal_cmd_src := exec.Command("curl", "-L", remarshal_src_url, "-o", "source.tar.gz")
	err = remarshal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	remarshal_cmd_direct := exec.Command("./binary")
	err = remarshal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
