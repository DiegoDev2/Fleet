package main

// OktaAwscli - Okta authentication for awscli
// Homepage: https://github.com/okta-awscli/okta-awscli

import (
	"fmt"
	
	"os/exec"
)

func installOktaAwscli() {
	// Método 1: Descargar y extraer .tar.gz
	oktaawscli_tar_url := "https://files.pythonhosted.org/packages/ed/2c/153d8ba330660d756fe6373fb4d1c13b99e63675570042de45aedf300bb7/okta-awscli-0.5.5.tar.gz"
	oktaawscli_cmd_tar := exec.Command("curl", "-L", oktaawscli_tar_url, "-o", "package.tar.gz")
	err := oktaawscli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oktaawscli_zip_url := "https://files.pythonhosted.org/packages/ed/2c/153d8ba330660d756fe6373fb4d1c13b99e63675570042de45aedf300bb7/okta-awscli-0.5.5.zip"
	oktaawscli_cmd_zip := exec.Command("curl", "-L", oktaawscli_zip_url, "-o", "package.zip")
	err = oktaawscli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oktaawscli_bin_url := "https://files.pythonhosted.org/packages/ed/2c/153d8ba330660d756fe6373fb4d1c13b99e63675570042de45aedf300bb7/okta-awscli-0.5.5.bin"
	oktaawscli_cmd_bin := exec.Command("curl", "-L", oktaawscli_bin_url, "-o", "binary.bin")
	err = oktaawscli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oktaawscli_src_url := "https://files.pythonhosted.org/packages/ed/2c/153d8ba330660d756fe6373fb4d1c13b99e63675570042de45aedf300bb7/okta-awscli-0.5.5.src.tar.gz"
	oktaawscli_cmd_src := exec.Command("curl", "-L", oktaawscli_src_url, "-o", "source.tar.gz")
	err = oktaawscli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oktaawscli_cmd_direct := exec.Command("./binary")
	err = oktaawscli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
