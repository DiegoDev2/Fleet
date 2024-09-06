package main

// Csvq - SQL-like query language for csv
// Homepage: https://mithrandie.github.io/csvq

import (
	"fmt"
	
	"os/exec"
)

func installCsvq() {
	// Método 1: Descargar y extraer .tar.gz
	csvq_tar_url := "https://github.com/mithrandie/csvq/archive/refs/tags/v1.18.1.tar.gz"
	csvq_cmd_tar := exec.Command("curl", "-L", csvq_tar_url, "-o", "package.tar.gz")
	err := csvq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csvq_zip_url := "https://github.com/mithrandie/csvq/archive/refs/tags/v1.18.1.zip"
	csvq_cmd_zip := exec.Command("curl", "-L", csvq_zip_url, "-o", "package.zip")
	err = csvq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csvq_bin_url := "https://github.com/mithrandie/csvq/archive/refs/tags/v1.18.1.bin"
	csvq_cmd_bin := exec.Command("curl", "-L", csvq_bin_url, "-o", "binary.bin")
	err = csvq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csvq_src_url := "https://github.com/mithrandie/csvq/archive/refs/tags/v1.18.1.src.tar.gz"
	csvq_cmd_src := exec.Command("curl", "-L", csvq_src_url, "-o", "source.tar.gz")
	err = csvq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csvq_cmd_direct := exec.Command("./binary")
	err = csvq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
