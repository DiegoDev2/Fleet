package main

// DbmlCli - Convert DBML file to SQL and vice versa
// Homepage: https://www.dbml.org/cli/

import (
	"fmt"
	
	"os/exec"
)

func installDbmlCli() {
	// Método 1: Descargar y extraer .tar.gz
	dbmlcli_tar_url := "https://registry.npmjs.org/@dbml/cli/-/cli-3.8.0.tgz"
	dbmlcli_cmd_tar := exec.Command("curl", "-L", dbmlcli_tar_url, "-o", "package.tar.gz")
	err := dbmlcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dbmlcli_zip_url := "https://registry.npmjs.org/@dbml/cli/-/cli-3.8.0.tgz"
	dbmlcli_cmd_zip := exec.Command("curl", "-L", dbmlcli_zip_url, "-o", "package.zip")
	err = dbmlcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dbmlcli_bin_url := "https://registry.npmjs.org/@dbml/cli/-/cli-3.8.0.tgz"
	dbmlcli_cmd_bin := exec.Command("curl", "-L", dbmlcli_bin_url, "-o", "binary.bin")
	err = dbmlcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dbmlcli_src_url := "https://registry.npmjs.org/@dbml/cli/-/cli-3.8.0.tgz"
	dbmlcli_cmd_src := exec.Command("curl", "-L", dbmlcli_src_url, "-o", "source.tar.gz")
	err = dbmlcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dbmlcli_cmd_direct := exec.Command("./binary")
	err = dbmlcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
