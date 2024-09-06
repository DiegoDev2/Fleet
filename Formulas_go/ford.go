package main

// Ford - Automatic documentation generator for modern Fortran programs
// Homepage: https://github.com/Fortran-FOSS-Programmers/ford

import (
	"fmt"
	
	"os/exec"
)

func installFord() {
	// Método 1: Descargar y extraer .tar.gz
	ford_tar_url := "https://files.pythonhosted.org/packages/a2/00/1dee70777917617df2c63bef8db8ec4e8a68495fae0d77b9208cdda6b458/ford-7.0.8.tar.gz"
	ford_cmd_tar := exec.Command("curl", "-L", ford_tar_url, "-o", "package.tar.gz")
	err := ford_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ford_zip_url := "https://files.pythonhosted.org/packages/a2/00/1dee70777917617df2c63bef8db8ec4e8a68495fae0d77b9208cdda6b458/ford-7.0.8.zip"
	ford_cmd_zip := exec.Command("curl", "-L", ford_zip_url, "-o", "package.zip")
	err = ford_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ford_bin_url := "https://files.pythonhosted.org/packages/a2/00/1dee70777917617df2c63bef8db8ec4e8a68495fae0d77b9208cdda6b458/ford-7.0.8.bin"
	ford_cmd_bin := exec.Command("curl", "-L", ford_bin_url, "-o", "binary.bin")
	err = ford_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ford_src_url := "https://files.pythonhosted.org/packages/a2/00/1dee70777917617df2c63bef8db8ec4e8a68495fae0d77b9208cdda6b458/ford-7.0.8.src.tar.gz"
	ford_cmd_src := exec.Command("curl", "-L", ford_src_url, "-o", "source.tar.gz")
	err = ford_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ford_cmd_direct := exec.Command("./binary")
	err = ford_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: graphviz")
exec.Command("latte", "install", "graphviz")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
