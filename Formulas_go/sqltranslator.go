package main

// SqlTranslator - Manipulate structured data definitions (SQL and more)
// Homepage: https://github.com/dbsrgits/sql-translator/

import (
	"fmt"
	
	"os/exec"
)

func installSqlTranslator() {
	// Método 1: Descargar y extraer .tar.gz
	sqltranslator_tar_url := "https://cpan.metacpan.org/authors/id/I/IL/ILMARI/SQL-Translator-1.62.tar.gz"
	sqltranslator_cmd_tar := exec.Command("curl", "-L", sqltranslator_tar_url, "-o", "package.tar.gz")
	err := sqltranslator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqltranslator_zip_url := "https://cpan.metacpan.org/authors/id/I/IL/ILMARI/SQL-Translator-1.62.zip"
	sqltranslator_cmd_zip := exec.Command("curl", "-L", sqltranslator_zip_url, "-o", "package.zip")
	err = sqltranslator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqltranslator_bin_url := "https://cpan.metacpan.org/authors/id/I/IL/ILMARI/SQL-Translator-1.62.bin"
	sqltranslator_cmd_bin := exec.Command("curl", "-L", sqltranslator_bin_url, "-o", "binary.bin")
	err = sqltranslator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqltranslator_src_url := "https://cpan.metacpan.org/authors/id/I/IL/ILMARI/SQL-Translator-1.62.src.tar.gz"
	sqltranslator_cmd_src := exec.Command("curl", "-L", sqltranslator_src_url, "-o", "source.tar.gz")
	err = sqltranslator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqltranslator_cmd_direct := exec.Command("./binary")
	err = sqltranslator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
