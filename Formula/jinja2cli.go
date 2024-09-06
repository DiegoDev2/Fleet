package main

// Jinja2Cli - CLI for the Jinja2 templating language
// Homepage: https://github.com/mattrobenolt/jinja2-cli

import (
	"fmt"
	
	"os/exec"
)

func installJinja2Cli() {
	// Método 1: Descargar y extraer .tar.gz
	jinja2cli_tar_url := "https://files.pythonhosted.org/packages/a4/22/c922839761b311b72ccc95c2ca2239311a3e80916458878962626f96922a/jinja2-cli-0.8.2.tar.gz"
	jinja2cli_cmd_tar := exec.Command("curl", "-L", jinja2cli_tar_url, "-o", "package.tar.gz")
	err := jinja2cli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jinja2cli_zip_url := "https://files.pythonhosted.org/packages/a4/22/c922839761b311b72ccc95c2ca2239311a3e80916458878962626f96922a/jinja2-cli-0.8.2.zip"
	jinja2cli_cmd_zip := exec.Command("curl", "-L", jinja2cli_zip_url, "-o", "package.zip")
	err = jinja2cli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jinja2cli_bin_url := "https://files.pythonhosted.org/packages/a4/22/c922839761b311b72ccc95c2ca2239311a3e80916458878962626f96922a/jinja2-cli-0.8.2.bin"
	jinja2cli_cmd_bin := exec.Command("curl", "-L", jinja2cli_bin_url, "-o", "binary.bin")
	err = jinja2cli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jinja2cli_src_url := "https://files.pythonhosted.org/packages/a4/22/c922839761b311b72ccc95c2ca2239311a3e80916458878962626f96922a/jinja2-cli-0.8.2.src.tar.gz"
	jinja2cli_cmd_src := exec.Command("curl", "-L", jinja2cli_src_url, "-o", "source.tar.gz")
	err = jinja2cli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jinja2cli_cmd_direct := exec.Command("./binary")
	err = jinja2cli_cmd_direct.Run()
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
