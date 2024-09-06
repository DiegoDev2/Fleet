package main

// Opa - Open source, general-purpose policy engine
// Homepage: https://www.openpolicyagent.org

import (
	"fmt"
	
	"os/exec"
)

func installOpa() {
	// Método 1: Descargar y extraer .tar.gz
	opa_tar_url := "https://github.com/open-policy-agent/opa/archive/refs/tags/v0.68.0.tar.gz"
	opa_cmd_tar := exec.Command("curl", "-L", opa_tar_url, "-o", "package.tar.gz")
	err := opa_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opa_zip_url := "https://github.com/open-policy-agent/opa/archive/refs/tags/v0.68.0.zip"
	opa_cmd_zip := exec.Command("curl", "-L", opa_zip_url, "-o", "package.zip")
	err = opa_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opa_bin_url := "https://github.com/open-policy-agent/opa/archive/refs/tags/v0.68.0.bin"
	opa_cmd_bin := exec.Command("curl", "-L", opa_bin_url, "-o", "binary.bin")
	err = opa_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opa_src_url := "https://github.com/open-policy-agent/opa/archive/refs/tags/v0.68.0.src.tar.gz"
	opa_cmd_src := exec.Command("curl", "-L", opa_src_url, "-o", "source.tar.gz")
	err = opa_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opa_cmd_direct := exec.Command("./binary")
	err = opa_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
