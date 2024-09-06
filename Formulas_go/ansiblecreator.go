package main

// AnsibleCreator - CLI tool for scaffolding Ansible Content
// Homepage: https://ansible.readthedocs.io/projects/creator/

import (
	"fmt"
	
	"os/exec"
)

func installAnsibleCreator() {
	// Método 1: Descargar y extraer .tar.gz
	ansiblecreator_tar_url := "https://files.pythonhosted.org/packages/f4/cc/4933e56147b0c98d18507f563074d0a06eb46587e932a378795a5fef638b/ansible_creator-24.8.0.tar.gz"
	ansiblecreator_cmd_tar := exec.Command("curl", "-L", ansiblecreator_tar_url, "-o", "package.tar.gz")
	err := ansiblecreator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ansiblecreator_zip_url := "https://files.pythonhosted.org/packages/f4/cc/4933e56147b0c98d18507f563074d0a06eb46587e932a378795a5fef638b/ansible_creator-24.8.0.zip"
	ansiblecreator_cmd_zip := exec.Command("curl", "-L", ansiblecreator_zip_url, "-o", "package.zip")
	err = ansiblecreator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ansiblecreator_bin_url := "https://files.pythonhosted.org/packages/f4/cc/4933e56147b0c98d18507f563074d0a06eb46587e932a378795a5fef638b/ansible_creator-24.8.0.bin"
	ansiblecreator_cmd_bin := exec.Command("curl", "-L", ansiblecreator_bin_url, "-o", "binary.bin")
	err = ansiblecreator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ansiblecreator_src_url := "https://files.pythonhosted.org/packages/f4/cc/4933e56147b0c98d18507f563074d0a06eb46587e932a378795a5fef638b/ansible_creator-24.8.0.src.tar.gz"
	ansiblecreator_cmd_src := exec.Command("curl", "-L", ansiblecreator_src_url, "-o", "source.tar.gz")
	err = ansiblecreator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ansiblecreator_cmd_direct := exec.Command("./binary")
	err = ansiblecreator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
