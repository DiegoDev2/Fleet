package main

// KitchenCompletion - Bash completion for Kitchen
// Homepage: https://github.com/MarkBorcherding/test-kitchen-bash-completion

import (
	"fmt"
	
	"os/exec"
)

func installKitchenCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	kitchencompletion_tar_url := "https://github.com/MarkBorcherding/test-kitchen-bash-completion/archive/refs/tags/v1.0.0.tar.gz"
	kitchencompletion_cmd_tar := exec.Command("curl", "-L", kitchencompletion_tar_url, "-o", "package.tar.gz")
	err := kitchencompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kitchencompletion_zip_url := "https://github.com/MarkBorcherding/test-kitchen-bash-completion/archive/refs/tags/v1.0.0.zip"
	kitchencompletion_cmd_zip := exec.Command("curl", "-L", kitchencompletion_zip_url, "-o", "package.zip")
	err = kitchencompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kitchencompletion_bin_url := "https://github.com/MarkBorcherding/test-kitchen-bash-completion/archive/refs/tags/v1.0.0.bin"
	kitchencompletion_cmd_bin := exec.Command("curl", "-L", kitchencompletion_bin_url, "-o", "binary.bin")
	err = kitchencompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kitchencompletion_src_url := "https://github.com/MarkBorcherding/test-kitchen-bash-completion/archive/refs/tags/v1.0.0.src.tar.gz"
	kitchencompletion_cmd_src := exec.Command("curl", "-L", kitchencompletion_src_url, "-o", "source.tar.gz")
	err = kitchencompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kitchencompletion_cmd_direct := exec.Command("./binary")
	err = kitchencompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
