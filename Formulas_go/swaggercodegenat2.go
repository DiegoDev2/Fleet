package main

// SwaggerCodegenAT2 - Generate clients, server stubs, and docs from an OpenAPI spec
// Homepage: https://swagger.io/tools/swagger-codegen/

import (
	"fmt"
	
	"os/exec"
)

func installSwaggerCodegenAT2() {
	// Método 1: Descargar y extraer .tar.gz
	swaggercodegenat2_tar_url := "https://github.com/swagger-api/swagger-codegen/archive/refs/tags/v2.4.43.tar.gz"
	swaggercodegenat2_cmd_tar := exec.Command("curl", "-L", swaggercodegenat2_tar_url, "-o", "package.tar.gz")
	err := swaggercodegenat2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swaggercodegenat2_zip_url := "https://github.com/swagger-api/swagger-codegen/archive/refs/tags/v2.4.43.zip"
	swaggercodegenat2_cmd_zip := exec.Command("curl", "-L", swaggercodegenat2_zip_url, "-o", "package.zip")
	err = swaggercodegenat2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swaggercodegenat2_bin_url := "https://github.com/swagger-api/swagger-codegen/archive/refs/tags/v2.4.43.bin"
	swaggercodegenat2_cmd_bin := exec.Command("curl", "-L", swaggercodegenat2_bin_url, "-o", "binary.bin")
	err = swaggercodegenat2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swaggercodegenat2_src_url := "https://github.com/swagger-api/swagger-codegen/archive/refs/tags/v2.4.43.src.tar.gz"
	swaggercodegenat2_cmd_src := exec.Command("curl", "-L", swaggercodegenat2_src_url, "-o", "source.tar.gz")
	err = swaggercodegenat2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swaggercodegenat2_cmd_direct := exec.Command("./binary")
	err = swaggercodegenat2_cmd_direct.Run()
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
