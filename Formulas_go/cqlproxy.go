package main

// CqlProxy - DataStax cql-proxy enables Cassandra apps to use Astra DB without code changes
// Homepage: https://github.com/datastax/cql-proxy

import (
	"fmt"
	
	"os/exec"
)

func installCqlProxy() {
	// Método 1: Descargar y extraer .tar.gz
	cqlproxy_tar_url := "https://github.com/datastax/cql-proxy/archive/refs/tags/v0.1.6.tar.gz"
	cqlproxy_cmd_tar := exec.Command("curl", "-L", cqlproxy_tar_url, "-o", "package.tar.gz")
	err := cqlproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cqlproxy_zip_url := "https://github.com/datastax/cql-proxy/archive/refs/tags/v0.1.6.zip"
	cqlproxy_cmd_zip := exec.Command("curl", "-L", cqlproxy_zip_url, "-o", "package.zip")
	err = cqlproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cqlproxy_bin_url := "https://github.com/datastax/cql-proxy/archive/refs/tags/v0.1.6.bin"
	cqlproxy_cmd_bin := exec.Command("curl", "-L", cqlproxy_bin_url, "-o", "binary.bin")
	err = cqlproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cqlproxy_src_url := "https://github.com/datastax/cql-proxy/archive/refs/tags/v0.1.6.src.tar.gz"
	cqlproxy_cmd_src := exec.Command("curl", "-L", cqlproxy_src_url, "-o", "source.tar.gz")
	err = cqlproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cqlproxy_cmd_direct := exec.Command("./binary")
	err = cqlproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
