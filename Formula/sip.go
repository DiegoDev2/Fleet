package main

// Sip - Tool to create Python bindings for C and C++ libraries
// Homepage: https://python-sip.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installSip() {
	// Método 1: Descargar y extraer .tar.gz
	sip_tar_url := "https://files.pythonhosted.org/packages/6e/52/36987b182711104d5e9f8831dd989085b1241fc627829c36ddf81640c372/sip-6.8.6.tar.gz"
	sip_cmd_tar := exec.Command("curl", "-L", sip_tar_url, "-o", "package.tar.gz")
	err := sip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sip_zip_url := "https://files.pythonhosted.org/packages/6e/52/36987b182711104d5e9f8831dd989085b1241fc627829c36ddf81640c372/sip-6.8.6.zip"
	sip_cmd_zip := exec.Command("curl", "-L", sip_zip_url, "-o", "package.zip")
	err = sip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sip_bin_url := "https://files.pythonhosted.org/packages/6e/52/36987b182711104d5e9f8831dd989085b1241fc627829c36ddf81640c372/sip-6.8.6.bin"
	sip_cmd_bin := exec.Command("curl", "-L", sip_bin_url, "-o", "binary.bin")
	err = sip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sip_src_url := "https://files.pythonhosted.org/packages/6e/52/36987b182711104d5e9f8831dd989085b1241fc627829c36ddf81640c372/sip-6.8.6.src.tar.gz"
	sip_cmd_src := exec.Command("curl", "-L", sip_src_url, "-o", "source.tar.gz")
	err = sip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sip_cmd_direct := exec.Command("./binary")
	err = sip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
