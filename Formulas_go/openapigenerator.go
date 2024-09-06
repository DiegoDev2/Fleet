package main

// OpenapiGenerator - Generate clients, server & docs from an OpenAPI spec (v2, v3)
// Homepage: https://openapi-generator.tech/

import (
	"fmt"
	
	"os/exec"
)

func installOpenapiGenerator() {
	// Método 1: Descargar y extraer .tar.gz
	openapigenerator_tar_url := "https://search.maven.org/remotecontent?filepath=org/openapitools/openapi-generator-cli/7.8.0/openapi-generator-cli-7.8.0.jar"
	openapigenerator_cmd_tar := exec.Command("curl", "-L", openapigenerator_tar_url, "-o", "package.tar.gz")
	err := openapigenerator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openapigenerator_zip_url := "https://search.maven.org/remotecontent?filepath=org/openapitools/openapi-generator-cli/7.8.0/openapi-generator-cli-7.8.0.jar"
	openapigenerator_cmd_zip := exec.Command("curl", "-L", openapigenerator_zip_url, "-o", "package.zip")
	err = openapigenerator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openapigenerator_bin_url := "https://search.maven.org/remotecontent?filepath=org/openapitools/openapi-generator-cli/7.8.0/openapi-generator-cli-7.8.0.jar"
	openapigenerator_cmd_bin := exec.Command("curl", "-L", openapigenerator_bin_url, "-o", "binary.bin")
	err = openapigenerator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openapigenerator_src_url := "https://search.maven.org/remotecontent?filepath=org/openapitools/openapi-generator-cli/7.8.0/openapi-generator-cli-7.8.0.jar"
	openapigenerator_cmd_src := exec.Command("curl", "-L", openapigenerator_src_url, "-o", "source.tar.gz")
	err = openapigenerator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openapigenerator_cmd_direct := exec.Command("./binary")
	err = openapigenerator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: maven")
exec.Command("latte", "install", "maven")
	fmt.Println("Instalando dependencia: openjdk@11")
exec.Command("latte", "install", "openjdk@11")
}
