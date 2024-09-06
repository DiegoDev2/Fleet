package main

// AnsibleLanguageServer - Language Server for Ansible Files
// Homepage: https://github.com/ansible/vscode-ansible

import (
	"fmt"
	
	"os/exec"
)

func installAnsibleLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	ansiblelanguageserver_tar_url := "https://registry.npmjs.org/@ansible/ansible-language-server/-/ansible-language-server-1.2.3.tgz"
	ansiblelanguageserver_cmd_tar := exec.Command("curl", "-L", ansiblelanguageserver_tar_url, "-o", "package.tar.gz")
	err := ansiblelanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansiblelanguageserver_zip_url := "https://registry.npmjs.org/@ansible/ansible-language-server/-/ansible-language-server-1.2.3.tgz"
	ansiblelanguageserver_cmd_zip := exec.Command("curl", "-L", ansiblelanguageserver_zip_url, "-o", "package.zip")
	err = ansiblelanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansiblelanguageserver_bin_url := "https://registry.npmjs.org/@ansible/ansible-language-server/-/ansible-language-server-1.2.3.tgz"
	ansiblelanguageserver_cmd_bin := exec.Command("curl", "-L", ansiblelanguageserver_bin_url, "-o", "binary.bin")
	err = ansiblelanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansiblelanguageserver_src_url := "https://registry.npmjs.org/@ansible/ansible-language-server/-/ansible-language-server-1.2.3.tgz"
	ansiblelanguageserver_cmd_src := exec.Command("curl", "-L", ansiblelanguageserver_src_url, "-o", "source.tar.gz")
	err = ansiblelanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansiblelanguageserver_cmd_direct := exec.Command("./binary")
	err = ansiblelanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
