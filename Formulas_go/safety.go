package main

// Safety - Checks Python dependencies for known vulnerabilities and suggests remediations
// Homepage: https://safetycli.com/product/safety-cli

import (
	"fmt"
	
	"os/exec"
)

func installSafety() {
	// Método 1: Descargar y extraer .tar.gz
	safety_tar_url := "https://files.pythonhosted.org/packages/44/76/08e25cff4ecaa9ad75f2ba3bf0d94b534400712289ea7a4ca07694348379/safety-3.2.7.tar.gz"
	safety_cmd_tar := exec.Command("curl", "-L", safety_tar_url, "-o", "package.tar.gz")
	err := safety_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	safety_zip_url := "https://files.pythonhosted.org/packages/44/76/08e25cff4ecaa9ad75f2ba3bf0d94b534400712289ea7a4ca07694348379/safety-3.2.7.zip"
	safety_cmd_zip := exec.Command("curl", "-L", safety_zip_url, "-o", "package.zip")
	err = safety_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	safety_bin_url := "https://files.pythonhosted.org/packages/44/76/08e25cff4ecaa9ad75f2ba3bf0d94b534400712289ea7a4ca07694348379/safety-3.2.7.bin"
	safety_cmd_bin := exec.Command("curl", "-L", safety_bin_url, "-o", "binary.bin")
	err = safety_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	safety_src_url := "https://files.pythonhosted.org/packages/44/76/08e25cff4ecaa9ad75f2ba3bf0d94b534400712289ea7a4ca07694348379/safety-3.2.7.src.tar.gz"
	safety_cmd_src := exec.Command("curl", "-L", safety_src_url, "-o", "source.tar.gz")
	err = safety_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	safety_cmd_direct := exec.Command("./binary")
	err = safety_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
