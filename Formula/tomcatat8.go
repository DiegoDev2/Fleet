package main

// TomcatAT8 - Implementation of Java Servlet and JavaServer Pages
// Homepage: https://tomcat.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installTomcatAT8() {
	// Método 1: Descargar y extraer .tar.gz
	tomcatat8_tar_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-8/v8.5.100/bin/apache-tomcat-8.5.100.tar.gz"
	tomcatat8_cmd_tar := exec.Command("curl", "-L", tomcatat8_tar_url, "-o", "package.tar.gz")
	err := tomcatat8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tomcatat8_zip_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-8/v8.5.100/bin/apache-tomcat-8.5.100.zip"
	tomcatat8_cmd_zip := exec.Command("curl", "-L", tomcatat8_zip_url, "-o", "package.zip")
	err = tomcatat8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tomcatat8_bin_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-8/v8.5.100/bin/apache-tomcat-8.5.100.bin"
	tomcatat8_cmd_bin := exec.Command("curl", "-L", tomcatat8_bin_url, "-o", "binary.bin")
	err = tomcatat8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tomcatat8_src_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-8/v8.5.100/bin/apache-tomcat-8.5.100.src.tar.gz"
	tomcatat8_cmd_src := exec.Command("curl", "-L", tomcatat8_src_url, "-o", "source.tar.gz")
	err = tomcatat8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tomcatat8_cmd_direct := exec.Command("./binary")
	err = tomcatat8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
