package main

// Zmqpp - High-level C++ binding for zeromq
// Homepage: https://zeromq.github.io/zmqpp/

import (
	"fmt"
	
	"os/exec"
)

func installZmqpp() {
	// Método 1: Descargar y extraer .tar.gz
	zmqpp_tar_url := "https://github.com/zeromq/zmqpp/archive/refs/tags/4.2.0.tar.gz"
	zmqpp_cmd_tar := exec.Command("curl", "-L", zmqpp_tar_url, "-o", "package.tar.gz")
	err := zmqpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zmqpp_zip_url := "https://github.com/zeromq/zmqpp/archive/refs/tags/4.2.0.zip"
	zmqpp_cmd_zip := exec.Command("curl", "-L", zmqpp_zip_url, "-o", "package.zip")
	err = zmqpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zmqpp_bin_url := "https://github.com/zeromq/zmqpp/archive/refs/tags/4.2.0.bin"
	zmqpp_cmd_bin := exec.Command("curl", "-L", zmqpp_bin_url, "-o", "binary.bin")
	err = zmqpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zmqpp_src_url := "https://github.com/zeromq/zmqpp/archive/refs/tags/4.2.0.src.tar.gz"
	zmqpp_cmd_src := exec.Command("curl", "-L", zmqpp_src_url, "-o", "source.tar.gz")
	err = zmqpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zmqpp_cmd_direct := exec.Command("./binary")
	err = zmqpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: zeromq")
exec.Command("latte", "install", "zeromq")
}
