package main

// Naml - Convert Kubernetes YAML to Golang
// Homepage: https://github.com/krisnova/naml

import (
	"fmt"
	
	"os/exec"
)

func installNaml() {
	// Método 1: Descargar y extraer .tar.gz
	naml_tar_url := "https://github.com/krisnova/naml/archive/refs/tags/v1.0.3.tar.gz"
	naml_cmd_tar := exec.Command("curl", "-L", naml_tar_url, "-o", "package.tar.gz")
	err := naml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	naml_zip_url := "https://github.com/krisnova/naml/archive/refs/tags/v1.0.3.zip"
	naml_cmd_zip := exec.Command("curl", "-L", naml_zip_url, "-o", "package.zip")
	err = naml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	naml_bin_url := "https://github.com/krisnova/naml/archive/refs/tags/v1.0.3.bin"
	naml_cmd_bin := exec.Command("curl", "-L", naml_bin_url, "-o", "binary.bin")
	err = naml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	naml_src_url := "https://github.com/krisnova/naml/archive/refs/tags/v1.0.3.src.tar.gz"
	naml_cmd_src := exec.Command("curl", "-L", naml_src_url, "-o", "source.tar.gz")
	err = naml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	naml_cmd_direct := exec.Command("./binary")
	err = naml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
