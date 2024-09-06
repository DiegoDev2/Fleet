package main

// Csvkit - Suite of command-line tools for converting to and working with CSV
// Homepage: https://csvkit.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installCsvkit() {
	// Método 1: Descargar y extraer .tar.gz
	csvkit_tar_url := "https://files.pythonhosted.org/packages/b6/29/51d7c3221669a4a63410f9be61178436109217d77a31b539f41ef6c1448e/csvkit-2.0.1.tar.gz"
	csvkit_cmd_tar := exec.Command("curl", "-L", csvkit_tar_url, "-o", "package.tar.gz")
	err := csvkit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	csvkit_zip_url := "https://files.pythonhosted.org/packages/b6/29/51d7c3221669a4a63410f9be61178436109217d77a31b539f41ef6c1448e/csvkit-2.0.1.zip"
	csvkit_cmd_zip := exec.Command("curl", "-L", csvkit_zip_url, "-o", "package.zip")
	err = csvkit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	csvkit_bin_url := "https://files.pythonhosted.org/packages/b6/29/51d7c3221669a4a63410f9be61178436109217d77a31b539f41ef6c1448e/csvkit-2.0.1.bin"
	csvkit_cmd_bin := exec.Command("curl", "-L", csvkit_bin_url, "-o", "binary.bin")
	err = csvkit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	csvkit_src_url := "https://files.pythonhosted.org/packages/b6/29/51d7c3221669a4a63410f9be61178436109217d77a31b539f41ef6c1448e/csvkit-2.0.1.src.tar.gz"
	csvkit_cmd_src := exec.Command("curl", "-L", csvkit_src_url, "-o", "source.tar.gz")
	err = csvkit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	csvkit_cmd_direct := exec.Command("./binary")
	err = csvkit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
