package main

// GoogleSqlTool - Command-line tool for executing common SQL statements
// Homepage: https://cloud.google.com/sql/docs/mysql-client

import (
	"fmt"
	
	"os/exec"
)

func installGoogleSqlTool() {
	// Método 1: Descargar y extraer .tar.gz
	googlesqltool_tar_url := "https://dl.google.com/cloudsql/tools/google_sql_tool.zip"
	googlesqltool_cmd_tar := exec.Command("curl", "-L", googlesqltool_tar_url, "-o", "package.tar.gz")
	err := googlesqltool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	googlesqltool_zip_url := "https://dl.google.com/cloudsql/tools/google_sql_tool.zip"
	googlesqltool_cmd_zip := exec.Command("curl", "-L", googlesqltool_zip_url, "-o", "package.zip")
	err = googlesqltool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	googlesqltool_bin_url := "https://dl.google.com/cloudsql/tools/google_sql_tool.zip"
	googlesqltool_cmd_bin := exec.Command("curl", "-L", googlesqltool_bin_url, "-o", "binary.bin")
	err = googlesqltool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	googlesqltool_src_url := "https://dl.google.com/cloudsql/tools/google_sql_tool.zip"
	googlesqltool_cmd_src := exec.Command("curl", "-L", googlesqltool_src_url, "-o", "source.tar.gz")
	err = googlesqltool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	googlesqltool_cmd_direct := exec.Command("./binary")
	err = googlesqltool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
