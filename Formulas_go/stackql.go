package main

// Stackql - SQL interface for arbitrary resources with full CRUD support
// Homepage: https://stackql.io/

import (
	"fmt"
	
	"os/exec"
)

func installStackql() {
	// Método 1: Descargar y extraer .tar.gz
	stackql_tar_url := "https://github.com/stackql/stackql/archive/refs/tags/v0.5.734.tar.gz"
	stackql_cmd_tar := exec.Command("curl", "-L", stackql_tar_url, "-o", "package.tar.gz")
	err := stackql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stackql_zip_url := "https://github.com/stackql/stackql/archive/refs/tags/v0.5.734.zip"
	stackql_cmd_zip := exec.Command("curl", "-L", stackql_zip_url, "-o", "package.zip")
	err = stackql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stackql_bin_url := "https://github.com/stackql/stackql/archive/refs/tags/v0.5.734.bin"
	stackql_cmd_bin := exec.Command("curl", "-L", stackql_bin_url, "-o", "binary.bin")
	err = stackql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stackql_src_url := "https://github.com/stackql/stackql/archive/refs/tags/v0.5.734.src.tar.gz"
	stackql_cmd_src := exec.Command("curl", "-L", stackql_src_url, "-o", "source.tar.gz")
	err = stackql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stackql_cmd_direct := exec.Command("./binary")
	err = stackql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
