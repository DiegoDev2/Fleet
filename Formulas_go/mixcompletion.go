package main

// MixCompletion - Elixir Mix completion plus shortcuts/colors
// Homepage: https://github.com/davidhq/mix-power-completion

import (
	"fmt"
	
	"os/exec"
)

func installMixCompletion() {
	// Método 1: Descargar y extraer .tar.gz
	mixcompletion_tar_url := "https://github.com/davidhq/mix-power-completion/archive/refs/tags/0.8.2.tar.gz"
	mixcompletion_cmd_tar := exec.Command("curl", "-L", mixcompletion_tar_url, "-o", "package.tar.gz")
	err := mixcompletion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mixcompletion_zip_url := "https://github.com/davidhq/mix-power-completion/archive/refs/tags/0.8.2.zip"
	mixcompletion_cmd_zip := exec.Command("curl", "-L", mixcompletion_zip_url, "-o", "package.zip")
	err = mixcompletion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mixcompletion_bin_url := "https://github.com/davidhq/mix-power-completion/archive/refs/tags/0.8.2.bin"
	mixcompletion_cmd_bin := exec.Command("curl", "-L", mixcompletion_bin_url, "-o", "binary.bin")
	err = mixcompletion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mixcompletion_src_url := "https://github.com/davidhq/mix-power-completion/archive/refs/tags/0.8.2.src.tar.gz"
	mixcompletion_cmd_src := exec.Command("curl", "-L", mixcompletion_src_url, "-o", "source.tar.gz")
	err = mixcompletion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mixcompletion_cmd_direct := exec.Command("./binary")
	err = mixcompletion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
