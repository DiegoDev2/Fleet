package main

// Testscript - Integration tests for command-line applications in .txtar format
// Homepage: https://github.com/rogpeppe/go-internal/tree/master/cmd/testscript

import (
	"fmt"
	
	"os/exec"
)

func installTestscript() {
	// Método 1: Descargar y extraer .tar.gz
	testscript_tar_url := "https://github.com/rogpeppe/go-internal/archive/refs/tags/v1.12.0.tar.gz"
	testscript_cmd_tar := exec.Command("curl", "-L", testscript_tar_url, "-o", "package.tar.gz")
	err := testscript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	testscript_zip_url := "https://github.com/rogpeppe/go-internal/archive/refs/tags/v1.12.0.zip"
	testscript_cmd_zip := exec.Command("curl", "-L", testscript_zip_url, "-o", "package.zip")
	err = testscript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	testscript_bin_url := "https://github.com/rogpeppe/go-internal/archive/refs/tags/v1.12.0.bin"
	testscript_cmd_bin := exec.Command("curl", "-L", testscript_bin_url, "-o", "binary.bin")
	err = testscript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	testscript_src_url := "https://github.com/rogpeppe/go-internal/archive/refs/tags/v1.12.0.src.tar.gz"
	testscript_cmd_src := exec.Command("curl", "-L", testscript_src_url, "-o", "source.tar.gz")
	err = testscript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	testscript_cmd_direct := exec.Command("./binary")
	err = testscript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
