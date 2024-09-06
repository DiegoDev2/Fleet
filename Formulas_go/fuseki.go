package main

// Fuseki - SPARQL server
// Homepage: https://jena.apache.org/documentation/fuseki2/

import (
	"fmt"
	
	"os/exec"
)

func installFuseki() {
	// Método 1: Descargar y extraer .tar.gz
	fuseki_tar_url := "https://www.apache.org/dyn/closer.lua?path=jena/binaries/apache-jena-fuseki-5.1.0.tar.gz"
	fuseki_cmd_tar := exec.Command("curl", "-L", fuseki_tar_url, "-o", "package.tar.gz")
	err := fuseki_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fuseki_zip_url := "https://www.apache.org/dyn/closer.lua?path=jena/binaries/apache-jena-fuseki-5.1.0.zip"
	fuseki_cmd_zip := exec.Command("curl", "-L", fuseki_zip_url, "-o", "package.zip")
	err = fuseki_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fuseki_bin_url := "https://www.apache.org/dyn/closer.lua?path=jena/binaries/apache-jena-fuseki-5.1.0.bin"
	fuseki_cmd_bin := exec.Command("curl", "-L", fuseki_bin_url, "-o", "binary.bin")
	err = fuseki_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fuseki_src_url := "https://www.apache.org/dyn/closer.lua?path=jena/binaries/apache-jena-fuseki-5.1.0.src.tar.gz"
	fuseki_cmd_src := exec.Command("curl", "-L", fuseki_src_url, "-o", "source.tar.gz")
	err = fuseki_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fuseki_cmd_direct := exec.Command("./binary")
	err = fuseki_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
