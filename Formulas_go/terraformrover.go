package main

// TerraformRover - Terraform Visualizer
// Homepage: https://github.com/im2nguyen/rover

import (
	"fmt"
	
	"os/exec"
)

func installTerraformRover() {
	// Método 1: Descargar y extraer .tar.gz
	terraformrover_tar_url := "https://github.com/im2nguyen/rover/archive/refs/tags/v0.3.3.tar.gz"
	terraformrover_cmd_tar := exec.Command("curl", "-L", terraformrover_tar_url, "-o", "package.tar.gz")
	err := terraformrover_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terraformrover_zip_url := "https://github.com/im2nguyen/rover/archive/refs/tags/v0.3.3.zip"
	terraformrover_cmd_zip := exec.Command("curl", "-L", terraformrover_zip_url, "-o", "package.zip")
	err = terraformrover_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terraformrover_bin_url := "https://github.com/im2nguyen/rover/archive/refs/tags/v0.3.3.bin"
	terraformrover_cmd_bin := exec.Command("curl", "-L", terraformrover_bin_url, "-o", "binary.bin")
	err = terraformrover_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terraformrover_src_url := "https://github.com/im2nguyen/rover/archive/refs/tags/v0.3.3.src.tar.gz"
	terraformrover_cmd_src := exec.Command("curl", "-L", terraformrover_src_url, "-o", "source.tar.gz")
	err = terraformrover_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terraformrover_cmd_direct := exec.Command("./binary")
	err = terraformrover_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: terraform")
exec.Command("latte", "install", "terraform")
}
