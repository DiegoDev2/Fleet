package main

// Textql - Executes SQL across text files
// Homepage: https://github.com/dinedal/textql

import (
	"fmt"
	
	"os/exec"
)

func installTextql() {
	// Método 1: Descargar y extraer .tar.gz
	textql_tar_url := "https://github.com/dinedal/textql/archive/refs/tags/2.0.3.tar.gz"
	textql_cmd_tar := exec.Command("curl", "-L", textql_tar_url, "-o", "package.tar.gz")
	err := textql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	textql_zip_url := "https://github.com/dinedal/textql/archive/refs/tags/2.0.3.zip"
	textql_cmd_zip := exec.Command("curl", "-L", textql_zip_url, "-o", "package.zip")
	err = textql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	textql_bin_url := "https://github.com/dinedal/textql/archive/refs/tags/2.0.3.bin"
	textql_cmd_bin := exec.Command("curl", "-L", textql_bin_url, "-o", "binary.bin")
	err = textql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	textql_src_url := "https://github.com/dinedal/textql/archive/refs/tags/2.0.3.src.tar.gz"
	textql_cmd_src := exec.Command("curl", "-L", textql_src_url, "-o", "source.tar.gz")
	err = textql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	textql_cmd_direct := exec.Command("./binary")
	err = textql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: glide")
	exec.Command("latte", "install", "glide").Run()
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
