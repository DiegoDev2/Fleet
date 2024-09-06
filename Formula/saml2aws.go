package main

// Saml2aws - Login and retrieve AWS temporary credentials using a SAML IDP
// Homepage: https://github.com/Versent/saml2aws

import (
	"fmt"
	
	"os/exec"
)

func installSaml2aws() {
	// Método 1: Descargar y extraer .tar.gz
	saml2aws_tar_url := "https://github.com/Versent/saml2aws/archive/refs/tags/v2.36.17.tar.gz"
	saml2aws_cmd_tar := exec.Command("curl", "-L", saml2aws_tar_url, "-o", "package.tar.gz")
	err := saml2aws_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	saml2aws_zip_url := "https://github.com/Versent/saml2aws/archive/refs/tags/v2.36.17.zip"
	saml2aws_cmd_zip := exec.Command("curl", "-L", saml2aws_zip_url, "-o", "package.zip")
	err = saml2aws_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	saml2aws_bin_url := "https://github.com/Versent/saml2aws/archive/refs/tags/v2.36.17.bin"
	saml2aws_cmd_bin := exec.Command("curl", "-L", saml2aws_bin_url, "-o", "binary.bin")
	err = saml2aws_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	saml2aws_src_url := "https://github.com/Versent/saml2aws/archive/refs/tags/v2.36.17.src.tar.gz"
	saml2aws_cmd_src := exec.Command("curl", "-L", saml2aws_src_url, "-o", "source.tar.gz")
	err = saml2aws_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	saml2aws_cmd_direct := exec.Command("./binary")
	err = saml2aws_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
