package main

// Hcl2json - Convert HCL2 to JSON
// Homepage: https://github.com/tmccombs/hcl2json

import (
	"fmt"
	
	"os/exec"
)

func installHcl2json() {
	// Método 1: Descargar y extraer .tar.gz
	hcl2json_tar_url := "https://github.com/tmccombs/hcl2json/archive/refs/tags/v0.6.4.tar.gz"
	hcl2json_cmd_tar := exec.Command("curl", "-L", hcl2json_tar_url, "-o", "package.tar.gz")
	err := hcl2json_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hcl2json_zip_url := "https://github.com/tmccombs/hcl2json/archive/refs/tags/v0.6.4.zip"
	hcl2json_cmd_zip := exec.Command("curl", "-L", hcl2json_zip_url, "-o", "package.zip")
	err = hcl2json_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hcl2json_bin_url := "https://github.com/tmccombs/hcl2json/archive/refs/tags/v0.6.4.bin"
	hcl2json_cmd_bin := exec.Command("curl", "-L", hcl2json_bin_url, "-o", "binary.bin")
	err = hcl2json_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hcl2json_src_url := "https://github.com/tmccombs/hcl2json/archive/refs/tags/v0.6.4.src.tar.gz"
	hcl2json_cmd_src := exec.Command("curl", "-L", hcl2json_src_url, "-o", "source.tar.gz")
	err = hcl2json_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hcl2json_cmd_direct := exec.Command("./binary")
	err = hcl2json_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
