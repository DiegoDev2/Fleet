package main

// Sqlfluff - SQL linter and auto-formatter for Humans
// Homepage: https://docs.sqlfluff.com/

import (
	"fmt"
	
	"os/exec"
)

func installSqlfluff() {
	// Método 1: Descargar y extraer .tar.gz
	sqlfluff_tar_url := "https://files.pythonhosted.org/packages/fa/bc/78715cf0193ef90135661b3cf52e1288eadb331aeb55162eadb48de3c5ff/sqlfluff-3.1.1.tar.gz"
	sqlfluff_cmd_tar := exec.Command("curl", "-L", sqlfluff_tar_url, "-o", "package.tar.gz")
	err := sqlfluff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sqlfluff_zip_url := "https://files.pythonhosted.org/packages/fa/bc/78715cf0193ef90135661b3cf52e1288eadb331aeb55162eadb48de3c5ff/sqlfluff-3.1.1.zip"
	sqlfluff_cmd_zip := exec.Command("curl", "-L", sqlfluff_zip_url, "-o", "package.zip")
	err = sqlfluff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sqlfluff_bin_url := "https://files.pythonhosted.org/packages/fa/bc/78715cf0193ef90135661b3cf52e1288eadb331aeb55162eadb48de3c5ff/sqlfluff-3.1.1.bin"
	sqlfluff_cmd_bin := exec.Command("curl", "-L", sqlfluff_bin_url, "-o", "binary.bin")
	err = sqlfluff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sqlfluff_src_url := "https://files.pythonhosted.org/packages/fa/bc/78715cf0193ef90135661b3cf52e1288eadb331aeb55162eadb48de3c5ff/sqlfluff-3.1.1.src.tar.gz"
	sqlfluff_cmd_src := exec.Command("curl", "-L", sqlfluff_src_url, "-o", "source.tar.gz")
	err = sqlfluff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sqlfluff_cmd_direct := exec.Command("./binary")
	err = sqlfluff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
