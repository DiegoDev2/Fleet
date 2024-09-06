package main

// Hermit - Manages isolated, self-bootstrapping sets of tools in software projects
// Homepage: https://cashapp.github.io/hermit

import (
	"fmt"
	
	"os/exec"
)

func installHermit() {
	// Método 1: Descargar y extraer .tar.gz
	hermit_tar_url := "https://github.com/cashapp/hermit/archive/refs/tags/v0.40.0.tar.gz"
	hermit_cmd_tar := exec.Command("curl", "-L", hermit_tar_url, "-o", "package.tar.gz")
	err := hermit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hermit_zip_url := "https://github.com/cashapp/hermit/archive/refs/tags/v0.40.0.zip"
	hermit_cmd_zip := exec.Command("curl", "-L", hermit_zip_url, "-o", "package.zip")
	err = hermit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hermit_bin_url := "https://github.com/cashapp/hermit/archive/refs/tags/v0.40.0.bin"
	hermit_cmd_bin := exec.Command("curl", "-L", hermit_bin_url, "-o", "binary.bin")
	err = hermit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hermit_src_url := "https://github.com/cashapp/hermit/archive/refs/tags/v0.40.0.src.tar.gz"
	hermit_cmd_src := exec.Command("curl", "-L", hermit_src_url, "-o", "source.tar.gz")
	err = hermit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hermit_cmd_direct := exec.Command("./binary")
	err = hermit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
