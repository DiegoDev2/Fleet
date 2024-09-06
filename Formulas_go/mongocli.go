package main

// Mongocli - MongoDB CLI enables you to manage your MongoDB in the Cloud
// Homepage: https://www.mongodb.com/docs/mongocli/stable/

import (
	"fmt"
	
	"os/exec"
)

func installMongocli() {
	// Método 1: Descargar y extraer .tar.gz
	mongocli_tar_url := "https://github.com/mongodb/mongodb-atlas-cli/archive/refs/tags/mongocli/v2.0.1.tar.gz"
	mongocli_cmd_tar := exec.Command("curl", "-L", mongocli_tar_url, "-o", "package.tar.gz")
	err := mongocli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mongocli_zip_url := "https://github.com/mongodb/mongodb-atlas-cli/archive/refs/tags/mongocli/v2.0.1.zip"
	mongocli_cmd_zip := exec.Command("curl", "-L", mongocli_zip_url, "-o", "package.zip")
	err = mongocli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mongocli_bin_url := "https://github.com/mongodb/mongodb-atlas-cli/archive/refs/tags/mongocli/v2.0.1.bin"
	mongocli_cmd_bin := exec.Command("curl", "-L", mongocli_bin_url, "-o", "binary.bin")
	err = mongocli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mongocli_src_url := "https://github.com/mongodb/mongodb-atlas-cli/archive/refs/tags/mongocli/v2.0.1.src.tar.gz"
	mongocli_cmd_src := exec.Command("curl", "-L", mongocli_src_url, "-o", "source.tar.gz")
	err = mongocli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mongocli_cmd_direct := exec.Command("./binary")
	err = mongocli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
