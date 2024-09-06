package main

// HasuraCli - Command-Line Interface for Hasura GraphQL Engine
// Homepage: https://hasura.io

import (
	"fmt"
	
	"os/exec"
)

func installHasuraCli() {
	// Método 1: Descargar y extraer .tar.gz
	hasuracli_tar_url := "https://github.com/hasura/graphql-engine/archive/refs/tags/v2.43.0.tar.gz"
	hasuracli_cmd_tar := exec.Command("curl", "-L", hasuracli_tar_url, "-o", "package.tar.gz")
	err := hasuracli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hasuracli_zip_url := "https://github.com/hasura/graphql-engine/archive/refs/tags/v2.43.0.zip"
	hasuracli_cmd_zip := exec.Command("curl", "-L", hasuracli_zip_url, "-o", "package.zip")
	err = hasuracli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hasuracli_bin_url := "https://github.com/hasura/graphql-engine/archive/refs/tags/v2.43.0.bin"
	hasuracli_cmd_bin := exec.Command("curl", "-L", hasuracli_bin_url, "-o", "binary.bin")
	err = hasuracli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hasuracli_src_url := "https://github.com/hasura/graphql-engine/archive/refs/tags/v2.43.0.src.tar.gz"
	hasuracli_cmd_src := exec.Command("curl", "-L", hasuracli_src_url, "-o", "source.tar.gz")
	err = hasuracli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hasuracli_cmd_direct := exec.Command("./binary")
	err = hasuracli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node@18")
	exec.Command("latte", "install", "node@18").Run()
}
