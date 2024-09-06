package main

// Couchdb - Apache CouchDB database server
// Homepage: https://couchdb.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installCouchdb() {
	// Método 1: Descargar y extraer .tar.gz
	couchdb_tar_url := "https://www.apache.org/dyn/closer.lua?path=couchdb/source/3.3.3/apache-couchdb-3.3.3.tar.gz"
	couchdb_cmd_tar := exec.Command("curl", "-L", couchdb_tar_url, "-o", "package.tar.gz")
	err := couchdb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	couchdb_zip_url := "https://www.apache.org/dyn/closer.lua?path=couchdb/source/3.3.3/apache-couchdb-3.3.3.zip"
	couchdb_cmd_zip := exec.Command("curl", "-L", couchdb_zip_url, "-o", "package.zip")
	err = couchdb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	couchdb_bin_url := "https://www.apache.org/dyn/closer.lua?path=couchdb/source/3.3.3/apache-couchdb-3.3.3.bin"
	couchdb_cmd_bin := exec.Command("curl", "-L", couchdb_bin_url, "-o", "binary.bin")
	err = couchdb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	couchdb_src_url := "https://www.apache.org/dyn/closer.lua?path=couchdb/source/3.3.3/apache-couchdb-3.3.3.src.tar.gz"
	couchdb_cmd_src := exec.Command("curl", "-L", couchdb_src_url, "-o", "source.tar.gz")
	err = couchdb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	couchdb_cmd_direct := exec.Command("./binary")
	err = couchdb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: autoconf-archive")
exec.Command("latte", "install", "autoconf-archive")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: erlang@25")
exec.Command("latte", "install", "erlang@25")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: spidermonkey@91")
exec.Command("latte", "install", "spidermonkey@91")
}
