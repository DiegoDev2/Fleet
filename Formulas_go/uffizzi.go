package main

// Uffizzi - Self-serve developer platforms in minutes, not months with k8s virtual clusters
// Homepage: https://uffizzi.com

import (
	"fmt"
	
	"os/exec"
)

func installUffizzi() {
	// Método 1: Descargar y extraer .tar.gz
	uffizzi_tar_url := "https://github.com/UffizziCloud/uffizzi_cli/archive/refs/tags/v2.4.11.tar.gz"
	uffizzi_cmd_tar := exec.Command("curl", "-L", uffizzi_tar_url, "-o", "package.tar.gz")
	err := uffizzi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uffizzi_zip_url := "https://github.com/UffizziCloud/uffizzi_cli/archive/refs/tags/v2.4.11.zip"
	uffizzi_cmd_zip := exec.Command("curl", "-L", uffizzi_zip_url, "-o", "package.zip")
	err = uffizzi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uffizzi_bin_url := "https://github.com/UffizziCloud/uffizzi_cli/archive/refs/tags/v2.4.11.bin"
	uffizzi_cmd_bin := exec.Command("curl", "-L", uffizzi_bin_url, "-o", "binary.bin")
	err = uffizzi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uffizzi_src_url := "https://github.com/UffizziCloud/uffizzi_cli/archive/refs/tags/v2.4.11.src.tar.gz"
	uffizzi_cmd_src := exec.Command("curl", "-L", uffizzi_src_url, "-o", "source.tar.gz")
	err = uffizzi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uffizzi_cmd_direct := exec.Command("./binary")
	err = uffizzi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
	fmt.Println("Instalando dependencia: skaffold")
exec.Command("latte", "install", "skaffold")
}
