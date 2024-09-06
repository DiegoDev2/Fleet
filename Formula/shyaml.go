package main

// Shyaml - Command-line YAML parser
// Homepage: https://github.com/0k/shyaml

import (
	"fmt"
	
	"os/exec"
)

func installShyaml() {
	// Método 1: Descargar y extraer .tar.gz
	shyaml_tar_url := "https://files.pythonhosted.org/packages/b9/59/7e6873fa73a476de053041d26d112b65d7e1e480b88a93b4baa77197bd04/shyaml-0.6.2.tar.gz"
	shyaml_cmd_tar := exec.Command("curl", "-L", shyaml_tar_url, "-o", "package.tar.gz")
	err := shyaml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shyaml_zip_url := "https://files.pythonhosted.org/packages/b9/59/7e6873fa73a476de053041d26d112b65d7e1e480b88a93b4baa77197bd04/shyaml-0.6.2.zip"
	shyaml_cmd_zip := exec.Command("curl", "-L", shyaml_zip_url, "-o", "package.zip")
	err = shyaml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shyaml_bin_url := "https://files.pythonhosted.org/packages/b9/59/7e6873fa73a476de053041d26d112b65d7e1e480b88a93b4baa77197bd04/shyaml-0.6.2.bin"
	shyaml_cmd_bin := exec.Command("curl", "-L", shyaml_bin_url, "-o", "binary.bin")
	err = shyaml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shyaml_src_url := "https://files.pythonhosted.org/packages/b9/59/7e6873fa73a476de053041d26d112b65d7e1e480b88a93b4baa77197bd04/shyaml-0.6.2.src.tar.gz"
	shyaml_cmd_src := exec.Command("curl", "-L", shyaml_src_url, "-o", "source.tar.gz")
	err = shyaml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shyaml_cmd_direct := exec.Command("./binary")
	err = shyaml_cmd_direct.Run()
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
