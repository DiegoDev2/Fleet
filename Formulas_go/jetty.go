package main

// Jetty - Java servlet engine and webserver
// Homepage: https://jetty.org/

import (
	"fmt"
	
	"os/exec"
)

func installJetty() {
	// Método 1: Descargar y extraer .tar.gz
	jetty_tar_url := "https://search.maven.org/remotecontent?filepath=org/eclipse/jetty/jetty-distribution/9.4.56.v20240826/jetty-distribution-9.4.56.v20240826.tar.gz"
	jetty_cmd_tar := exec.Command("curl", "-L", jetty_tar_url, "-o", "package.tar.gz")
	err := jetty_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jetty_zip_url := "https://search.maven.org/remotecontent?filepath=org/eclipse/jetty/jetty-distribution/9.4.56.v20240826/jetty-distribution-9.4.56.v20240826.zip"
	jetty_cmd_zip := exec.Command("curl", "-L", jetty_zip_url, "-o", "package.zip")
	err = jetty_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jetty_bin_url := "https://search.maven.org/remotecontent?filepath=org/eclipse/jetty/jetty-distribution/9.4.56.v20240826/jetty-distribution-9.4.56.v20240826.bin"
	jetty_cmd_bin := exec.Command("curl", "-L", jetty_bin_url, "-o", "binary.bin")
	err = jetty_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jetty_src_url := "https://search.maven.org/remotecontent?filepath=org/eclipse/jetty/jetty-distribution/9.4.56.v20240826/jetty-distribution-9.4.56.v20240826.src.tar.gz"
	jetty_cmd_src := exec.Command("curl", "-L", jetty_src_url, "-o", "source.tar.gz")
	err = jetty_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jetty_cmd_direct := exec.Command("./binary")
	err = jetty_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
}
