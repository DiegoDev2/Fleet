package main

// Ignite - Build, launch, and maintain any crypto application with Ignite CLI
// Homepage: https://github.com/ignite/cli

import (
	"fmt"
	
	"os/exec"
)

func installIgnite() {
	// Método 1: Descargar y extraer .tar.gz
	ignite_tar_url := "https://github.com/ignite/cli/archive/refs/tags/v28.5.1.tar.gz"
	ignite_cmd_tar := exec.Command("curl", "-L", ignite_tar_url, "-o", "package.tar.gz")
	err := ignite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ignite_zip_url := "https://github.com/ignite/cli/archive/refs/tags/v28.5.1.zip"
	ignite_cmd_zip := exec.Command("curl", "-L", ignite_zip_url, "-o", "package.zip")
	err = ignite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ignite_bin_url := "https://github.com/ignite/cli/archive/refs/tags/v28.5.1.bin"
	ignite_cmd_bin := exec.Command("curl", "-L", ignite_bin_url, "-o", "binary.bin")
	err = ignite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ignite_src_url := "https://github.com/ignite/cli/archive/refs/tags/v28.5.1.src.tar.gz"
	ignite_cmd_src := exec.Command("curl", "-L", ignite_src_url, "-o", "source.tar.gz")
	err = ignite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ignite_cmd_direct := exec.Command("./binary")
	err = ignite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
