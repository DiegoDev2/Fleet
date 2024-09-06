package main

// OscCli - Official Outscale CLI providing connectors to Outscale API
// Homepage: https://github.com/outscale/osc-cli

import (
	"fmt"
	
	"os/exec"
)

func installOscCli() {
	// Método 1: Descargar y extraer .tar.gz
	osccli_tar_url := "https://files.pythonhosted.org/packages/02/cd/f1b796f5e7a301f6a3c0b910be07188cbfd329d2758e036d24ef26b4ee96/osc-sdk-1.11.0.tar.gz"
	osccli_cmd_tar := exec.Command("curl", "-L", osccli_tar_url, "-o", "package.tar.gz")
	err := osccli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	osccli_zip_url := "https://files.pythonhosted.org/packages/02/cd/f1b796f5e7a301f6a3c0b910be07188cbfd329d2758e036d24ef26b4ee96/osc-sdk-1.11.0.zip"
	osccli_cmd_zip := exec.Command("curl", "-L", osccli_zip_url, "-o", "package.zip")
	err = osccli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	osccli_bin_url := "https://files.pythonhosted.org/packages/02/cd/f1b796f5e7a301f6a3c0b910be07188cbfd329d2758e036d24ef26b4ee96/osc-sdk-1.11.0.bin"
	osccli_cmd_bin := exec.Command("curl", "-L", osccli_bin_url, "-o", "binary.bin")
	err = osccli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	osccli_src_url := "https://files.pythonhosted.org/packages/02/cd/f1b796f5e7a301f6a3c0b910be07188cbfd329d2758e036d24ef26b4ee96/osc-sdk-1.11.0.src.tar.gz"
	osccli_cmd_src := exec.Command("curl", "-L", osccli_src_url, "-o", "source.tar.gz")
	err = osccli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	osccli_cmd_direct := exec.Command("./binary")
	err = osccli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
