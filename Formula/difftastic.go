package main

// Difftastic - Diff that understands syntax
// Homepage: https://github.com/Wilfred/difftastic

import (
	"fmt"
	
	"os/exec"
)

func installDifftastic() {
	// Método 1: Descargar y extraer .tar.gz
	difftastic_tar_url := "https://github.com/Wilfred/difftastic/archive/refs/tags/0.60.0.tar.gz"
	difftastic_cmd_tar := exec.Command("curl", "-L", difftastic_tar_url, "-o", "package.tar.gz")
	err := difftastic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	difftastic_zip_url := "https://github.com/Wilfred/difftastic/archive/refs/tags/0.60.0.zip"
	difftastic_cmd_zip := exec.Command("curl", "-L", difftastic_zip_url, "-o", "package.zip")
	err = difftastic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	difftastic_bin_url := "https://github.com/Wilfred/difftastic/archive/refs/tags/0.60.0.bin"
	difftastic_cmd_bin := exec.Command("curl", "-L", difftastic_bin_url, "-o", "binary.bin")
	err = difftastic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	difftastic_src_url := "https://github.com/Wilfred/difftastic/archive/refs/tags/0.60.0.src.tar.gz"
	difftastic_cmd_src := exec.Command("curl", "-L", difftastic_src_url, "-o", "source.tar.gz")
	err = difftastic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	difftastic_cmd_direct := exec.Command("./binary")
	err = difftastic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
