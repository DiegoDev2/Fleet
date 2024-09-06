package main

// Orientdb - Graph database
// Homepage: https://orientdb.org/

import (
	"fmt"
	
	"os/exec"
)

func installOrientdb() {
	// Método 1: Descargar y extraer .tar.gz
	orientdb_tar_url := "https://search.maven.org/remotecontent?filepath=com/orientechnologies/orientdb-community/3.2.33/orientdb-community-3.2.33.zip"
	orientdb_cmd_tar := exec.Command("curl", "-L", orientdb_tar_url, "-o", "package.tar.gz")
	err := orientdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	orientdb_zip_url := "https://search.maven.org/remotecontent?filepath=com/orientechnologies/orientdb-community/3.2.33/orientdb-community-3.2.33.zip"
	orientdb_cmd_zip := exec.Command("curl", "-L", orientdb_zip_url, "-o", "package.zip")
	err = orientdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	orientdb_bin_url := "https://search.maven.org/remotecontent?filepath=com/orientechnologies/orientdb-community/3.2.33/orientdb-community-3.2.33.zip"
	orientdb_cmd_bin := exec.Command("curl", "-L", orientdb_bin_url, "-o", "binary.bin")
	err = orientdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	orientdb_src_url := "https://search.maven.org/remotecontent?filepath=com/orientechnologies/orientdb-community/3.2.33/orientdb-community-3.2.33.zip"
	orientdb_cmd_src := exec.Command("curl", "-L", orientdb_src_url, "-o", "source.tar.gz")
	err = orientdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	orientdb_cmd_direct := exec.Command("./binary")
	err = orientdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
