package main

// Neo4j - Robust (fully ACID) transactional property graph database
// Homepage: https://neo4j.com/

import (
	"fmt"
	
	"os/exec"
)

func installNeo4j() {
	// Método 1: Descargar y extraer .tar.gz
	neo4j_tar_url := "https://neo4j.com/artifact.php?name=neo4j-community-5.23.0-unix.tar.gz"
	neo4j_cmd_tar := exec.Command("curl", "-L", neo4j_tar_url, "-o", "package.tar.gz")
	err := neo4j_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neo4j_zip_url := "https://neo4j.com/artifact.php?name=neo4j-community-5.23.0-unix.zip"
	neo4j_cmd_zip := exec.Command("curl", "-L", neo4j_zip_url, "-o", "package.zip")
	err = neo4j_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neo4j_bin_url := "https://neo4j.com/artifact.php?name=neo4j-community-5.23.0-unix.bin"
	neo4j_cmd_bin := exec.Command("curl", "-L", neo4j_bin_url, "-o", "binary.bin")
	err = neo4j_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neo4j_src_url := "https://neo4j.com/artifact.php?name=neo4j-community-5.23.0-unix.src.tar.gz"
	neo4j_cmd_src := exec.Command("curl", "-L", neo4j_src_url, "-o", "source.tar.gz")
	err = neo4j_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neo4j_cmd_direct := exec.Command("./binary")
	err = neo4j_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cypher-shell")
exec.Command("latte", "install", "cypher-shell")
	fmt.Println("Instalando dependencia: openjdk@21")
exec.Command("latte", "install", "openjdk@21")
}
