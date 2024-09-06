package main

// CheckJsonschema - JSON Schema CLI
// Homepage: https://github.com/python-jsonschema/check-jsonschema

import (
	"fmt"
	
	"os/exec"
)

func installCheckJsonschema() {
	// Método 1: Descargar y extraer .tar.gz
	checkjsonschema_tar_url := "https://files.pythonhosted.org/packages/3a/63/46ada3bfef271aa4c5ccc86a01a2ca521bfa3a0d7b083dbe76e8adcc460c/check_jsonschema-0.29.2.tar.gz"
	checkjsonschema_cmd_tar := exec.Command("curl", "-L", checkjsonschema_tar_url, "-o", "package.tar.gz")
	err := checkjsonschema_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	checkjsonschema_zip_url := "https://files.pythonhosted.org/packages/3a/63/46ada3bfef271aa4c5ccc86a01a2ca521bfa3a0d7b083dbe76e8adcc460c/check_jsonschema-0.29.2.zip"
	checkjsonschema_cmd_zip := exec.Command("curl", "-L", checkjsonschema_zip_url, "-o", "package.zip")
	err = checkjsonschema_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	checkjsonschema_bin_url := "https://files.pythonhosted.org/packages/3a/63/46ada3bfef271aa4c5ccc86a01a2ca521bfa3a0d7b083dbe76e8adcc460c/check_jsonschema-0.29.2.bin"
	checkjsonschema_cmd_bin := exec.Command("curl", "-L", checkjsonschema_bin_url, "-o", "binary.bin")
	err = checkjsonschema_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	checkjsonschema_src_url := "https://files.pythonhosted.org/packages/3a/63/46ada3bfef271aa4c5ccc86a01a2ca521bfa3a0d7b083dbe76e8adcc460c/check_jsonschema-0.29.2.src.tar.gz"
	checkjsonschema_cmd_src := exec.Command("curl", "-L", checkjsonschema_src_url, "-o", "source.tar.gz")
	err = checkjsonschema_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	checkjsonschema_cmd_direct := exec.Command("./binary")
	err = checkjsonschema_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
