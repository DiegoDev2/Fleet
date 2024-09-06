package main

// Parliament - AWS IAM linting library
// Homepage: https://github.com/duo-labs/parliament

import (
	"fmt"
	
	"os/exec"
)

func installParliament() {
	// Método 1: Descargar y extraer .tar.gz
	parliament_tar_url := "https://files.pythonhosted.org/packages/a6/3f/b7262b8a7c8d41c243950c5858cefc29652623599a6fafb2f753621b9702/parliament-1.6.3.tar.gz"
	parliament_cmd_tar := exec.Command("curl", "-L", parliament_tar_url, "-o", "package.tar.gz")
	err := parliament_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	parliament_zip_url := "https://files.pythonhosted.org/packages/a6/3f/b7262b8a7c8d41c243950c5858cefc29652623599a6fafb2f753621b9702/parliament-1.6.3.zip"
	parliament_cmd_zip := exec.Command("curl", "-L", parliament_zip_url, "-o", "package.zip")
	err = parliament_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	parliament_bin_url := "https://files.pythonhosted.org/packages/a6/3f/b7262b8a7c8d41c243950c5858cefc29652623599a6fafb2f753621b9702/parliament-1.6.3.bin"
	parliament_cmd_bin := exec.Command("curl", "-L", parliament_bin_url, "-o", "binary.bin")
	err = parliament_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	parliament_src_url := "https://files.pythonhosted.org/packages/a6/3f/b7262b8a7c8d41c243950c5858cefc29652623599a6fafb2f753621b9702/parliament-1.6.3.src.tar.gz"
	parliament_cmd_src := exec.Command("curl", "-L", parliament_src_url, "-o", "source.tar.gz")
	err = parliament_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	parliament_cmd_direct := exec.Command("./binary")
	err = parliament_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
