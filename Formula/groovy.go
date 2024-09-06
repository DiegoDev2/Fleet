package main

// Groovy - Java-based scripting language
// Homepage: https://www.groovy-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installGroovy() {
	// Método 1: Descargar y extraer .tar.gz
	groovy_tar_url := "https://groovy.jfrog.io/artifactory/dist-release-local/groovy-zips/apache-groovy-binary-4.0.22.zip"
	groovy_cmd_tar := exec.Command("curl", "-L", groovy_tar_url, "-o", "package.tar.gz")
	err := groovy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	groovy_zip_url := "https://groovy.jfrog.io/artifactory/dist-release-local/groovy-zips/apache-groovy-binary-4.0.22.zip"
	groovy_cmd_zip := exec.Command("curl", "-L", groovy_zip_url, "-o", "package.zip")
	err = groovy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	groovy_bin_url := "https://groovy.jfrog.io/artifactory/dist-release-local/groovy-zips/apache-groovy-binary-4.0.22.zip"
	groovy_cmd_bin := exec.Command("curl", "-L", groovy_bin_url, "-o", "binary.bin")
	err = groovy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	groovy_src_url := "https://groovy.jfrog.io/artifactory/dist-release-local/groovy-zips/apache-groovy-binary-4.0.22.zip"
	groovy_cmd_src := exec.Command("curl", "-L", groovy_src_url, "-o", "source.tar.gz")
	err = groovy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	groovy_cmd_direct := exec.Command("./binary")
	err = groovy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: maven")
	exec.Command("latte", "install", "maven").Run()
}
