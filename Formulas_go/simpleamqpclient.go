package main

// SimpleAmqpClient - C++ interface to rabbitmq-c
// Homepage: https://github.com/alanxz/SimpleAmqpClient

import (
	"fmt"
	
	"os/exec"
)

func installSimpleAmqpClient() {
	// Método 1: Descargar y extraer .tar.gz
	simpleamqpclient_tar_url := "https://github.com/alanxz/SimpleAmqpClient/archive/refs/tags/v2.5.1.tar.gz"
	simpleamqpclient_cmd_tar := exec.Command("curl", "-L", simpleamqpclient_tar_url, "-o", "package.tar.gz")
	err := simpleamqpclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	simpleamqpclient_zip_url := "https://github.com/alanxz/SimpleAmqpClient/archive/refs/tags/v2.5.1.zip"
	simpleamqpclient_cmd_zip := exec.Command("curl", "-L", simpleamqpclient_zip_url, "-o", "package.zip")
	err = simpleamqpclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	simpleamqpclient_bin_url := "https://github.com/alanxz/SimpleAmqpClient/archive/refs/tags/v2.5.1.bin"
	simpleamqpclient_cmd_bin := exec.Command("curl", "-L", simpleamqpclient_bin_url, "-o", "binary.bin")
	err = simpleamqpclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	simpleamqpclient_src_url := "https://github.com/alanxz/SimpleAmqpClient/archive/refs/tags/v2.5.1.src.tar.gz"
	simpleamqpclient_cmd_src := exec.Command("curl", "-L", simpleamqpclient_src_url, "-o", "source.tar.gz")
	err = simpleamqpclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	simpleamqpclient_cmd_direct := exec.Command("./binary")
	err = simpleamqpclient_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: rabbitmq-c")
exec.Command("latte", "install", "rabbitmq-c")
}
