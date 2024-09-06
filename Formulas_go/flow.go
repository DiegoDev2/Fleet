package main

// Flow - Static type checker for JavaScript
// Homepage: https://flow.org/

import (
	"fmt"
	
	"os/exec"
)

func installFlow() {
	// Método 1: Descargar y extraer .tar.gz
	flow_tar_url := "https://github.com/facebook/flow/archive/refs/tags/v0.245.2.tar.gz"
	flow_cmd_tar := exec.Command("curl", "-L", flow_tar_url, "-o", "package.tar.gz")
	err := flow_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flow_zip_url := "https://github.com/facebook/flow/archive/refs/tags/v0.245.2.zip"
	flow_cmd_zip := exec.Command("curl", "-L", flow_zip_url, "-o", "package.zip")
	err = flow_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flow_bin_url := "https://github.com/facebook/flow/archive/refs/tags/v0.245.2.bin"
	flow_cmd_bin := exec.Command("curl", "-L", flow_bin_url, "-o", "binary.bin")
	err = flow_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flow_src_url := "https://github.com/facebook/flow/archive/refs/tags/v0.245.2.src.tar.gz"
	flow_cmd_src := exec.Command("curl", "-L", flow_src_url, "-o", "source.tar.gz")
	err = flow_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flow_cmd_direct := exec.Command("./binary")
	err = flow_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ocaml")
exec.Command("latte", "install", "ocaml")
	fmt.Println("Instalando dependencia: opam")
exec.Command("latte", "install", "opam")
}
