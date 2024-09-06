package main

// Arangodb - Multi-Model NoSQL Database
// Homepage: https://www.arangodb.com/

import (
	"fmt"
	
	"os/exec"
)

func installArangodb() {
	// Método 1: Descargar y extraer .tar.gz
	arangodb_tar_url := "https://download.arangodb.com/Source/ArangoDB-3.10.4.tar.bz2"
	arangodb_cmd_tar := exec.Command("curl", "-L", arangodb_tar_url, "-o", "package.tar.gz")
	err := arangodb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arangodb_zip_url := "https://download.arangodb.com/Source/ArangoDB-3.10.4.tar.bz2"
	arangodb_cmd_zip := exec.Command("curl", "-L", arangodb_zip_url, "-o", "package.zip")
	err = arangodb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arangodb_bin_url := "https://download.arangodb.com/Source/ArangoDB-3.10.4.tar.bz2"
	arangodb_cmd_bin := exec.Command("curl", "-L", arangodb_bin_url, "-o", "binary.bin")
	err = arangodb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arangodb_src_url := "https://download.arangodb.com/Source/ArangoDB-3.10.4.tar.bz2"
	arangodb_cmd_src := exec.Command("curl", "-L", arangodb_src_url, "-o", "source.tar.gz")
	err = arangodb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arangodb_cmd_direct := exec.Command("./binary")
	err = arangodb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: openssl@1.1")
	exec.Command("latte", "install", "openssl@1.1").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: gcc@10")
	exec.Command("latte", "install", "gcc@10").Run()
}
