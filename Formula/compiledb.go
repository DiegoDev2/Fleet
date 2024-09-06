package main

// Compiledb - Generate a Clang compilation database for Make-based build systems
// Homepage: https://github.com/nickdiego/compiledb

import (
	"fmt"
	
	"os/exec"
)

func installCompiledb() {
	// Método 1: Descargar y extraer .tar.gz
	compiledb_tar_url := "https://files.pythonhosted.org/packages/76/62/30fb04404b1d4a454f414f792553d142e8acc5da27fddcce911fff0fe570/compiledb-0.10.1.tar.gz"
	compiledb_cmd_tar := exec.Command("curl", "-L", compiledb_tar_url, "-o", "package.tar.gz")
	err := compiledb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	compiledb_zip_url := "https://files.pythonhosted.org/packages/76/62/30fb04404b1d4a454f414f792553d142e8acc5da27fddcce911fff0fe570/compiledb-0.10.1.zip"
	compiledb_cmd_zip := exec.Command("curl", "-L", compiledb_zip_url, "-o", "package.zip")
	err = compiledb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	compiledb_bin_url := "https://files.pythonhosted.org/packages/76/62/30fb04404b1d4a454f414f792553d142e8acc5da27fddcce911fff0fe570/compiledb-0.10.1.bin"
	compiledb_cmd_bin := exec.Command("curl", "-L", compiledb_bin_url, "-o", "binary.bin")
	err = compiledb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	compiledb_src_url := "https://files.pythonhosted.org/packages/76/62/30fb04404b1d4a454f414f792553d142e8acc5da27fddcce911fff0fe570/compiledb-0.10.1.src.tar.gz"
	compiledb_cmd_src := exec.Command("curl", "-L", compiledb_src_url, "-o", "source.tar.gz")
	err = compiledb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	compiledb_cmd_direct := exec.Command("./binary")
	err = compiledb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
