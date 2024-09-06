package main

// Cli11 - Simple and intuitive command-line parser for C++11
// Homepage: https://cliutils.github.io/CLI11/book/

import (
	"fmt"
	
	"os/exec"
)

func installCli11() {
	// Método 1: Descargar y extraer .tar.gz
	cli11_tar_url := "https://github.com/CLIUtils/CLI11/archive/refs/tags/v2.4.2.tar.gz"
	cli11_cmd_tar := exec.Command("curl", "-L", cli11_tar_url, "-o", "package.tar.gz")
	err := cli11_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cli11_zip_url := "https://github.com/CLIUtils/CLI11/archive/refs/tags/v2.4.2.zip"
	cli11_cmd_zip := exec.Command("curl", "-L", cli11_zip_url, "-o", "package.zip")
	err = cli11_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cli11_bin_url := "https://github.com/CLIUtils/CLI11/archive/refs/tags/v2.4.2.bin"
	cli11_cmd_bin := exec.Command("curl", "-L", cli11_bin_url, "-o", "binary.bin")
	err = cli11_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cli11_src_url := "https://github.com/CLIUtils/CLI11/archive/refs/tags/v2.4.2.src.tar.gz"
	cli11_cmd_src := exec.Command("curl", "-L", cli11_src_url, "-o", "source.tar.gz")
	err = cli11_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cli11_cmd_direct := exec.Command("./binary")
	err = cli11_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
