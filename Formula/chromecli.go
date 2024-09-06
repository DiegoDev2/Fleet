package main

// ChromeCli - Control Google Chrome from the command-line
// Homepage: https://github.com/prasmussen/chrome-cli

import (
	"fmt"
	
	"os/exec"
)

func installChromeCli() {
	// Método 1: Descargar y extraer .tar.gz
	chromecli_tar_url := "https://github.com/prasmussen/chrome-cli/archive/refs/tags/1.10.0.tar.gz"
	chromecli_cmd_tar := exec.Command("curl", "-L", chromecli_tar_url, "-o", "package.tar.gz")
	err := chromecli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chromecli_zip_url := "https://github.com/prasmussen/chrome-cli/archive/refs/tags/1.10.0.zip"
	chromecli_cmd_zip := exec.Command("curl", "-L", chromecli_zip_url, "-o", "package.zip")
	err = chromecli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chromecli_bin_url := "https://github.com/prasmussen/chrome-cli/archive/refs/tags/1.10.0.bin"
	chromecli_cmd_bin := exec.Command("curl", "-L", chromecli_bin_url, "-o", "binary.bin")
	err = chromecli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chromecli_src_url := "https://github.com/prasmussen/chrome-cli/archive/refs/tags/1.10.0.src.tar.gz"
	chromecli_cmd_src := exec.Command("curl", "-L", chromecli_src_url, "-o", "source.tar.gz")
	err = chromecli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chromecli_cmd_direct := exec.Command("./binary")
	err = chromecli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
