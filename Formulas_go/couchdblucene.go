package main

// CouchdbLucene - Full-text search of CouchDB documents using Lucene
// Homepage: https://github.com/rnewson/couchdb-lucene

import (
	"fmt"
	
	"os/exec"
)

func installCouchdbLucene() {
	// Método 1: Descargar y extraer .tar.gz
	couchdblucene_tar_url := "https://github.com/rnewson/couchdb-lucene/archive/refs/tags/v2.1.0.tar.gz"
	couchdblucene_cmd_tar := exec.Command("curl", "-L", couchdblucene_tar_url, "-o", "package.tar.gz")
	err := couchdblucene_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	couchdblucene_zip_url := "https://github.com/rnewson/couchdb-lucene/archive/refs/tags/v2.1.0.zip"
	couchdblucene_cmd_zip := exec.Command("curl", "-L", couchdblucene_zip_url, "-o", "package.zip")
	err = couchdblucene_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	couchdblucene_bin_url := "https://github.com/rnewson/couchdb-lucene/archive/refs/tags/v2.1.0.bin"
	couchdblucene_cmd_bin := exec.Command("curl", "-L", couchdblucene_bin_url, "-o", "binary.bin")
	err = couchdblucene_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	couchdblucene_src_url := "https://github.com/rnewson/couchdb-lucene/archive/refs/tags/v2.1.0.src.tar.gz"
	couchdblucene_cmd_src := exec.Command("curl", "-L", couchdblucene_src_url, "-o", "source.tar.gz")
	err = couchdblucene_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	couchdblucene_cmd_direct := exec.Command("./binary")
	err = couchdblucene_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
	fmt.Println("Instalando dependencia: couchdb")
exec.Command("latte", "install", "couchdb")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
