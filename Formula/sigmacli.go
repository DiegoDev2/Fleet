package main

// SigmaCli - CLI based on pySigma
// Homepage: https://github.com/SigmaHQ/sigma-cli

import (
	"fmt"
	
	"os/exec"
)

func installSigmaCli() {
	// Método 1: Descargar y extraer .tar.gz
	sigmacli_tar_url := "https://files.pythonhosted.org/packages/70/e8/6a4e6aa2875494af43483a37c1715039d42a0ba54cb1353db5c3ebfded69/sigma_cli-1.0.4.tar.gz"
	sigmacli_cmd_tar := exec.Command("curl", "-L", sigmacli_tar_url, "-o", "package.tar.gz")
	err := sigmacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sigmacli_zip_url := "https://files.pythonhosted.org/packages/70/e8/6a4e6aa2875494af43483a37c1715039d42a0ba54cb1353db5c3ebfded69/sigma_cli-1.0.4.zip"
	sigmacli_cmd_zip := exec.Command("curl", "-L", sigmacli_zip_url, "-o", "package.zip")
	err = sigmacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sigmacli_bin_url := "https://files.pythonhosted.org/packages/70/e8/6a4e6aa2875494af43483a37c1715039d42a0ba54cb1353db5c3ebfded69/sigma_cli-1.0.4.bin"
	sigmacli_cmd_bin := exec.Command("curl", "-L", sigmacli_bin_url, "-o", "binary.bin")
	err = sigmacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sigmacli_src_url := "https://files.pythonhosted.org/packages/70/e8/6a4e6aa2875494af43483a37c1715039d42a0ba54cb1353db5c3ebfded69/sigma_cli-1.0.4.src.tar.gz"
	sigmacli_cmd_src := exec.Command("curl", "-L", sigmacli_src_url, "-o", "source.tar.gz")
	err = sigmacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sigmacli_cmd_direct := exec.Command("./binary")
	err = sigmacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
