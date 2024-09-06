package main

// ThriftAT09 - Framework for scalable cross-language services development
// Homepage: https://thrift.apache.org

import (
	"fmt"
	
	"os/exec"
)

func installThriftAT09() {
	// Método 1: Descargar y extraer .tar.gz
	thriftat09_tar_url := "https://github.com/apache/thrift/archive/refs/tags/0.9.3.1.tar.gz"
	thriftat09_cmd_tar := exec.Command("curl", "-L", thriftat09_tar_url, "-o", "package.tar.gz")
	err := thriftat09_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	thriftat09_zip_url := "https://github.com/apache/thrift/archive/refs/tags/0.9.3.1.zip"
	thriftat09_cmd_zip := exec.Command("curl", "-L", thriftat09_zip_url, "-o", "package.zip")
	err = thriftat09_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	thriftat09_bin_url := "https://github.com/apache/thrift/archive/refs/tags/0.9.3.1.bin"
	thriftat09_cmd_bin := exec.Command("curl", "-L", thriftat09_bin_url, "-o", "binary.bin")
	err = thriftat09_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	thriftat09_src_url := "https://github.com/apache/thrift/archive/refs/tags/0.9.3.1.src.tar.gz"
	thriftat09_cmd_src := exec.Command("curl", "-L", thriftat09_src_url, "-o", "source.tar.gz")
	err = thriftat09_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	thriftat09_cmd_direct := exec.Command("./binary")
	err = thriftat09_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
}
