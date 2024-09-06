package main

// C7n - Rules engine for cloud security, cost optimization, and governance
// Homepage: https://github.com/cloud-custodian/cloud-custodian

import (
	"fmt"
	
	"os/exec"
)

func installC7n() {
	// Método 1: Descargar y extraer .tar.gz
	c7n_tar_url := "https://github.com/cloud-custodian/cloud-custodian/archive/refs/tags/0.9.40.0.tar.gz"
	c7n_cmd_tar := exec.Command("curl", "-L", c7n_tar_url, "-o", "package.tar.gz")
	err := c7n_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	c7n_zip_url := "https://github.com/cloud-custodian/cloud-custodian/archive/refs/tags/0.9.40.0.zip"
	c7n_cmd_zip := exec.Command("curl", "-L", c7n_zip_url, "-o", "package.zip")
	err = c7n_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	c7n_bin_url := "https://github.com/cloud-custodian/cloud-custodian/archive/refs/tags/0.9.40.0.bin"
	c7n_cmd_bin := exec.Command("curl", "-L", c7n_bin_url, "-o", "binary.bin")
	err = c7n_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	c7n_src_url := "https://github.com/cloud-custodian/cloud-custodian/archive/refs/tags/0.9.40.0.src.tar.gz"
	c7n_cmd_src := exec.Command("curl", "-L", c7n_src_url, "-o", "source.tar.gz")
	err = c7n_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	c7n_cmd_direct := exec.Command("./binary")
	err = c7n_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
