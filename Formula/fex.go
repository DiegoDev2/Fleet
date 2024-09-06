package main

// Fex - Powerful field extraction tool
// Homepage: https://www.semicomplete.com/projects/fex/

import (
	"fmt"
	
	"os/exec"
)

func installFex() {
	// Método 1: Descargar y extraer .tar.gz
	fex_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/semicomplete/fex-2.0.0.tar.gz"
	fex_cmd_tar := exec.Command("curl", "-L", fex_tar_url, "-o", "package.tar.gz")
	err := fex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fex_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/semicomplete/fex-2.0.0.zip"
	fex_cmd_zip := exec.Command("curl", "-L", fex_zip_url, "-o", "package.zip")
	err = fex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fex_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/semicomplete/fex-2.0.0.bin"
	fex_cmd_bin := exec.Command("curl", "-L", fex_bin_url, "-o", "binary.bin")
	err = fex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fex_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/semicomplete/fex-2.0.0.src.tar.gz"
	fex_cmd_src := exec.Command("curl", "-L", fex_src_url, "-o", "source.tar.gz")
	err = fex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fex_cmd_direct := exec.Command("./binary")
	err = fex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
