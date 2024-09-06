package main

// Virtualfish - Python virtual environment manager for the fish shell
// Homepage: https://virtualfish.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installVirtualfish() {
	// Método 1: Descargar y extraer .tar.gz
	virtualfish_tar_url := "https://files.pythonhosted.org/packages/1f/4e/343d044d61e80a44163d15ad2f6ca20eca0cb4fef4058caf8e5e55fc3dd9/virtualfish-2.5.9.tar.gz"
	virtualfish_cmd_tar := exec.Command("curl", "-L", virtualfish_tar_url, "-o", "package.tar.gz")
	err := virtualfish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	virtualfish_zip_url := "https://files.pythonhosted.org/packages/1f/4e/343d044d61e80a44163d15ad2f6ca20eca0cb4fef4058caf8e5e55fc3dd9/virtualfish-2.5.9.zip"
	virtualfish_cmd_zip := exec.Command("curl", "-L", virtualfish_zip_url, "-o", "package.zip")
	err = virtualfish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	virtualfish_bin_url := "https://files.pythonhosted.org/packages/1f/4e/343d044d61e80a44163d15ad2f6ca20eca0cb4fef4058caf8e5e55fc3dd9/virtualfish-2.5.9.bin"
	virtualfish_cmd_bin := exec.Command("curl", "-L", virtualfish_bin_url, "-o", "binary.bin")
	err = virtualfish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	virtualfish_src_url := "https://files.pythonhosted.org/packages/1f/4e/343d044d61e80a44163d15ad2f6ca20eca0cb4fef4058caf8e5e55fc3dd9/virtualfish-2.5.9.src.tar.gz"
	virtualfish_cmd_src := exec.Command("curl", "-L", virtualfish_src_url, "-o", "source.tar.gz")
	err = virtualfish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	virtualfish_cmd_direct := exec.Command("./binary")
	err = virtualfish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fish")
exec.Command("latte", "install", "fish")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
