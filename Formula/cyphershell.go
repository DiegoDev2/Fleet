package main

// CypherShell - Command-line shell where you can execute Cypher against Neo4j
// Homepage: https://neo4j.com

import (
	"fmt"
	
	"os/exec"
)

func installCypherShell() {
	// Método 1: Descargar y extraer .tar.gz
	cyphershell_tar_url := "https://dist.neo4j.org/cypher-shell/cypher-shell-5.23.0.zip"
	cyphershell_cmd_tar := exec.Command("curl", "-L", cyphershell_tar_url, "-o", "package.tar.gz")
	err := cyphershell_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cyphershell_zip_url := "https://dist.neo4j.org/cypher-shell/cypher-shell-5.23.0.zip"
	cyphershell_cmd_zip := exec.Command("curl", "-L", cyphershell_zip_url, "-o", "package.zip")
	err = cyphershell_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cyphershell_bin_url := "https://dist.neo4j.org/cypher-shell/cypher-shell-5.23.0.zip"
	cyphershell_cmd_bin := exec.Command("curl", "-L", cyphershell_bin_url, "-o", "binary.bin")
	err = cyphershell_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cyphershell_src_url := "https://dist.neo4j.org/cypher-shell/cypher-shell-5.23.0.zip"
	cyphershell_cmd_src := exec.Command("curl", "-L", cyphershell_src_url, "-o", "source.tar.gz")
	err = cyphershell_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cyphershell_cmd_direct := exec.Command("./binary")
	err = cyphershell_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@21")
	exec.Command("latte", "install", "openjdk@21").Run()
}
