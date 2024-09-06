package main

// Cfripper - Library and CLI tool to analyse CloudFormation templates for security issues
// Homepage: https://cfripper.readthedocs.io

import (
	"fmt"
	
	"os/exec"
)

func installCfripper() {
	// Método 1: Descargar y extraer .tar.gz
	cfripper_tar_url := "https://files.pythonhosted.org/packages/d2/93/6f8a1d1bc18b933231a7d79e97beaedcae395c95dc1505f14cc56cfd46aa/cfripper-1.16.0.tar.gz"
	cfripper_cmd_tar := exec.Command("curl", "-L", cfripper_tar_url, "-o", "package.tar.gz")
	err := cfripper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cfripper_zip_url := "https://files.pythonhosted.org/packages/d2/93/6f8a1d1bc18b933231a7d79e97beaedcae395c95dc1505f14cc56cfd46aa/cfripper-1.16.0.zip"
	cfripper_cmd_zip := exec.Command("curl", "-L", cfripper_zip_url, "-o", "package.zip")
	err = cfripper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cfripper_bin_url := "https://files.pythonhosted.org/packages/d2/93/6f8a1d1bc18b933231a7d79e97beaedcae395c95dc1505f14cc56cfd46aa/cfripper-1.16.0.bin"
	cfripper_cmd_bin := exec.Command("curl", "-L", cfripper_bin_url, "-o", "binary.bin")
	err = cfripper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cfripper_src_url := "https://files.pythonhosted.org/packages/d2/93/6f8a1d1bc18b933231a7d79e97beaedcae395c95dc1505f14cc56cfd46aa/cfripper-1.16.0.src.tar.gz"
	cfripper_cmd_src := exec.Command("curl", "-L", cfripper_src_url, "-o", "source.tar.gz")
	err = cfripper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cfripper_cmd_direct := exec.Command("./binary")
	err = cfripper_cmd_direct.Run()
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
