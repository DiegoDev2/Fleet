package main

// Darglint - Python docstring argument linter
// Homepage: https://github.com/terrencepreilly/darglint

import (
	"fmt"
	
	"os/exec"
)

func installDarglint() {
	// Método 1: Descargar y extraer .tar.gz
	darglint_tar_url := "https://files.pythonhosted.org/packages/d4/2c/86e8549e349388c18ca8a4ff8661bb5347da550f598656d32a98eaaf91cc/darglint-1.8.1.tar.gz"
	darglint_cmd_tar := exec.Command("curl", "-L", darglint_tar_url, "-o", "package.tar.gz")
	err := darglint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	darglint_zip_url := "https://files.pythonhosted.org/packages/d4/2c/86e8549e349388c18ca8a4ff8661bb5347da550f598656d32a98eaaf91cc/darglint-1.8.1.zip"
	darglint_cmd_zip := exec.Command("curl", "-L", darglint_zip_url, "-o", "package.zip")
	err = darglint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	darglint_bin_url := "https://files.pythonhosted.org/packages/d4/2c/86e8549e349388c18ca8a4ff8661bb5347da550f598656d32a98eaaf91cc/darglint-1.8.1.bin"
	darglint_cmd_bin := exec.Command("curl", "-L", darglint_bin_url, "-o", "binary.bin")
	err = darglint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	darglint_src_url := "https://files.pythonhosted.org/packages/d4/2c/86e8549e349388c18ca8a4ff8661bb5347da550f598656d32a98eaaf91cc/darglint-1.8.1.src.tar.gz"
	darglint_cmd_src := exec.Command("curl", "-L", darglint_src_url, "-o", "source.tar.gz")
	err = darglint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	darglint_cmd_direct := exec.Command("./binary")
	err = darglint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
}
