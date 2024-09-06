package main

// PythonYq - Command-line YAML and XML processor that wraps jq
// Homepage: https://kislyuk.github.io/yq/

import (
	"fmt"
	
	"os/exec"
)

func installPythonYq() {
	// Método 1: Descargar y extraer .tar.gz
	pythonyq_tar_url := "https://files.pythonhosted.org/packages/38/6a/eb9721ed0929d0f55d167c2222d288b529723afbef0a07ed7aa6cca72380/yq-3.4.3.tar.gz"
	pythonyq_cmd_tar := exec.Command("curl", "-L", pythonyq_tar_url, "-o", "package.tar.gz")
	err := pythonyq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonyq_zip_url := "https://files.pythonhosted.org/packages/38/6a/eb9721ed0929d0f55d167c2222d288b529723afbef0a07ed7aa6cca72380/yq-3.4.3.zip"
	pythonyq_cmd_zip := exec.Command("curl", "-L", pythonyq_zip_url, "-o", "package.zip")
	err = pythonyq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonyq_bin_url := "https://files.pythonhosted.org/packages/38/6a/eb9721ed0929d0f55d167c2222d288b529723afbef0a07ed7aa6cca72380/yq-3.4.3.bin"
	pythonyq_cmd_bin := exec.Command("curl", "-L", pythonyq_bin_url, "-o", "binary.bin")
	err = pythonyq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonyq_src_url := "https://files.pythonhosted.org/packages/38/6a/eb9721ed0929d0f55d167c2222d288b529723afbef0a07ed7aa6cca72380/yq-3.4.3.src.tar.gz"
	pythonyq_cmd_src := exec.Command("curl", "-L", pythonyq_src_url, "-o", "source.tar.gz")
	err = pythonyq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonyq_cmd_direct := exec.Command("./binary")
	err = pythonyq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jq")
	exec.Command("latte", "install", "jq").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
