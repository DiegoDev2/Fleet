package main

// AnsibleAT9 - Automate deployment, configuration, and upgrading
// Homepage: https://www.ansible.com/

import (
	"fmt"
	
	"os/exec"
)

func installAnsibleAT9() {
	// Método 1: Descargar y extraer .tar.gz
	ansibleat9_tar_url := "https://files.pythonhosted.org/packages/e4/85/23d1a1884f8c6bd437edc5a0f55709c77250f357a198469d8060071237f9/ansible-9.6.0.tar.gz"
	ansibleat9_cmd_tar := exec.Command("curl", "-L", ansibleat9_tar_url, "-o", "package.tar.gz")
	err := ansibleat9_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansibleat9_zip_url := "https://files.pythonhosted.org/packages/e4/85/23d1a1884f8c6bd437edc5a0f55709c77250f357a198469d8060071237f9/ansible-9.6.0.zip"
	ansibleat9_cmd_zip := exec.Command("curl", "-L", ansibleat9_zip_url, "-o", "package.zip")
	err = ansibleat9_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansibleat9_bin_url := "https://files.pythonhosted.org/packages/e4/85/23d1a1884f8c6bd437edc5a0f55709c77250f357a198469d8060071237f9/ansible-9.6.0.bin"
	ansibleat9_cmd_bin := exec.Command("curl", "-L", ansibleat9_bin_url, "-o", "binary.bin")
	err = ansibleat9_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansibleat9_src_url := "https://files.pythonhosted.org/packages/e4/85/23d1a1884f8c6bd437edc5a0f55709c77250f357a198469d8060071237f9/ansible-9.6.0.src.tar.gz"
	ansibleat9_cmd_src := exec.Command("curl", "-L", ansibleat9_src_url, "-o", "source.tar.gz")
	err = ansibleat9_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansibleat9_cmd_direct := exec.Command("./binary")
	err = ansibleat9_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: libssh")
exec.Command("latte", "install", "libssh")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
