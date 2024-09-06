package main

// Tomcat - Implementation of Java Servlet and JavaServer Pages
// Homepage: https://tomcat.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installTomcat() {
	// Método 1: Descargar y extraer .tar.gz
	tomcat_tar_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-10/v10.1.28/bin/apache-tomcat-10.1.28.tar.gz"
	tomcat_cmd_tar := exec.Command("curl", "-L", tomcat_tar_url, "-o", "package.tar.gz")
	err := tomcat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tomcat_zip_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-10/v10.1.28/bin/apache-tomcat-10.1.28.zip"
	tomcat_cmd_zip := exec.Command("curl", "-L", tomcat_zip_url, "-o", "package.zip")
	err = tomcat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tomcat_bin_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-10/v10.1.28/bin/apache-tomcat-10.1.28.bin"
	tomcat_cmd_bin := exec.Command("curl", "-L", tomcat_bin_url, "-o", "binary.bin")
	err = tomcat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tomcat_src_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-10/v10.1.28/bin/apache-tomcat-10.1.28.src.tar.gz"
	tomcat_cmd_src := exec.Command("curl", "-L", tomcat_src_url, "-o", "source.tar.gz")
	err = tomcat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tomcat_cmd_direct := exec.Command("./binary")
	err = tomcat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
