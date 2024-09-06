package main

// Fobis - KISS build tool for automatically building modern Fortran projects
// Homepage: https://github.com/szaghi/FoBiS

import (
	"fmt"
	
	"os/exec"
)

func installFobis() {
	// Método 1: Descargar y extraer .tar.gz
	fobis_tar_url := "https://files.pythonhosted.org/packages/53/3a/5533ab0277977027478b4c1285bb20b6beb221b222403b10398fb24e81a2/FoBiS.py-3.0.5.tar.gz"
	fobis_cmd_tar := exec.Command("curl", "-L", fobis_tar_url, "-o", "package.tar.gz")
	err := fobis_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fobis_zip_url := "https://files.pythonhosted.org/packages/53/3a/5533ab0277977027478b4c1285bb20b6beb221b222403b10398fb24e81a2/FoBiS.py-3.0.5.zip"
	fobis_cmd_zip := exec.Command("curl", "-L", fobis_zip_url, "-o", "package.zip")
	err = fobis_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fobis_bin_url := "https://files.pythonhosted.org/packages/53/3a/5533ab0277977027478b4c1285bb20b6beb221b222403b10398fb24e81a2/FoBiS.py-3.0.5.bin"
	fobis_cmd_bin := exec.Command("curl", "-L", fobis_bin_url, "-o", "binary.bin")
	err = fobis_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fobis_src_url := "https://files.pythonhosted.org/packages/53/3a/5533ab0277977027478b4c1285bb20b6beb221b222403b10398fb24e81a2/FoBiS.py-3.0.5.src.tar.gz"
	fobis_cmd_src := exec.Command("curl", "-L", fobis_src_url, "-o", "source.tar.gz")
	err = fobis_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fobis_cmd_direct := exec.Command("./binary")
	err = fobis_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: graphviz")
	exec.Command("latte", "install", "graphviz").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
