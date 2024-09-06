package main

// Liquibase - Library for database change tracking
// Homepage: https://www.liquibase.org/

import (
	"fmt"
	
	"os/exec"
)

func installLiquibase() {
	// Método 1: Descargar y extraer .tar.gz
	liquibase_tar_url := "https://github.com/liquibase/liquibase/releases/download/v4.29.2/liquibase-4.29.2.tar.gz"
	liquibase_cmd_tar := exec.Command("curl", "-L", liquibase_tar_url, "-o", "package.tar.gz")
	err := liquibase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	liquibase_zip_url := "https://github.com/liquibase/liquibase/releases/download/v4.29.2/liquibase-4.29.2.zip"
	liquibase_cmd_zip := exec.Command("curl", "-L", liquibase_zip_url, "-o", "package.zip")
	err = liquibase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	liquibase_bin_url := "https://github.com/liquibase/liquibase/releases/download/v4.29.2/liquibase-4.29.2.bin"
	liquibase_cmd_bin := exec.Command("curl", "-L", liquibase_bin_url, "-o", "binary.bin")
	err = liquibase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	liquibase_src_url := "https://github.com/liquibase/liquibase/releases/download/v4.29.2/liquibase-4.29.2.src.tar.gz"
	liquibase_cmd_src := exec.Command("curl", "-L", liquibase_src_url, "-o", "source.tar.gz")
	err = liquibase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	liquibase_cmd_direct := exec.Command("./binary")
	err = liquibase_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
