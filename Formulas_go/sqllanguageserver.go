package main

// SqlLanguageServer - Language Server for SQL
// Homepage: https://github.com/joe-re/sql-language-server

import (
	"fmt"
	
	"os/exec"
)

func installSqlLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	sqllanguageserver_tar_url := "https://registry.npmjs.org/sql-language-server/-/sql-language-server-1.7.0.tgz"
	sqllanguageserver_cmd_tar := exec.Command("curl", "-L", sqllanguageserver_tar_url, "-o", "package.tar.gz")
	err := sqllanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqllanguageserver_zip_url := "https://registry.npmjs.org/sql-language-server/-/sql-language-server-1.7.0.tgz"
	sqllanguageserver_cmd_zip := exec.Command("curl", "-L", sqllanguageserver_zip_url, "-o", "package.zip")
	err = sqllanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqllanguageserver_bin_url := "https://registry.npmjs.org/sql-language-server/-/sql-language-server-1.7.0.tgz"
	sqllanguageserver_cmd_bin := exec.Command("curl", "-L", sqllanguageserver_bin_url, "-o", "binary.bin")
	err = sqllanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqllanguageserver_src_url := "https://registry.npmjs.org/sql-language-server/-/sql-language-server-1.7.0.tgz"
	sqllanguageserver_cmd_src := exec.Command("curl", "-L", sqllanguageserver_src_url, "-o", "source.tar.gz")
	err = sqllanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqllanguageserver_cmd_direct := exec.Command("./binary")
	err = sqllanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: terminal-notifier")
exec.Command("latte", "install", "terminal-notifier")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
}
