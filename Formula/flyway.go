package main

// Flyway - Database version control to control migrations
// Homepage: https://flywaydb.org/

import (
	"fmt"
	
	"os/exec"
)

func installFlyway() {
	// Método 1: Descargar y extraer .tar.gz
	flyway_tar_url := "https://search.maven.org/remotecontent?filepath=org/flywaydb/flyway-commandline/10.17.3/flyway-commandline-10.17.3.tar.gz"
	flyway_cmd_tar := exec.Command("curl", "-L", flyway_tar_url, "-o", "package.tar.gz")
	err := flyway_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flyway_zip_url := "https://search.maven.org/remotecontent?filepath=org/flywaydb/flyway-commandline/10.17.3/flyway-commandline-10.17.3.zip"
	flyway_cmd_zip := exec.Command("curl", "-L", flyway_zip_url, "-o", "package.zip")
	err = flyway_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flyway_bin_url := "https://search.maven.org/remotecontent?filepath=org/flywaydb/flyway-commandline/10.17.3/flyway-commandline-10.17.3.bin"
	flyway_cmd_bin := exec.Command("curl", "-L", flyway_bin_url, "-o", "binary.bin")
	err = flyway_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flyway_src_url := "https://search.maven.org/remotecontent?filepath=org/flywaydb/flyway-commandline/10.17.3/flyway-commandline-10.17.3.src.tar.gz"
	flyway_cmd_src := exec.Command("curl", "-L", flyway_src_url, "-o", "source.tar.gz")
	err = flyway_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flyway_cmd_direct := exec.Command("./binary")
	err = flyway_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
