package main

// WoodpeckerCli - CLI client for the Woodpecker Continuous Integration server
// Homepage: https://woodpecker-ci.org/

import (
	"fmt"
	
	"os/exec"
)

func installWoodpeckerCli() {
	// Método 1: Descargar y extraer .tar.gz
	woodpeckercli_tar_url := "https://github.com/woodpecker-ci/woodpecker/archive/refs/tags/v2.7.0.tar.gz"
	woodpeckercli_cmd_tar := exec.Command("curl", "-L", woodpeckercli_tar_url, "-o", "package.tar.gz")
	err := woodpeckercli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	woodpeckercli_zip_url := "https://github.com/woodpecker-ci/woodpecker/archive/refs/tags/v2.7.0.zip"
	woodpeckercli_cmd_zip := exec.Command("curl", "-L", woodpeckercli_zip_url, "-o", "package.zip")
	err = woodpeckercli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	woodpeckercli_bin_url := "https://github.com/woodpecker-ci/woodpecker/archive/refs/tags/v2.7.0.bin"
	woodpeckercli_cmd_bin := exec.Command("curl", "-L", woodpeckercli_bin_url, "-o", "binary.bin")
	err = woodpeckercli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	woodpeckercli_src_url := "https://github.com/woodpecker-ci/woodpecker/archive/refs/tags/v2.7.0.src.tar.gz"
	woodpeckercli_cmd_src := exec.Command("curl", "-L", woodpeckercli_src_url, "-o", "source.tar.gz")
	err = woodpeckercli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	woodpeckercli_cmd_direct := exec.Command("./binary")
	err = woodpeckercli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
