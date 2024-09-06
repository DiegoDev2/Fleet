package main

// Postgraphile - GraphQL schema created by reflection over a PostgreSQL schema
// Homepage: https://www.graphile.org/postgraphile/

import (
	"fmt"
	
	"os/exec"
)

func installPostgraphile() {
	// Método 1: Descargar y extraer .tar.gz
	postgraphile_tar_url := "https://registry.npmjs.org/postgraphile/-/postgraphile-4.14.0.tgz"
	postgraphile_cmd_tar := exec.Command("curl", "-L", postgraphile_tar_url, "-o", "package.tar.gz")
	err := postgraphile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	postgraphile_zip_url := "https://registry.npmjs.org/postgraphile/-/postgraphile-4.14.0.tgz"
	postgraphile_cmd_zip := exec.Command("curl", "-L", postgraphile_zip_url, "-o", "package.zip")
	err = postgraphile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	postgraphile_bin_url := "https://registry.npmjs.org/postgraphile/-/postgraphile-4.14.0.tgz"
	postgraphile_cmd_bin := exec.Command("curl", "-L", postgraphile_bin_url, "-o", "binary.bin")
	err = postgraphile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	postgraphile_src_url := "https://registry.npmjs.org/postgraphile/-/postgraphile-4.14.0.tgz"
	postgraphile_cmd_src := exec.Command("curl", "-L", postgraphile_src_url, "-o", "source.tar.gz")
	err = postgraphile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	postgraphile_cmd_direct := exec.Command("./binary")
	err = postgraphile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: postgresql@16")
exec.Command("latte", "install", "postgresql@16")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
