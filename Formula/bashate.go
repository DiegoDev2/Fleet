package main

// Bashate - Code style enforcement for bash programs
// Homepage: https://github.com/openstack/bashate

import (
	"fmt"
	
	"os/exec"
)

func installBashate() {
	// Método 1: Descargar y extraer .tar.gz
	bashate_tar_url := "https://files.pythonhosted.org/packages/4d/0c/35b92b742cc9da7788db16cfafda2f38505e19045ae1ee204ec238ece93f/bashate-2.1.1.tar.gz"
	bashate_cmd_tar := exec.Command("curl", "-L", bashate_tar_url, "-o", "package.tar.gz")
	err := bashate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bashate_zip_url := "https://files.pythonhosted.org/packages/4d/0c/35b92b742cc9da7788db16cfafda2f38505e19045ae1ee204ec238ece93f/bashate-2.1.1.zip"
	bashate_cmd_zip := exec.Command("curl", "-L", bashate_zip_url, "-o", "package.zip")
	err = bashate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bashate_bin_url := "https://files.pythonhosted.org/packages/4d/0c/35b92b742cc9da7788db16cfafda2f38505e19045ae1ee204ec238ece93f/bashate-2.1.1.bin"
	bashate_cmd_bin := exec.Command("curl", "-L", bashate_bin_url, "-o", "binary.bin")
	err = bashate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bashate_src_url := "https://files.pythonhosted.org/packages/4d/0c/35b92b742cc9da7788db16cfafda2f38505e19045ae1ee204ec238ece93f/bashate-2.1.1.src.tar.gz"
	bashate_cmd_src := exec.Command("curl", "-L", bashate_src_url, "-o", "source.tar.gz")
	err = bashate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bashate_cmd_direct := exec.Command("./binary")
	err = bashate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
