package main

// SolrAT811 - Enterprise search platform from the Apache Lucene project
// Homepage: https://solr.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installSolrAT811() {
	// Método 1: Descargar y extraer .tar.gz
	solrat811_tar_url := "https://dlcdn.apache.org/lucene/solr/8.11.3/solr-8.11.3.tgz"
	solrat811_cmd_tar := exec.Command("curl", "-L", solrat811_tar_url, "-o", "package.tar.gz")
	err := solrat811_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	solrat811_zip_url := "https://dlcdn.apache.org/lucene/solr/8.11.3/solr-8.11.3.tgz"
	solrat811_cmd_zip := exec.Command("curl", "-L", solrat811_zip_url, "-o", "package.zip")
	err = solrat811_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	solrat811_bin_url := "https://dlcdn.apache.org/lucene/solr/8.11.3/solr-8.11.3.tgz"
	solrat811_cmd_bin := exec.Command("curl", "-L", solrat811_bin_url, "-o", "binary.bin")
	err = solrat811_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	solrat811_src_url := "https://dlcdn.apache.org/lucene/solr/8.11.3/solr-8.11.3.tgz"
	solrat811_cmd_src := exec.Command("curl", "-L", solrat811_src_url, "-o", "source.tar.gz")
	err = solrat811_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	solrat811_cmd_direct := exec.Command("./binary")
	err = solrat811_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
