package main

// Mongosh - MongoDB Shell to connect, configure, query, and work with your MongoDB database
// Homepage: https://github.com/mongodb-js/mongosh

import (
	"fmt"
	
	"os/exec"
)

func installMongosh() {
	// Método 1: Descargar y extraer .tar.gz
	mongosh_tar_url := "https://registry.npmjs.org/@mongosh/cli-repl/-/cli-repl-2.3.1.tgz"
	mongosh_cmd_tar := exec.Command("curl", "-L", mongosh_tar_url, "-o", "package.tar.gz")
	err := mongosh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mongosh_zip_url := "https://registry.npmjs.org/@mongosh/cli-repl/-/cli-repl-2.3.1.tgz"
	mongosh_cmd_zip := exec.Command("curl", "-L", mongosh_zip_url, "-o", "package.zip")
	err = mongosh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mongosh_bin_url := "https://registry.npmjs.org/@mongosh/cli-repl/-/cli-repl-2.3.1.tgz"
	mongosh_cmd_bin := exec.Command("curl", "-L", mongosh_bin_url, "-o", "binary.bin")
	err = mongosh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mongosh_src_url := "https://registry.npmjs.org/@mongosh/cli-repl/-/cli-repl-2.3.1.tgz"
	mongosh_cmd_src := exec.Command("curl", "-L", mongosh_src_url, "-o", "source.tar.gz")
	err = mongosh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mongosh_cmd_direct := exec.Command("./binary")
	err = mongosh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
