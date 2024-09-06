package main

// Awscurl - Curl like simplicity to access AWS resources
// Homepage: https://github.com/okigan/awscurl

import (
	"fmt"
	
	"os/exec"
)

func installAwscurl() {
	// Método 1: Descargar y extraer .tar.gz
	awscurl_tar_url := "https://files.pythonhosted.org/packages/f0/53/68500d2e61aff7549f878a9227eea5c80eaf6ffcad7c134c576360b1bae7/awscurl-0.36.tar.gz"
	awscurl_cmd_tar := exec.Command("curl", "-L", awscurl_tar_url, "-o", "package.tar.gz")
	err := awscurl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awscurl_zip_url := "https://files.pythonhosted.org/packages/f0/53/68500d2e61aff7549f878a9227eea5c80eaf6ffcad7c134c576360b1bae7/awscurl-0.36.zip"
	awscurl_cmd_zip := exec.Command("curl", "-L", awscurl_zip_url, "-o", "package.zip")
	err = awscurl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awscurl_bin_url := "https://files.pythonhosted.org/packages/f0/53/68500d2e61aff7549f878a9227eea5c80eaf6ffcad7c134c576360b1bae7/awscurl-0.36.bin"
	awscurl_cmd_bin := exec.Command("curl", "-L", awscurl_bin_url, "-o", "binary.bin")
	err = awscurl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awscurl_src_url := "https://files.pythonhosted.org/packages/f0/53/68500d2e61aff7549f878a9227eea5c80eaf6ffcad7c134c576360b1bae7/awscurl-0.36.src.tar.gz"
	awscurl_cmd_src := exec.Command("curl", "-L", awscurl_src_url, "-o", "source.tar.gz")
	err = awscurl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awscurl_cmd_direct := exec.Command("./binary")
	err = awscurl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
