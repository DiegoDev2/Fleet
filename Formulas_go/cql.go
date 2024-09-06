package main

// Cql - Decentralized SQL database with blockchain features
// Homepage: https://covenantsql.io

import (
	"fmt"
	
	"os/exec"
)

func installCql() {
	// Método 1: Descargar y extraer .tar.gz
	cql_tar_url := "https://github.com/CovenantSQL/CovenantSQL/archive/refs/tags/v0.8.1.tar.gz"
	cql_cmd_tar := exec.Command("curl", "-L", cql_tar_url, "-o", "package.tar.gz")
	err := cql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cql_zip_url := "https://github.com/CovenantSQL/CovenantSQL/archive/refs/tags/v0.8.1.zip"
	cql_cmd_zip := exec.Command("curl", "-L", cql_zip_url, "-o", "package.zip")
	err = cql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cql_bin_url := "https://github.com/CovenantSQL/CovenantSQL/archive/refs/tags/v0.8.1.bin"
	cql_cmd_bin := exec.Command("curl", "-L", cql_bin_url, "-o", "binary.bin")
	err = cql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cql_src_url := "https://github.com/CovenantSQL/CovenantSQL/archive/refs/tags/v0.8.1.src.tar.gz"
	cql_cmd_src := exec.Command("curl", "-L", cql_src_url, "-o", "source.tar.gz")
	err = cql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cql_cmd_direct := exec.Command("./binary")
	err = cql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
