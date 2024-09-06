package main

// Gh - GitHub command-line tool
// Homepage: https://cli.github.com/

import (
	"fmt"
	
	"os/exec"
)

func installGh() {
	// Método 1: Descargar y extraer .tar.gz
	gh_tar_url := "https://github.com/cli/cli/archive/refs/tags/v2.55.0.tar.gz"
	gh_cmd_tar := exec.Command("curl", "-L", gh_tar_url, "-o", "package.tar.gz")
	err := gh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gh_zip_url := "https://github.com/cli/cli/archive/refs/tags/v2.55.0.zip"
	gh_cmd_zip := exec.Command("curl", "-L", gh_zip_url, "-o", "package.zip")
	err = gh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gh_bin_url := "https://github.com/cli/cli/archive/refs/tags/v2.55.0.bin"
	gh_cmd_bin := exec.Command("curl", "-L", gh_bin_url, "-o", "binary.bin")
	err = gh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gh_src_url := "https://github.com/cli/cli/archive/refs/tags/v2.55.0.src.tar.gz"
	gh_cmd_src := exec.Command("curl", "-L", gh_src_url, "-o", "source.tar.gz")
	err = gh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gh_cmd_direct := exec.Command("./binary")
	err = gh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
