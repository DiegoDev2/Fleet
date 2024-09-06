package main

// BashUnit - Bash unit testing enterprise edition framework for professionals
// Homepage: https://github.com/pgrange/bash_unit

import (
	"fmt"
	
	"os/exec"
)

func installBashUnit() {
	// Método 1: Descargar y extraer .tar.gz
	bashunit_tar_url := "https://github.com/pgrange/bash_unit/archive/refs/tags/v2.3.1.tar.gz"
	bashunit_cmd_tar := exec.Command("curl", "-L", bashunit_tar_url, "-o", "package.tar.gz")
	err := bashunit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bashunit_zip_url := "https://github.com/pgrange/bash_unit/archive/refs/tags/v2.3.1.zip"
	bashunit_cmd_zip := exec.Command("curl", "-L", bashunit_zip_url, "-o", "package.zip")
	err = bashunit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bashunit_bin_url := "https://github.com/pgrange/bash_unit/archive/refs/tags/v2.3.1.bin"
	bashunit_cmd_bin := exec.Command("curl", "-L", bashunit_bin_url, "-o", "binary.bin")
	err = bashunit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bashunit_src_url := "https://github.com/pgrange/bash_unit/archive/refs/tags/v2.3.1.src.tar.gz"
	bashunit_cmd_src := exec.Command("curl", "-L", bashunit_src_url, "-o", "source.tar.gz")
	err = bashunit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bashunit_cmd_direct := exec.Command("./binary")
	err = bashunit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
