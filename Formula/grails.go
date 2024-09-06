package main

// Grails - Web application framework for the Groovy language
// Homepage: https://grails.org

import (
	"fmt"
	
	"os/exec"
)

func installGrails() {
	// Método 1: Descargar y extraer .tar.gz
	grails_tar_url := "https://github.com/grails/grails-core/releases/download/v6.2.0/grails-6.2.0.zip"
	grails_cmd_tar := exec.Command("curl", "-L", grails_tar_url, "-o", "package.tar.gz")
	err := grails_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grails_zip_url := "https://github.com/grails/grails-core/releases/download/v6.2.0/grails-6.2.0.zip"
	grails_cmd_zip := exec.Command("curl", "-L", grails_zip_url, "-o", "package.zip")
	err = grails_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grails_bin_url := "https://github.com/grails/grails-core/releases/download/v6.2.0/grails-6.2.0.zip"
	grails_cmd_bin := exec.Command("curl", "-L", grails_bin_url, "-o", "binary.bin")
	err = grails_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grails_src_url := "https://github.com/grails/grails-core/releases/download/v6.2.0/grails-6.2.0.zip"
	grails_cmd_src := exec.Command("curl", "-L", grails_src_url, "-o", "source.tar.gz")
	err = grails_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grails_cmd_direct := exec.Command("./binary")
	err = grails_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
