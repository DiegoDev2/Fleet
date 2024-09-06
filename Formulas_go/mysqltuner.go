package main

// Mysqltuner - Increase performance and stability of a MySQL installation
// Homepage: https://raw.github.com/major/MySQLTuner-perl/master/mysqltuner.pl

import (
	"fmt"
	
	"os/exec"
)

func installMysqltuner() {
	// Método 1: Descargar y extraer .tar.gz
	mysqltuner_tar_url := "https://github.com/major/MySQLTuner-perl/archive/refs/tags/v2.6.0.tar.gz"
	mysqltuner_cmd_tar := exec.Command("curl", "-L", mysqltuner_tar_url, "-o", "package.tar.gz")
	err := mysqltuner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mysqltuner_zip_url := "https://github.com/major/MySQLTuner-perl/archive/refs/tags/v2.6.0.zip"
	mysqltuner_cmd_zip := exec.Command("curl", "-L", mysqltuner_zip_url, "-o", "package.zip")
	err = mysqltuner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mysqltuner_bin_url := "https://github.com/major/MySQLTuner-perl/archive/refs/tags/v2.6.0.bin"
	mysqltuner_cmd_bin := exec.Command("curl", "-L", mysqltuner_bin_url, "-o", "binary.bin")
	err = mysqltuner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mysqltuner_src_url := "https://github.com/major/MySQLTuner-perl/archive/refs/tags/v2.6.0.src.tar.gz"
	mysqltuner_cmd_src := exec.Command("curl", "-L", mysqltuner_src_url, "-o", "source.tar.gz")
	err = mysqltuner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mysqltuner_cmd_direct := exec.Command("./binary")
	err = mysqltuner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
