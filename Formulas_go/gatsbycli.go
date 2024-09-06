package main

// GatsbyCli - Gatsby command-line interface
// Homepage: https://www.gatsbyjs.org/docs/gatsby-cli/

import (
	"fmt"
	
	"os/exec"
)

func installGatsbyCli() {
	// Método 1: Descargar y extraer .tar.gz
	gatsbycli_tar_url := "https://registry.npmjs.org/gatsby-cli/-/gatsby-cli-5.13.3.tgz"
	gatsbycli_cmd_tar := exec.Command("curl", "-L", gatsbycli_tar_url, "-o", "package.tar.gz")
	err := gatsbycli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gatsbycli_zip_url := "https://registry.npmjs.org/gatsby-cli/-/gatsby-cli-5.13.3.tgz"
	gatsbycli_cmd_zip := exec.Command("curl", "-L", gatsbycli_zip_url, "-o", "package.zip")
	err = gatsbycli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gatsbycli_bin_url := "https://registry.npmjs.org/gatsby-cli/-/gatsby-cli-5.13.3.tgz"
	gatsbycli_cmd_bin := exec.Command("curl", "-L", gatsbycli_bin_url, "-o", "binary.bin")
	err = gatsbycli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gatsbycli_src_url := "https://registry.npmjs.org/gatsby-cli/-/gatsby-cli-5.13.3.tgz"
	gatsbycli_cmd_src := exec.Command("curl", "-L", gatsbycli_src_url, "-o", "source.tar.gz")
	err = gatsbycli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gatsbycli_cmd_direct := exec.Command("./binary")
	err = gatsbycli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: xsel")
exec.Command("latte", "install", "xsel")
}
