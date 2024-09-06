package main

// Trino - Distributed SQL query engine for big data
// Homepage: https://trino.io

import (
	"fmt"
	
	"os/exec"
)

func installTrino() {
	// Método 1: Descargar y extraer .tar.gz
	trino_tar_url := "https://search.maven.org/remotecontent?filepath=io/trino/trino-server/455/trino-server-455.tar.gz"
	trino_cmd_tar := exec.Command("curl", "-L", trino_tar_url, "-o", "package.tar.gz")
	err := trino_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trino_zip_url := "https://search.maven.org/remotecontent?filepath=io/trino/trino-server/455/trino-server-455.zip"
	trino_cmd_zip := exec.Command("curl", "-L", trino_zip_url, "-o", "package.zip")
	err = trino_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trino_bin_url := "https://search.maven.org/remotecontent?filepath=io/trino/trino-server/455/trino-server-455.bin"
	trino_cmd_bin := exec.Command("curl", "-L", trino_bin_url, "-o", "binary.bin")
	err = trino_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trino_src_url := "https://search.maven.org/remotecontent?filepath=io/trino/trino-server/455/trino-server-455.src.tar.gz"
	trino_cmd_src := exec.Command("curl", "-L", trino_src_url, "-o", "source.tar.gz")
	err = trino_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trino_cmd_direct := exec.Command("./binary")
	err = trino_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-tar")
	exec.Command("latte", "install", "gnu-tar").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
