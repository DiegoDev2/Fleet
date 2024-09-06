package main

// Swagger2markupCli - Swagger to AsciiDoc or Markdown converter
// Homepage: https://github.com/Swagger2Markup/swagger2markup

import (
	"fmt"
	
	"os/exec"
)

func installSwagger2markupCli() {
	// Método 1: Descargar y extraer .tar.gz
	swagger2markupcli_tar_url := "https://search.maven.org/remotecontent?filepath=io/github/swagger2markup/swagger2markup-cli/1.3.3/swagger2markup-cli-1.3.3.jar"
	swagger2markupcli_cmd_tar := exec.Command("curl", "-L", swagger2markupcli_tar_url, "-o", "package.tar.gz")
	err := swagger2markupcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swagger2markupcli_zip_url := "https://search.maven.org/remotecontent?filepath=io/github/swagger2markup/swagger2markup-cli/1.3.3/swagger2markup-cli-1.3.3.jar"
	swagger2markupcli_cmd_zip := exec.Command("curl", "-L", swagger2markupcli_zip_url, "-o", "package.zip")
	err = swagger2markupcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swagger2markupcli_bin_url := "https://search.maven.org/remotecontent?filepath=io/github/swagger2markup/swagger2markup-cli/1.3.3/swagger2markup-cli-1.3.3.jar"
	swagger2markupcli_cmd_bin := exec.Command("curl", "-L", swagger2markupcli_bin_url, "-o", "binary.bin")
	err = swagger2markupcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swagger2markupcli_src_url := "https://search.maven.org/remotecontent?filepath=io/github/swagger2markup/swagger2markup-cli/1.3.3/swagger2markup-cli-1.3.3.jar"
	swagger2markupcli_cmd_src := exec.Command("curl", "-L", swagger2markupcli_src_url, "-o", "source.tar.gz")
	err = swagger2markupcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swagger2markupcli_cmd_direct := exec.Command("./binary")
	err = swagger2markupcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
