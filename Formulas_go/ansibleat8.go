package main

// AnsibleAT8 - Automate deployment, configuration, and upgrading
// Homepage: https://www.ansible.com/

import (
	"fmt"
	
	"os/exec"
)

func installAnsibleAT8() {
	// Método 1: Descargar y extraer .tar.gz
	ansibleat8_tar_url := "https://files.pythonhosted.org/packages/90/25/55e09468efe564f3b48c47a7e082bd84d4f0d064af60ac8458eba4667994/ansible-8.7.0.tar.gz"
	ansibleat8_cmd_tar := exec.Command("curl", "-L", ansibleat8_tar_url, "-o", "package.tar.gz")
	err := ansibleat8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansibleat8_zip_url := "https://files.pythonhosted.org/packages/90/25/55e09468efe564f3b48c47a7e082bd84d4f0d064af60ac8458eba4667994/ansible-8.7.0.zip"
	ansibleat8_cmd_zip := exec.Command("curl", "-L", ansibleat8_zip_url, "-o", "package.zip")
	err = ansibleat8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansibleat8_bin_url := "https://files.pythonhosted.org/packages/90/25/55e09468efe564f3b48c47a7e082bd84d4f0d064af60ac8458eba4667994/ansible-8.7.0.bin"
	ansibleat8_cmd_bin := exec.Command("curl", "-L", ansibleat8_bin_url, "-o", "binary.bin")
	err = ansibleat8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansibleat8_src_url := "https://files.pythonhosted.org/packages/90/25/55e09468efe564f3b48c47a7e082bd84d4f0d064af60ac8458eba4667994/ansible-8.7.0.src.tar.gz"
	ansibleat8_cmd_src := exec.Command("curl", "-L", ansibleat8_src_url, "-o", "source.tar.gz")
	err = ansibleat8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansibleat8_cmd_direct := exec.Command("./binary")
	err = ansibleat8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
}
