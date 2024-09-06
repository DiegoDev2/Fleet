package main

// Typescript - Language for application scale JavaScript development
// Homepage: https://www.typescriptlang.org/

import (
	"fmt"
	
	"os/exec"
)

func installTypescript() {
	// Método 1: Descargar y extraer .tar.gz
	typescript_tar_url := "https://registry.npmjs.org/typescript/-/typescript-5.5.4.tgz"
	typescript_cmd_tar := exec.Command("curl", "-L", typescript_tar_url, "-o", "package.tar.gz")
	err := typescript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	typescript_zip_url := "https://registry.npmjs.org/typescript/-/typescript-5.5.4.tgz"
	typescript_cmd_zip := exec.Command("curl", "-L", typescript_zip_url, "-o", "package.zip")
	err = typescript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	typescript_bin_url := "https://registry.npmjs.org/typescript/-/typescript-5.5.4.tgz"
	typescript_cmd_bin := exec.Command("curl", "-L", typescript_bin_url, "-o", "binary.bin")
	err = typescript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	typescript_src_url := "https://registry.npmjs.org/typescript/-/typescript-5.5.4.tgz"
	typescript_cmd_src := exec.Command("curl", "-L", typescript_src_url, "-o", "source.tar.gz")
	err = typescript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	typescript_cmd_direct := exec.Command("./binary")
	err = typescript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
