package main

// Solr - Enterprise search platform from the Apache Lucene project
// Homepage: https://solr.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installSolr() {
	// Método 1: Descargar y extraer .tar.gz
	solr_tar_url := "https://dlcdn.apache.org/solr/solr/9.6.1/solr-9.6.1.tgz"
	solr_cmd_tar := exec.Command("curl", "-L", solr_tar_url, "-o", "package.tar.gz")
	err := solr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	solr_zip_url := "https://dlcdn.apache.org/solr/solr/9.6.1/solr-9.6.1.tgz"
	solr_cmd_zip := exec.Command("curl", "-L", solr_zip_url, "-o", "package.zip")
	err = solr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	solr_bin_url := "https://dlcdn.apache.org/solr/solr/9.6.1/solr-9.6.1.tgz"
	solr_cmd_bin := exec.Command("curl", "-L", solr_bin_url, "-o", "binary.bin")
	err = solr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	solr_src_url := "https://dlcdn.apache.org/solr/solr/9.6.1/solr-9.6.1.tgz"
	solr_cmd_src := exec.Command("curl", "-L", solr_src_url, "-o", "source.tar.gz")
	err = solr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	solr_cmd_direct := exec.Command("./binary")
	err = solr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
