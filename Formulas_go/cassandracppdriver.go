package main

// CassandraCppDriver - DataStax C/C++ Driver for Apache Cassandra
// Homepage: https://docs.datastax.com/en/developer/cpp-driver/latest

import (
	"fmt"
	
	"os/exec"
)

func installCassandraCppDriver() {
	// Método 1: Descargar y extraer .tar.gz
	cassandracppdriver_tar_url := "https://github.com/datastax/cpp-driver/archive/refs/tags/2.17.1.tar.gz"
	cassandracppdriver_cmd_tar := exec.Command("curl", "-L", cassandracppdriver_tar_url, "-o", "package.tar.gz")
	err := cassandracppdriver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cassandracppdriver_zip_url := "https://github.com/datastax/cpp-driver/archive/refs/tags/2.17.1.zip"
	cassandracppdriver_cmd_zip := exec.Command("curl", "-L", cassandracppdriver_zip_url, "-o", "package.zip")
	err = cassandracppdriver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cassandracppdriver_bin_url := "https://github.com/datastax/cpp-driver/archive/refs/tags/2.17.1.bin"
	cassandracppdriver_cmd_bin := exec.Command("curl", "-L", cassandracppdriver_bin_url, "-o", "binary.bin")
	err = cassandracppdriver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cassandracppdriver_src_url := "https://github.com/datastax/cpp-driver/archive/refs/tags/2.17.1.src.tar.gz"
	cassandracppdriver_cmd_src := exec.Command("curl", "-L", cassandracppdriver_src_url, "-o", "source.tar.gz")
	err = cassandracppdriver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cassandracppdriver_cmd_direct := exec.Command("./binary")
	err = cassandracppdriver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
