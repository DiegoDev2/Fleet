package main

// GruntCli - JavaScript Task Runner
// Homepage: https://gruntjs.com/

import (
	"fmt"
	
	"os/exec"
)

func installGruntCli() {
	// Método 1: Descargar y extraer .tar.gz
	gruntcli_tar_url := "https://registry.npmjs.org/grunt-cli/-/grunt-cli-1.5.0.tgz"
	gruntcli_cmd_tar := exec.Command("curl", "-L", gruntcli_tar_url, "-o", "package.tar.gz")
	err := gruntcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gruntcli_zip_url := "https://registry.npmjs.org/grunt-cli/-/grunt-cli-1.5.0.tgz"
	gruntcli_cmd_zip := exec.Command("curl", "-L", gruntcli_zip_url, "-o", "package.zip")
	err = gruntcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gruntcli_bin_url := "https://registry.npmjs.org/grunt-cli/-/grunt-cli-1.5.0.tgz"
	gruntcli_cmd_bin := exec.Command("curl", "-L", gruntcli_bin_url, "-o", "binary.bin")
	err = gruntcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gruntcli_src_url := "https://registry.npmjs.org/grunt-cli/-/grunt-cli-1.5.0.tgz"
	gruntcli_cmd_src := exec.Command("curl", "-L", gruntcli_src_url, "-o", "source.tar.gz")
	err = gruntcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gruntcli_cmd_direct := exec.Command("./binary")
	err = gruntcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
