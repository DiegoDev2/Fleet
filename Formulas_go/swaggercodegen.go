package main

// SwaggerCodegen - Generate clients, server stubs, and docs from an OpenAPI spec
// Homepage: https://swagger.io/tools/swagger-codegen/

import (
	"fmt"
	
	"os/exec"
)

func installSwaggerCodegen() {
	// Método 1: Descargar y extraer .tar.gz
	swaggercodegen_tar_url := "https://github.com/swagger-api/swagger-codegen/archive/refs/tags/v3.0.62.tar.gz"
	swaggercodegen_cmd_tar := exec.Command("curl", "-L", swaggercodegen_tar_url, "-o", "package.tar.gz")
	err := swaggercodegen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swaggercodegen_zip_url := "https://github.com/swagger-api/swagger-codegen/archive/refs/tags/v3.0.62.zip"
	swaggercodegen_cmd_zip := exec.Command("curl", "-L", swaggercodegen_zip_url, "-o", "package.zip")
	err = swaggercodegen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swaggercodegen_bin_url := "https://github.com/swagger-api/swagger-codegen/archive/refs/tags/v3.0.62.bin"
	swaggercodegen_cmd_bin := exec.Command("curl", "-L", swaggercodegen_bin_url, "-o", "binary.bin")
	err = swaggercodegen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swaggercodegen_src_url := "https://github.com/swagger-api/swagger-codegen/archive/refs/tags/v3.0.62.src.tar.gz"
	swaggercodegen_cmd_src := exec.Command("curl", "-L", swaggercodegen_src_url, "-o", "source.tar.gz")
	err = swaggercodegen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swaggercodegen_cmd_direct := exec.Command("./binary")
	err = swaggercodegen_cmd_direct.Run()
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
