package main

// Prestodb - Distributed SQL query engine for big data
// Homepage: https://prestodb.io

import (
	"fmt"
	
	"os/exec"
)

func installPrestodb() {
	// Método 1: Descargar y extraer .tar.gz
	prestodb_tar_url := "https://search.maven.org/remotecontent?filepath=com/facebook/presto/presto-server/0.288.1/presto-server-0.288.1.tar.gz"
	prestodb_cmd_tar := exec.Command("curl", "-L", prestodb_tar_url, "-o", "package.tar.gz")
	err := prestodb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	prestodb_zip_url := "https://search.maven.org/remotecontent?filepath=com/facebook/presto/presto-server/0.288.1/presto-server-0.288.1.zip"
	prestodb_cmd_zip := exec.Command("curl", "-L", prestodb_zip_url, "-o", "package.zip")
	err = prestodb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	prestodb_bin_url := "https://search.maven.org/remotecontent?filepath=com/facebook/presto/presto-server/0.288.1/presto-server-0.288.1.bin"
	prestodb_cmd_bin := exec.Command("curl", "-L", prestodb_bin_url, "-o", "binary.bin")
	err = prestodb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	prestodb_src_url := "https://search.maven.org/remotecontent?filepath=com/facebook/presto/presto-server/0.288.1/presto-server-0.288.1.src.tar.gz"
	prestodb_cmd_src := exec.Command("curl", "-L", prestodb_src_url, "-o", "source.tar.gz")
	err = prestodb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	prestodb_cmd_direct := exec.Command("./binary")
	err = prestodb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
