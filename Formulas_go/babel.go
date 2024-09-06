package main

// Babel - Compiler for writing next generation JavaScript
// Homepage: https://babeljs.io/

import (
	"fmt"
	
	"os/exec"
)

func installBabel() {
	// Método 1: Descargar y extraer .tar.gz
	babel_tar_url := "https://registry.npmjs.org/@babel/core/-/core-7.25.2.tgz"
	babel_cmd_tar := exec.Command("curl", "-L", babel_tar_url, "-o", "package.tar.gz")
	err := babel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	babel_zip_url := "https://registry.npmjs.org/@babel/core/-/core-7.25.2.tgz"
	babel_cmd_zip := exec.Command("curl", "-L", babel_zip_url, "-o", "package.zip")
	err = babel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	babel_bin_url := "https://registry.npmjs.org/@babel/core/-/core-7.25.2.tgz"
	babel_cmd_bin := exec.Command("curl", "-L", babel_bin_url, "-o", "binary.bin")
	err = babel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	babel_src_url := "https://registry.npmjs.org/@babel/core/-/core-7.25.2.tgz"
	babel_cmd_src := exec.Command("curl", "-L", babel_src_url, "-o", "source.tar.gz")
	err = babel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	babel_cmd_direct := exec.Command("./binary")
	err = babel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
