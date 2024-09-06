package main

// Cycode - Boost security in your dev lifecycle via SAST, SCA, Secrets & IaC scanning
// Homepage: https://github.com/cycodehq/cycode-cli

import (
	"fmt"
	
	"os/exec"
)

func installCycode() {
	// Método 1: Descargar y extraer .tar.gz
	cycode_tar_url := "https://files.pythonhosted.org/packages/7c/cb/956d7b17b920a3ab645cac54b3026982259cd19c93b58d04a2f759861d40/cycode-1.10.8.tar.gz"
	cycode_cmd_tar := exec.Command("curl", "-L", cycode_tar_url, "-o", "package.tar.gz")
	err := cycode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cycode_zip_url := "https://files.pythonhosted.org/packages/7c/cb/956d7b17b920a3ab645cac54b3026982259cd19c93b58d04a2f759861d40/cycode-1.10.8.zip"
	cycode_cmd_zip := exec.Command("curl", "-L", cycode_zip_url, "-o", "package.zip")
	err = cycode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cycode_bin_url := "https://files.pythonhosted.org/packages/7c/cb/956d7b17b920a3ab645cac54b3026982259cd19c93b58d04a2f759861d40/cycode-1.10.8.bin"
	cycode_cmd_bin := exec.Command("curl", "-L", cycode_bin_url, "-o", "binary.bin")
	err = cycode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cycode_src_url := "https://files.pythonhosted.org/packages/7c/cb/956d7b17b920a3ab645cac54b3026982259cd19c93b58d04a2f759861d40/cycode-1.10.8.src.tar.gz"
	cycode_cmd_src := exec.Command("curl", "-L", cycode_src_url, "-o", "source.tar.gz")
	err = cycode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cycode_cmd_direct := exec.Command("./binary")
	err = cycode_cmd_direct.Run()
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
