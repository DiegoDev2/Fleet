package main

// Cliam - Cloud agnostic IAM permissions enumerator
// Homepage: https://github.com/securisec/cliam

import (
	"fmt"
	
	"os/exec"
)

func installCliam() {
	// Método 1: Descargar y extraer .tar.gz
	cliam_tar_url := "https://github.com/securisec/cliam/archive/refs/tags/2.2.0.tar.gz"
	cliam_cmd_tar := exec.Command("curl", "-L", cliam_tar_url, "-o", "package.tar.gz")
	err := cliam_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cliam_zip_url := "https://github.com/securisec/cliam/archive/refs/tags/2.2.0.zip"
	cliam_cmd_zip := exec.Command("curl", "-L", cliam_zip_url, "-o", "package.zip")
	err = cliam_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cliam_bin_url := "https://github.com/securisec/cliam/archive/refs/tags/2.2.0.bin"
	cliam_cmd_bin := exec.Command("curl", "-L", cliam_bin_url, "-o", "binary.bin")
	err = cliam_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cliam_src_url := "https://github.com/securisec/cliam/archive/refs/tags/2.2.0.src.tar.gz"
	cliam_cmd_src := exec.Command("curl", "-L", cliam_src_url, "-o", "source.tar.gz")
	err = cliam_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cliam_cmd_direct := exec.Command("./binary")
	err = cliam_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
