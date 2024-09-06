package main

// Jmeter - Load testing and performance measurement application
// Homepage: https://jmeter.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installJmeter() {
	// Método 1: Descargar y extraer .tar.gz
	jmeter_tar_url := "https://www.apache.org/dyn/closer.lua?path=jmeter/binaries/apache-jmeter-5.6.3.tgz"
	jmeter_cmd_tar := exec.Command("curl", "-L", jmeter_tar_url, "-o", "package.tar.gz")
	err := jmeter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jmeter_zip_url := "https://www.apache.org/dyn/closer.lua?path=jmeter/binaries/apache-jmeter-5.6.3.tgz"
	jmeter_cmd_zip := exec.Command("curl", "-L", jmeter_zip_url, "-o", "package.zip")
	err = jmeter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jmeter_bin_url := "https://www.apache.org/dyn/closer.lua?path=jmeter/binaries/apache-jmeter-5.6.3.tgz"
	jmeter_cmd_bin := exec.Command("curl", "-L", jmeter_bin_url, "-o", "binary.bin")
	err = jmeter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jmeter_src_url := "https://www.apache.org/dyn/closer.lua?path=jmeter/binaries/apache-jmeter-5.6.3.tgz"
	jmeter_cmd_src := exec.Command("curl", "-L", jmeter_src_url, "-o", "source.tar.gz")
	err = jmeter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jmeter_cmd_direct := exec.Command("./binary")
	err = jmeter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@21")
	exec.Command("latte", "install", "openjdk@21").Run()
}
