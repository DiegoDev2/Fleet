package main

// AnsibleLint - Checks ansible playbooks for practices and behaviour
// Homepage: https://ansible-lint.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installAnsibleLint() {
	// Método 1: Descargar y extraer .tar.gz
	ansiblelint_tar_url := "https://files.pythonhosted.org/packages/c2/44/8751d04184a2c4f3e60be27ef7a14f04257096303dcb8122167672d3cd6a/ansible_lint-24.7.0.tar.gz"
	ansiblelint_cmd_tar := exec.Command("curl", "-L", ansiblelint_tar_url, "-o", "package.tar.gz")
	err := ansiblelint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansiblelint_zip_url := "https://files.pythonhosted.org/packages/c2/44/8751d04184a2c4f3e60be27ef7a14f04257096303dcb8122167672d3cd6a/ansible_lint-24.7.0.zip"
	ansiblelint_cmd_zip := exec.Command("curl", "-L", ansiblelint_zip_url, "-o", "package.zip")
	err = ansiblelint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansiblelint_bin_url := "https://files.pythonhosted.org/packages/c2/44/8751d04184a2c4f3e60be27ef7a14f04257096303dcb8122167672d3cd6a/ansible_lint-24.7.0.bin"
	ansiblelint_cmd_bin := exec.Command("curl", "-L", ansiblelint_bin_url, "-o", "binary.bin")
	err = ansiblelint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansiblelint_src_url := "https://files.pythonhosted.org/packages/c2/44/8751d04184a2c4f3e60be27ef7a14f04257096303dcb8122167672d3cd6a/ansible_lint-24.7.0.src.tar.gz"
	ansiblelint_cmd_src := exec.Command("curl", "-L", ansiblelint_src_url, "-o", "source.tar.gz")
	err = ansiblelint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansiblelint_cmd_direct := exec.Command("./binary")
	err = ansiblelint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: ansible")
exec.Command("latte", "install", "ansible")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
