package main

// MarpCli - Easily convert Marp Markdown files into static HTML/CSS, PDF, PPT and images
// Homepage: https://github.com/marp-team/marp-cli

import (
	"fmt"
	
	"os/exec"
)

func installMarpCli() {
	// Método 1: Descargar y extraer .tar.gz
	marpcli_tar_url := "https://registry.npmjs.org/@marp-team/marp-cli/-/marp-cli-3.4.0.tgz"
	marpcli_cmd_tar := exec.Command("curl", "-L", marpcli_tar_url, "-o", "package.tar.gz")
	err := marpcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	marpcli_zip_url := "https://registry.npmjs.org/@marp-team/marp-cli/-/marp-cli-3.4.0.tgz"
	marpcli_cmd_zip := exec.Command("curl", "-L", marpcli_zip_url, "-o", "package.zip")
	err = marpcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	marpcli_bin_url := "https://registry.npmjs.org/@marp-team/marp-cli/-/marp-cli-3.4.0.tgz"
	marpcli_cmd_bin := exec.Command("curl", "-L", marpcli_bin_url, "-o", "binary.bin")
	err = marpcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	marpcli_src_url := "https://registry.npmjs.org/@marp-team/marp-cli/-/marp-cli-3.4.0.tgz"
	marpcli_cmd_src := exec.Command("curl", "-L", marpcli_src_url, "-o", "source.tar.gz")
	err = marpcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	marpcli_cmd_direct := exec.Command("./binary")
	err = marpcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
