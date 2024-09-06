package main

// Log4cxx - Library of C++ classes for flexible logging
// Homepage: https://logging.apache.org/log4cxx/index.html

import (
	"fmt"
	
	"os/exec"
)

func installLog4cxx() {
	// Método 1: Descargar y extraer .tar.gz
	log4cxx_tar_url := "https://www.apache.org/dyn/closer.lua?path=logging/log4cxx/1.2.0/apache-log4cxx-1.2.0.tar.gz"
	log4cxx_cmd_tar := exec.Command("curl", "-L", log4cxx_tar_url, "-o", "package.tar.gz")
	err := log4cxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	log4cxx_zip_url := "https://www.apache.org/dyn/closer.lua?path=logging/log4cxx/1.2.0/apache-log4cxx-1.2.0.zip"
	log4cxx_cmd_zip := exec.Command("curl", "-L", log4cxx_zip_url, "-o", "package.zip")
	err = log4cxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	log4cxx_bin_url := "https://www.apache.org/dyn/closer.lua?path=logging/log4cxx/1.2.0/apache-log4cxx-1.2.0.bin"
	log4cxx_cmd_bin := exec.Command("curl", "-L", log4cxx_bin_url, "-o", "binary.bin")
	err = log4cxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	log4cxx_src_url := "https://www.apache.org/dyn/closer.lua?path=logging/log4cxx/1.2.0/apache-log4cxx-1.2.0.src.tar.gz"
	log4cxx_cmd_src := exec.Command("curl", "-L", log4cxx_src_url, "-o", "source.tar.gz")
	err = log4cxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	log4cxx_cmd_direct := exec.Command("./binary")
	err = log4cxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: apr-util")
exec.Command("latte", "install", "apr-util")
}
