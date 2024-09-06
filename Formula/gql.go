package main

// Gql - Git Query language is a SQL like language to perform queries on .git files
// Homepage: https://github.com/AmrDeveloper/GQL

import (
	"fmt"
	
	"os/exec"
)

func installGql() {
	// Método 1: Descargar y extraer .tar.gz
	gql_tar_url := "https://github.com/AmrDeveloper/GQL/archive/refs/tags/0.26.0.tar.gz"
	gql_cmd_tar := exec.Command("curl", "-L", gql_tar_url, "-o", "package.tar.gz")
	err := gql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gql_zip_url := "https://github.com/AmrDeveloper/GQL/archive/refs/tags/0.26.0.zip"
	gql_cmd_zip := exec.Command("curl", "-L", gql_zip_url, "-o", "package.zip")
	err = gql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gql_bin_url := "https://github.com/AmrDeveloper/GQL/archive/refs/tags/0.26.0.bin"
	gql_cmd_bin := exec.Command("curl", "-L", gql_bin_url, "-o", "binary.bin")
	err = gql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gql_src_url := "https://github.com/AmrDeveloper/GQL/archive/refs/tags/0.26.0.src.tar.gz"
	gql_cmd_src := exec.Command("curl", "-L", gql_src_url, "-o", "source.tar.gz")
	err = gql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gql_cmd_direct := exec.Command("./binary")
	err = gql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
