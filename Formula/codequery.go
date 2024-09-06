package main

// Codequery - Code-understanding, code-browsing or code-search tool
// Homepage: https://ruben2020.github.io/codequery/

import (
	"fmt"
	
	"os/exec"
)

func installCodequery() {
	// Método 1: Descargar y extraer .tar.gz
	codequery_tar_url := "https://github.com/ruben2020/codequery/archive/refs/tags/v1.0.0.tar.gz"
	codequery_cmd_tar := exec.Command("curl", "-L", codequery_tar_url, "-o", "package.tar.gz")
	err := codequery_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	codequery_zip_url := "https://github.com/ruben2020/codequery/archive/refs/tags/v1.0.0.zip"
	codequery_cmd_zip := exec.Command("curl", "-L", codequery_zip_url, "-o", "package.zip")
	err = codequery_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	codequery_bin_url := "https://github.com/ruben2020/codequery/archive/refs/tags/v1.0.0.bin"
	codequery_cmd_bin := exec.Command("curl", "-L", codequery_bin_url, "-o", "binary.bin")
	err = codequery_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	codequery_src_url := "https://github.com/ruben2020/codequery/archive/refs/tags/v1.0.0.src.tar.gz"
	codequery_cmd_src := exec.Command("curl", "-L", codequery_src_url, "-o", "source.tar.gz")
	err = codequery_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	codequery_cmd_direct := exec.Command("./binary")
	err = codequery_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
}
