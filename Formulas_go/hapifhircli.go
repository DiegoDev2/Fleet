package main

// HapiFhirCli - Command-line interface for the HAPI FHIR library
// Homepage: https://hapifhir.io/

import (
	"fmt"
	
	"os/exec"
)

func installHapiFhirCli() {
	// Método 1: Descargar y extraer .tar.gz
	hapifhircli_tar_url := "https://github.com/hapifhir/hapi-fhir/releases/download/v7.4.0/hapi-fhir-7.4.0-cli.zip"
	hapifhircli_cmd_tar := exec.Command("curl", "-L", hapifhircli_tar_url, "-o", "package.tar.gz")
	err := hapifhircli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hapifhircli_zip_url := "https://github.com/hapifhir/hapi-fhir/releases/download/v7.4.0/hapi-fhir-7.4.0-cli.zip"
	hapifhircli_cmd_zip := exec.Command("curl", "-L", hapifhircli_zip_url, "-o", "package.zip")
	err = hapifhircli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hapifhircli_bin_url := "https://github.com/hapifhir/hapi-fhir/releases/download/v7.4.0/hapi-fhir-7.4.0-cli.zip"
	hapifhircli_cmd_bin := exec.Command("curl", "-L", hapifhircli_bin_url, "-o", "binary.bin")
	err = hapifhircli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hapifhircli_src_url := "https://github.com/hapifhir/hapi-fhir/releases/download/v7.4.0/hapi-fhir-7.4.0-cli.zip"
	hapifhircli_cmd_src := exec.Command("curl", "-L", hapifhircli_src_url, "-o", "source.tar.gz")
	err = hapifhircli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hapifhircli_cmd_direct := exec.Command("./binary")
	err = hapifhircli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
