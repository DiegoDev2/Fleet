package main

// Oj - JSON parser and visualization tool
// Homepage: https://github.com/ohler55/ojg

import (
	"fmt"
	
	"os/exec"
)

func installOj() {
	// Método 1: Descargar y extraer .tar.gz
	oj_tar_url := "https://github.com/ohler55/ojg/archive/refs/tags/v1.24.0.tar.gz"
	oj_cmd_tar := exec.Command("curl", "-L", oj_tar_url, "-o", "package.tar.gz")
	err := oj_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oj_zip_url := "https://github.com/ohler55/ojg/archive/refs/tags/v1.24.0.zip"
	oj_cmd_zip := exec.Command("curl", "-L", oj_zip_url, "-o", "package.zip")
	err = oj_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oj_bin_url := "https://github.com/ohler55/ojg/archive/refs/tags/v1.24.0.bin"
	oj_cmd_bin := exec.Command("curl", "-L", oj_bin_url, "-o", "binary.bin")
	err = oj_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oj_src_url := "https://github.com/ohler55/ojg/archive/refs/tags/v1.24.0.src.tar.gz"
	oj_cmd_src := exec.Command("curl", "-L", oj_src_url, "-o", "source.tar.gz")
	err = oj_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oj_cmd_direct := exec.Command("./binary")
	err = oj_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
