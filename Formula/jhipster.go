package main

// Jhipster - Generate, develop and deploy Spring Boot + Angular/React applications
// Homepage: https://www.jhipster.tech/

import (
	"fmt"
	
	"os/exec"
)

func installJhipster() {
	// Método 1: Descargar y extraer .tar.gz
	jhipster_tar_url := "https://registry.npmjs.org/generator-jhipster/-/generator-jhipster-8.7.0.tgz"
	jhipster_cmd_tar := exec.Command("curl", "-L", jhipster_tar_url, "-o", "package.tar.gz")
	err := jhipster_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jhipster_zip_url := "https://registry.npmjs.org/generator-jhipster/-/generator-jhipster-8.7.0.tgz"
	jhipster_cmd_zip := exec.Command("curl", "-L", jhipster_zip_url, "-o", "package.zip")
	err = jhipster_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jhipster_bin_url := "https://registry.npmjs.org/generator-jhipster/-/generator-jhipster-8.7.0.tgz"
	jhipster_cmd_bin := exec.Command("curl", "-L", jhipster_bin_url, "-o", "binary.bin")
	err = jhipster_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jhipster_src_url := "https://registry.npmjs.org/generator-jhipster/-/generator-jhipster-8.7.0.tgz"
	jhipster_cmd_src := exec.Command("curl", "-L", jhipster_src_url, "-o", "source.tar.gz")
	err = jhipster_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jhipster_cmd_direct := exec.Command("./binary")
	err = jhipster_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
