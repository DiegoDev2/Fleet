package main

// Logstash - Tool for managing events and logs
// Homepage: https://www.elastic.co/products/logstash

import (
	"fmt"
	
	"os/exec"
)

func installLogstash() {
	// Método 1: Descargar y extraer .tar.gz
	logstash_tar_url := "https://github.com/elastic/logstash/archive/refs/tags/v8.15.1.tar.gz"
	logstash_cmd_tar := exec.Command("curl", "-L", logstash_tar_url, "-o", "package.tar.gz")
	err := logstash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	logstash_zip_url := "https://github.com/elastic/logstash/archive/refs/tags/v8.15.1.zip"
	logstash_cmd_zip := exec.Command("curl", "-L", logstash_zip_url, "-o", "package.zip")
	err = logstash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	logstash_bin_url := "https://github.com/elastic/logstash/archive/refs/tags/v8.15.1.bin"
	logstash_cmd_bin := exec.Command("curl", "-L", logstash_bin_url, "-o", "binary.bin")
	err = logstash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	logstash_src_url := "https://github.com/elastic/logstash/archive/refs/tags/v8.15.1.src.tar.gz"
	logstash_cmd_src := exec.Command("curl", "-L", logstash_src_url, "-o", "source.tar.gz")
	err = logstash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	logstash_cmd_direct := exec.Command("./binary")
	err = logstash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@17")
	exec.Command("latte", "install", "openjdk@17").Run()
}
