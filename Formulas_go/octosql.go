package main

// Octosql - SQL query tool to analyze data from different file formats and databases
// Homepage: https://github.com/cube2222/octosql/

import (
	"fmt"
	
	"os/exec"
)

func installOctosql() {
	// Método 1: Descargar y extraer .tar.gz
	octosql_tar_url := "https://github.com/cube2222/octosql/archive/refs/tags/v0.13.0.tar.gz"
	octosql_cmd_tar := exec.Command("curl", "-L", octosql_tar_url, "-o", "package.tar.gz")
	err := octosql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	octosql_zip_url := "https://github.com/cube2222/octosql/archive/refs/tags/v0.13.0.zip"
	octosql_cmd_zip := exec.Command("curl", "-L", octosql_zip_url, "-o", "package.zip")
	err = octosql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	octosql_bin_url := "https://github.com/cube2222/octosql/archive/refs/tags/v0.13.0.bin"
	octosql_cmd_bin := exec.Command("curl", "-L", octosql_bin_url, "-o", "binary.bin")
	err = octosql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	octosql_src_url := "https://github.com/cube2222/octosql/archive/refs/tags/v0.13.0.src.tar.gz"
	octosql_cmd_src := exec.Command("curl", "-L", octosql_src_url, "-o", "source.tar.gz")
	err = octosql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	octosql_cmd_direct := exec.Command("./binary")
	err = octosql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
