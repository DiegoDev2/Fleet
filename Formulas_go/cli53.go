package main

// Cli53 - Command-line tool for Amazon Route 53
// Homepage: https://github.com/barnybug/cli53

import (
	"fmt"
	
	"os/exec"
)

func installCli53() {
	// Método 1: Descargar y extraer .tar.gz
	cli53_tar_url := "https://github.com/barnybug/cli53/archive/refs/tags/0.8.22.tar.gz"
	cli53_cmd_tar := exec.Command("curl", "-L", cli53_tar_url, "-o", "package.tar.gz")
	err := cli53_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cli53_zip_url := "https://github.com/barnybug/cli53/archive/refs/tags/0.8.22.zip"
	cli53_cmd_zip := exec.Command("curl", "-L", cli53_zip_url, "-o", "package.zip")
	err = cli53_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cli53_bin_url := "https://github.com/barnybug/cli53/archive/refs/tags/0.8.22.bin"
	cli53_cmd_bin := exec.Command("curl", "-L", cli53_bin_url, "-o", "binary.bin")
	err = cli53_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cli53_src_url := "https://github.com/barnybug/cli53/archive/refs/tags/0.8.22.src.tar.gz"
	cli53_cmd_src := exec.Command("curl", "-L", cli53_src_url, "-o", "source.tar.gz")
	err = cli53_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cli53_cmd_direct := exec.Command("./binary")
	err = cli53_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
