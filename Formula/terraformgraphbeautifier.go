package main

// TerraformGraphBeautifier - CLI to beautify `terraform graph` output
// Homepage: https://github.com/pcasteran/terraform-graph-beautifier

import (
	"fmt"
	
	"os/exec"
)

func installTerraformGraphBeautifier() {
	// Método 1: Descargar y extraer .tar.gz
	terraformgraphbeautifier_tar_url := "https://github.com/pcasteran/terraform-graph-beautifier/archive/refs/tags/v0.3.4.tar.gz"
	terraformgraphbeautifier_cmd_tar := exec.Command("curl", "-L", terraformgraphbeautifier_tar_url, "-o", "package.tar.gz")
	err := terraformgraphbeautifier_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terraformgraphbeautifier_zip_url := "https://github.com/pcasteran/terraform-graph-beautifier/archive/refs/tags/v0.3.4.zip"
	terraformgraphbeautifier_cmd_zip := exec.Command("curl", "-L", terraformgraphbeautifier_zip_url, "-o", "package.zip")
	err = terraformgraphbeautifier_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terraformgraphbeautifier_bin_url := "https://github.com/pcasteran/terraform-graph-beautifier/archive/refs/tags/v0.3.4.bin"
	terraformgraphbeautifier_cmd_bin := exec.Command("curl", "-L", terraformgraphbeautifier_bin_url, "-o", "binary.bin")
	err = terraformgraphbeautifier_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terraformgraphbeautifier_src_url := "https://github.com/pcasteran/terraform-graph-beautifier/archive/refs/tags/v0.3.4.src.tar.gz"
	terraformgraphbeautifier_cmd_src := exec.Command("curl", "-L", terraformgraphbeautifier_src_url, "-o", "source.tar.gz")
	err = terraformgraphbeautifier_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terraformgraphbeautifier_cmd_direct := exec.Command("./binary")
	err = terraformgraphbeautifier_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
