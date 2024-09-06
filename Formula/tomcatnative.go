package main

// TomcatNative - Lets Tomcat use some native resources for performance
// Homepage: https://tomcat.apache.org/native-doc/

import (
	"fmt"
	
	"os/exec"
)

func installTomcatNative() {
	// Método 1: Descargar y extraer .tar.gz
	tomcatnative_tar_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-connectors/native/2.0.8/source/tomcat-native-2.0.8-src.tar.gz"
	tomcatnative_cmd_tar := exec.Command("curl", "-L", tomcatnative_tar_url, "-o", "package.tar.gz")
	err := tomcatnative_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tomcatnative_zip_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-connectors/native/2.0.8/source/tomcat-native-2.0.8-src.zip"
	tomcatnative_cmd_zip := exec.Command("curl", "-L", tomcatnative_zip_url, "-o", "package.zip")
	err = tomcatnative_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tomcatnative_bin_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-connectors/native/2.0.8/source/tomcat-native-2.0.8-src.bin"
	tomcatnative_cmd_bin := exec.Command("curl", "-L", tomcatnative_bin_url, "-o", "binary.bin")
	err = tomcatnative_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tomcatnative_src_url := "https://www.apache.org/dyn/closer.lua?path=tomcat/tomcat-connectors/native/2.0.8/source/tomcat-native-2.0.8-src.src.tar.gz"
	tomcatnative_cmd_src := exec.Command("curl", "-L", tomcatnative_src_url, "-o", "source.tar.gz")
	err = tomcatnative_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tomcatnative_cmd_direct := exec.Command("./binary")
	err = tomcatnative_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tomcat")
	exec.Command("latte", "install", "tomcat").Run()
	fmt.Println("Instalando dependencia: apr")
	exec.Command("latte", "install", "apr").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
