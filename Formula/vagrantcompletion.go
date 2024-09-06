package main

// VagrantCompletion - Bash completion for Vagrant
// Homepage: https://github.com/hashicorp/vagrant

import (
	"fmt"
	
	"os/exec"
)

func installVagrantCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	vagrantcompletion_tar_url := "https://github.com/hashicorp/vagrant/archive/refs/tags/v2.3.7.tar.gz"
	vagrantcompletion_cmd_tar := exec.Command("curl", "-L", vagrantcompletion_tar_url, "-o", "package.tar.gz")
	err := vagrantcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vagrantcompletion_zip_url := "https://github.com/hashicorp/vagrant/archive/refs/tags/v2.3.7.zip"
	vagrantcompletion_cmd_zip := exec.Command("curl", "-L", vagrantcompletion_zip_url, "-o", "package.zip")
	err = vagrantcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vagrantcompletion_bin_url := "https://github.com/hashicorp/vagrant/archive/refs/tags/v2.3.7.bin"
	vagrantcompletion_cmd_bin := exec.Command("curl", "-L", vagrantcompletion_bin_url, "-o", "binary.bin")
	err = vagrantcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vagrantcompletion_src_url := "https://github.com/hashicorp/vagrant/archive/refs/tags/v2.3.7.src.tar.gz"
	vagrantcompletion_cmd_src := exec.Command("curl", "-L", vagrantcompletion_src_url, "-o", "source.tar.gz")
	err = vagrantcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vagrantcompletion_cmd_direct := exec.Command("./binary")
	err = vagrantcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
