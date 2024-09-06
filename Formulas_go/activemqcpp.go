package main

// ActivemqCpp - C++ API for message brokers such as Apache ActiveMQ
// Homepage: https://activemq.apache.org/components/cms/

import (
	"fmt"
	
	"os/exec"
)

func installActivemqCpp() {
	// Método 1: Descargar y extraer .tar.gz
	activemqcpp_tar_url := "https://www.apache.org/dyn/closer.lua?path=activemq/activemq-cpp/3.9.5/activemq-cpp-library-3.9.5-src.tar.bz2"
	activemqcpp_cmd_tar := exec.Command("curl", "-L", activemqcpp_tar_url, "-o", "package.tar.gz")
	err := activemqcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	activemqcpp_zip_url := "https://www.apache.org/dyn/closer.lua?path=activemq/activemq-cpp/3.9.5/activemq-cpp-library-3.9.5-src.tar.bz2"
	activemqcpp_cmd_zip := exec.Command("curl", "-L", activemqcpp_zip_url, "-o", "package.zip")
	err = activemqcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	activemqcpp_bin_url := "https://www.apache.org/dyn/closer.lua?path=activemq/activemq-cpp/3.9.5/activemq-cpp-library-3.9.5-src.tar.bz2"
	activemqcpp_cmd_bin := exec.Command("curl", "-L", activemqcpp_bin_url, "-o", "binary.bin")
	err = activemqcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	activemqcpp_src_url := "https://www.apache.org/dyn/closer.lua?path=activemq/activemq-cpp/3.9.5/activemq-cpp-library-3.9.5-src.tar.bz2"
	activemqcpp_cmd_src := exec.Command("curl", "-L", activemqcpp_src_url, "-o", "source.tar.gz")
	err = activemqcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	activemqcpp_cmd_direct := exec.Command("./binary")
	err = activemqcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: apr")
exec.Command("latte", "install", "apr")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
