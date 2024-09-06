package main

// Ansible - Automate deployment, configuration, and upgrading
// Homepage: https://www.ansible.com/

import (
	"fmt"
	
	"os/exec"
)

func installAnsible() {
	// Método 1: Descargar y extraer .tar.gz
	ansible_tar_url := "https://files.pythonhosted.org/packages/94/e2/b34a077f05f97982dfddf6ec17ab5e2e476412b3ff542e30d21d22cf2e2d/ansible-10.3.0.tar.gz"
	ansible_cmd_tar := exec.Command("curl", "-L", ansible_tar_url, "-o", "package.tar.gz")
	err := ansible_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansible_zip_url := "https://files.pythonhosted.org/packages/94/e2/b34a077f05f97982dfddf6ec17ab5e2e476412b3ff542e30d21d22cf2e2d/ansible-10.3.0.zip"
	ansible_cmd_zip := exec.Command("curl", "-L", ansible_zip_url, "-o", "package.zip")
	err = ansible_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansible_bin_url := "https://files.pythonhosted.org/packages/94/e2/b34a077f05f97982dfddf6ec17ab5e2e476412b3ff542e30d21d22cf2e2d/ansible-10.3.0.bin"
	ansible_cmd_bin := exec.Command("curl", "-L", ansible_bin_url, "-o", "binary.bin")
	err = ansible_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansible_src_url := "https://files.pythonhosted.org/packages/94/e2/b34a077f05f97982dfddf6ec17ab5e2e476412b3ff542e30d21d22cf2e2d/ansible-10.3.0.src.tar.gz"
	ansible_cmd_src := exec.Command("curl", "-L", ansible_src_url, "-o", "source.tar.gz")
	err = ansible_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansible_cmd_direct := exec.Command("./binary")
	err = ansible_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libssh")
	exec.Command("latte", "install", "libssh").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
