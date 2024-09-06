package main

// AnsibleAT7 - Automate deployment, configuration, and upgrading
// Homepage: https://www.ansible.com/

import (
	"fmt"
	
	"os/exec"
)

func installAnsibleAT7() {
	// Método 1: Descargar y extraer .tar.gz
	ansibleat7_tar_url := "https://files.pythonhosted.org/packages/39/47/bef8fd8bc2b6e7b5058b61565959c91819eccb8be119a66f8524c0252c62/ansible-7.7.0.tar.gz"
	ansibleat7_cmd_tar := exec.Command("curl", "-L", ansibleat7_tar_url, "-o", "package.tar.gz")
	err := ansibleat7_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansibleat7_zip_url := "https://files.pythonhosted.org/packages/39/47/bef8fd8bc2b6e7b5058b61565959c91819eccb8be119a66f8524c0252c62/ansible-7.7.0.zip"
	ansibleat7_cmd_zip := exec.Command("curl", "-L", ansibleat7_zip_url, "-o", "package.zip")
	err = ansibleat7_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansibleat7_bin_url := "https://files.pythonhosted.org/packages/39/47/bef8fd8bc2b6e7b5058b61565959c91819eccb8be119a66f8524c0252c62/ansible-7.7.0.bin"
	ansibleat7_cmd_bin := exec.Command("curl", "-L", ansibleat7_bin_url, "-o", "binary.bin")
	err = ansibleat7_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansibleat7_src_url := "https://files.pythonhosted.org/packages/39/47/bef8fd8bc2b6e7b5058b61565959c91819eccb8be119a66f8524c0252c62/ansible-7.7.0.src.tar.gz"
	ansibleat7_cmd_src := exec.Command("curl", "-L", ansibleat7_src_url, "-o", "source.tar.gz")
	err = ansibleat7_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansibleat7_cmd_direct := exec.Command("./binary")
	err = ansibleat7_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: cffi")
exec.Command("latte", "install", "cffi")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pycparser")
exec.Command("latte", "install", "pycparser")
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: pyyaml")
exec.Command("latte", "install", "pyyaml")
	fmt.Println("Instalando dependencia: six")
exec.Command("latte", "install", "six")
}
