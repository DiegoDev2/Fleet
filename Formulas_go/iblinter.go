package main

// Iblinter - Linter tool for Interface Builder
// Homepage: https://github.com/IBDecodable/IBLinter

import (
	"fmt"
	
	"os/exec"
)

func installIblinter() {
	// Método 1: Descargar y extraer .tar.gz
	iblinter_tar_url := "https://github.com/IBDecodable/IBLinter/archive/refs/tags/0.5.0.tar.gz"
	iblinter_cmd_tar := exec.Command("curl", "-L", iblinter_tar_url, "-o", "package.tar.gz")
	err := iblinter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iblinter_zip_url := "https://github.com/IBDecodable/IBLinter/archive/refs/tags/0.5.0.zip"
	iblinter_cmd_zip := exec.Command("curl", "-L", iblinter_zip_url, "-o", "package.zip")
	err = iblinter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iblinter_bin_url := "https://github.com/IBDecodable/IBLinter/archive/refs/tags/0.5.0.bin"
	iblinter_cmd_bin := exec.Command("curl", "-L", iblinter_bin_url, "-o", "binary.bin")
	err = iblinter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iblinter_src_url := "https://github.com/IBDecodable/IBLinter/archive/refs/tags/0.5.0.src.tar.gz"
	iblinter_cmd_src := exec.Command("curl", "-L", iblinter_src_url, "-o", "source.tar.gz")
	err = iblinter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iblinter_cmd_direct := exec.Command("./binary")
	err = iblinter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
