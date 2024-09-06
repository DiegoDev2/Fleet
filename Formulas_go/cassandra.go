package main

// Cassandra - Eventually consistent, distributed key-value store
// Homepage: https://cassandra.apache.org

import (
	"fmt"
	
	"os/exec"
)

func installCassandra() {
	// Método 1: Descargar y extraer .tar.gz
	cassandra_tar_url := "https://www.apache.org/dyn/closer.lua?path=cassandra/4.1.6/apache-cassandra-4.1.6-bin.tar.gz"
	cassandra_cmd_tar := exec.Command("curl", "-L", cassandra_tar_url, "-o", "package.tar.gz")
	err := cassandra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cassandra_zip_url := "https://www.apache.org/dyn/closer.lua?path=cassandra/4.1.6/apache-cassandra-4.1.6-bin.zip"
	cassandra_cmd_zip := exec.Command("curl", "-L", cassandra_zip_url, "-o", "package.zip")
	err = cassandra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cassandra_bin_url := "https://www.apache.org/dyn/closer.lua?path=cassandra/4.1.6/apache-cassandra-4.1.6-bin.bin"
	cassandra_cmd_bin := exec.Command("curl", "-L", cassandra_bin_url, "-o", "binary.bin")
	err = cassandra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cassandra_src_url := "https://www.apache.org/dyn/closer.lua?path=cassandra/4.1.6/apache-cassandra-4.1.6-bin.src.tar.gz"
	cassandra_cmd_src := exec.Command("curl", "-L", cassandra_src_url, "-o", "source.tar.gz")
	err = cassandra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cassandra_cmd_direct := exec.Command("./binary")
	err = cassandra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libev")
exec.Command("latte", "install", "libev")
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
