package main

// GulpCli - Command-line utility for Gulp
// Homepage: https://github.com/gulpjs/gulp-cli

import (
	"fmt"
	
	"os/exec"
)

func installGulpCli() {
	// Método 1: Descargar y extraer .tar.gz
	gulpcli_tar_url := "https://registry.npmjs.org/gulp-cli/-/gulp-cli-3.0.0.tgz"
	gulpcli_cmd_tar := exec.Command("curl", "-L", gulpcli_tar_url, "-o", "package.tar.gz")
	err := gulpcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gulpcli_zip_url := "https://registry.npmjs.org/gulp-cli/-/gulp-cli-3.0.0.tgz"
	gulpcli_cmd_zip := exec.Command("curl", "-L", gulpcli_zip_url, "-o", "package.zip")
	err = gulpcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gulpcli_bin_url := "https://registry.npmjs.org/gulp-cli/-/gulp-cli-3.0.0.tgz"
	gulpcli_cmd_bin := exec.Command("curl", "-L", gulpcli_bin_url, "-o", "binary.bin")
	err = gulpcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gulpcli_src_url := "https://registry.npmjs.org/gulp-cli/-/gulp-cli-3.0.0.tgz"
	gulpcli_cmd_src := exec.Command("curl", "-L", gulpcli_src_url, "-o", "source.tar.gz")
	err = gulpcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gulpcli_cmd_direct := exec.Command("./binary")
	err = gulpcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
