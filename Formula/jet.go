package main

// Jet - Type safe SQL builder with code generation and auto query result data mapping
// Homepage: https://github.com/go-jet/jet

import (
	"fmt"
	
	"os/exec"
)

func installJet() {
	// Método 1: Descargar y extraer .tar.gz
	jet_tar_url := "https://github.com/go-jet/jet/archive/refs/tags/v2.11.1.tar.gz"
	jet_cmd_tar := exec.Command("curl", "-L", jet_tar_url, "-o", "package.tar.gz")
	err := jet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jet_zip_url := "https://github.com/go-jet/jet/archive/refs/tags/v2.11.1.zip"
	jet_cmd_zip := exec.Command("curl", "-L", jet_zip_url, "-o", "package.zip")
	err = jet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jet_bin_url := "https://github.com/go-jet/jet/archive/refs/tags/v2.11.1.bin"
	jet_cmd_bin := exec.Command("curl", "-L", jet_bin_url, "-o", "binary.bin")
	err = jet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jet_src_url := "https://github.com/go-jet/jet/archive/refs/tags/v2.11.1.src.tar.gz"
	jet_cmd_src := exec.Command("curl", "-L", jet_src_url, "-o", "source.tar.gz")
	err = jet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jet_cmd_direct := exec.Command("./binary")
	err = jet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
