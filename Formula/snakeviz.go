package main

// Snakeviz - Web-based viewer for Python profiler output
// Homepage: https://jiffyclub.github.io/snakeviz/

import (
	"fmt"
	
	"os/exec"
)

func installSnakeviz() {
	// Método 1: Descargar y extraer .tar.gz
	snakeviz_tar_url := "https://files.pythonhosted.org/packages/64/9b/3983c41e913676d55e4b3de869aa0561e053ad3505f1fd35181670244b70/snakeviz-2.2.0.tar.gz"
	snakeviz_cmd_tar := exec.Command("curl", "-L", snakeviz_tar_url, "-o", "package.tar.gz")
	err := snakeviz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	snakeviz_zip_url := "https://files.pythonhosted.org/packages/64/9b/3983c41e913676d55e4b3de869aa0561e053ad3505f1fd35181670244b70/snakeviz-2.2.0.zip"
	snakeviz_cmd_zip := exec.Command("curl", "-L", snakeviz_zip_url, "-o", "package.zip")
	err = snakeviz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	snakeviz_bin_url := "https://files.pythonhosted.org/packages/64/9b/3983c41e913676d55e4b3de869aa0561e053ad3505f1fd35181670244b70/snakeviz-2.2.0.bin"
	snakeviz_cmd_bin := exec.Command("curl", "-L", snakeviz_bin_url, "-o", "binary.bin")
	err = snakeviz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	snakeviz_src_url := "https://files.pythonhosted.org/packages/64/9b/3983c41e913676d55e4b3de869aa0561e053ad3505f1fd35181670244b70/snakeviz-2.2.0.src.tar.gz"
	snakeviz_cmd_src := exec.Command("curl", "-L", snakeviz_src_url, "-o", "source.tar.gz")
	err = snakeviz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	snakeviz_cmd_direct := exec.Command("./binary")
	err = snakeviz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
