package main

// GraphqlCli - Command-line tool for common GraphQL development workflows
// Homepage: https://github.com/Urigo/graphql-cli

import (
	"fmt"
	
	"os/exec"
)

func installGraphqlCli() {
	// Método 1: Descargar y extraer .tar.gz
	graphqlcli_tar_url := "https://registry.npmjs.org/graphql-cli/-/graphql-cli-4.1.0.tgz"
	graphqlcli_cmd_tar := exec.Command("curl", "-L", graphqlcli_tar_url, "-o", "package.tar.gz")
	err := graphqlcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	graphqlcli_zip_url := "https://registry.npmjs.org/graphql-cli/-/graphql-cli-4.1.0.tgz"
	graphqlcli_cmd_zip := exec.Command("curl", "-L", graphqlcli_zip_url, "-o", "package.zip")
	err = graphqlcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	graphqlcli_bin_url := "https://registry.npmjs.org/graphql-cli/-/graphql-cli-4.1.0.tgz"
	graphqlcli_cmd_bin := exec.Command("curl", "-L", graphqlcli_bin_url, "-o", "binary.bin")
	err = graphqlcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	graphqlcli_src_url := "https://registry.npmjs.org/graphql-cli/-/graphql-cli-4.1.0.tgz"
	graphqlcli_cmd_src := exec.Command("curl", "-L", graphqlcli_src_url, "-o", "source.tar.gz")
	err = graphqlcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	graphqlcli_cmd_direct := exec.Command("./binary")
	err = graphqlcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
