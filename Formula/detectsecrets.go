package main

// DetectSecrets - Enterprise friendly way of detecting and preventing secrets in code
// Homepage: https://github.com/Yelp/detect-secrets

import (
	"fmt"
	
	"os/exec"
)

func installDetectSecrets() {
	// Método 1: Descargar y extraer .tar.gz
	detectsecrets_tar_url := "https://files.pythonhosted.org/packages/69/67/382a863fff94eae5a0cf05542179169a1c49a4c8784a9480621e2066ca7d/detect_secrets-1.5.0.tar.gz"
	detectsecrets_cmd_tar := exec.Command("curl", "-L", detectsecrets_tar_url, "-o", "package.tar.gz")
	err := detectsecrets_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	detectsecrets_zip_url := "https://files.pythonhosted.org/packages/69/67/382a863fff94eae5a0cf05542179169a1c49a4c8784a9480621e2066ca7d/detect_secrets-1.5.0.zip"
	detectsecrets_cmd_zip := exec.Command("curl", "-L", detectsecrets_zip_url, "-o", "package.zip")
	err = detectsecrets_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	detectsecrets_bin_url := "https://files.pythonhosted.org/packages/69/67/382a863fff94eae5a0cf05542179169a1c49a4c8784a9480621e2066ca7d/detect_secrets-1.5.0.bin"
	detectsecrets_cmd_bin := exec.Command("curl", "-L", detectsecrets_bin_url, "-o", "binary.bin")
	err = detectsecrets_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	detectsecrets_src_url := "https://files.pythonhosted.org/packages/69/67/382a863fff94eae5a0cf05542179169a1c49a4c8784a9480621e2066ca7d/detect_secrets-1.5.0.src.tar.gz"
	detectsecrets_cmd_src := exec.Command("curl", "-L", detectsecrets_src_url, "-o", "source.tar.gz")
	err = detectsecrets_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	detectsecrets_cmd_direct := exec.Command("./binary")
	err = detectsecrets_cmd_direct.Run()
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
