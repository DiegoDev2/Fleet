package main

// Activemq - Apache ActiveMQ: powerful open source messaging server
// Homepage: https://activemq.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installActivemq() {
	// Método 1: Descargar y extraer .tar.gz
	activemq_tar_url := "https://www.apache.org/dyn/closer.lua?path=activemq/6.1.3/apache-activemq-6.1.3-bin.tar.gz"
	activemq_cmd_tar := exec.Command("curl", "-L", activemq_tar_url, "-o", "package.tar.gz")
	err := activemq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	activemq_zip_url := "https://www.apache.org/dyn/closer.lua?path=activemq/6.1.3/apache-activemq-6.1.3-bin.zip"
	activemq_cmd_zip := exec.Command("curl", "-L", activemq_zip_url, "-o", "package.zip")
	err = activemq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	activemq_bin_url := "https://www.apache.org/dyn/closer.lua?path=activemq/6.1.3/apache-activemq-6.1.3-bin.bin"
	activemq_cmd_bin := exec.Command("curl", "-L", activemq_bin_url, "-o", "binary.bin")
	err = activemq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	activemq_src_url := "https://www.apache.org/dyn/closer.lua?path=activemq/6.1.3/apache-activemq-6.1.3-bin.src.tar.gz"
	activemq_cmd_src := exec.Command("curl", "-L", activemq_src_url, "-o", "source.tar.gz")
	err = activemq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	activemq_cmd_direct := exec.Command("./binary")
	err = activemq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: java-service-wrapper")
	exec.Command("latte", "install", "java-service-wrapper").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
