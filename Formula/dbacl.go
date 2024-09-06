package main

// Dbacl - Digramic Bayesian classifier
// Homepage: https://dbacl.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDbacl() {
	// Método 1: Descargar y extraer .tar.gz
	dbacl_tar_url := "https://downloads.sourceforge.net/project/dbacl/dbacl/1.14.1/dbacl-1.14.1.tar.gz"
	dbacl_cmd_tar := exec.Command("curl", "-L", dbacl_tar_url, "-o", "package.tar.gz")
	err := dbacl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dbacl_zip_url := "https://downloads.sourceforge.net/project/dbacl/dbacl/1.14.1/dbacl-1.14.1.zip"
	dbacl_cmd_zip := exec.Command("curl", "-L", dbacl_zip_url, "-o", "package.zip")
	err = dbacl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dbacl_bin_url := "https://downloads.sourceforge.net/project/dbacl/dbacl/1.14.1/dbacl-1.14.1.bin"
	dbacl_cmd_bin := exec.Command("curl", "-L", dbacl_bin_url, "-o", "binary.bin")
	err = dbacl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dbacl_src_url := "https://downloads.sourceforge.net/project/dbacl/dbacl/1.14.1/dbacl-1.14.1.src.tar.gz"
	dbacl_cmd_src := exec.Command("curl", "-L", dbacl_src_url, "-o", "source.tar.gz")
	err = dbacl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dbacl_cmd_direct := exec.Command("./binary")
	err = dbacl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
