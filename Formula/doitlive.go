package main

// Doitlive - Replay stored shell commands for live presentations
// Homepage: https://doitlive.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installDoitlive() {
	// Método 1: Descargar y extraer .tar.gz
	doitlive_tar_url := "https://files.pythonhosted.org/packages/32/75/c94e4d4e7fac8606e199fad35a00b33e4252d00078f25285f91e97e546c0/doitlive-5.1.0.tar.gz"
	doitlive_cmd_tar := exec.Command("curl", "-L", doitlive_tar_url, "-o", "package.tar.gz")
	err := doitlive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	doitlive_zip_url := "https://files.pythonhosted.org/packages/32/75/c94e4d4e7fac8606e199fad35a00b33e4252d00078f25285f91e97e546c0/doitlive-5.1.0.zip"
	doitlive_cmd_zip := exec.Command("curl", "-L", doitlive_zip_url, "-o", "package.zip")
	err = doitlive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	doitlive_bin_url := "https://files.pythonhosted.org/packages/32/75/c94e4d4e7fac8606e199fad35a00b33e4252d00078f25285f91e97e546c0/doitlive-5.1.0.bin"
	doitlive_cmd_bin := exec.Command("curl", "-L", doitlive_bin_url, "-o", "binary.bin")
	err = doitlive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	doitlive_src_url := "https://files.pythonhosted.org/packages/32/75/c94e4d4e7fac8606e199fad35a00b33e4252d00078f25285f91e97e546c0/doitlive-5.1.0.src.tar.gz"
	doitlive_cmd_src := exec.Command("curl", "-L", doitlive_src_url, "-o", "source.tar.gz")
	err = doitlive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	doitlive_cmd_direct := exec.Command("./binary")
	err = doitlive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
